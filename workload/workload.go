package workload

import (
	"context"
	"errors"
	"github.com/artikell/valkey-tpc/storage"
	"github.com/spf13/viper"
	"log/slog"
	"math/rand"
	"sync"
)

type ActionMode uint32

const (
	VarConfigName = "config"
	VarClient     = "client"
	VarClientType = "type"
	VarClientNum  = "client-num"
	VarRequestNum = "req-num"

	MaxDiceNum = 1000

	ModePrepare ActionMode = 0
	ModeBench   ActionMode = 1
)

type actionRun func(ctx context.Context, s *storage.Storage) error

type Action struct {
	name   string
	run    actionRun
	weight int
}

type WorkLoad struct {
	prepareAction actionRun
	action        []*Action
	paramsVerify  func(v *viper.Viper) error
}

var workLoadTable = map[string]*WorkLoad{}

func (w *WorkLoad) RunOneAction(ctx context.Context, n int) error {
	s := 0
	// 1. Find One Action
	var act *Action = nil
	for _, a := range w.action {
		if n < s {
			act = a
			break
		} else {
			s += a.weight
		}
	}
	if act == nil {
		return errors.New("no action")
	}
	// 2. Run Action
	err := act.run(ctx, ctx.Value(VarClient).(*storage.Storage))
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkLoad) RunOneThread(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < ctx.Value(VarClientNum).(int); i++ {
		n := rand.Intn(MaxDiceNum)
		err := w.RunOneAction(ctx, n)
		if err != nil {
			slog.Error("Action run error", err.Error())
		}
	}
	wg.Done()
}

func RunWorkLoad(name string, mode ActionMode) {
	// 1. Load Workload by Name
	w, ok := workLoadTable[name]
	if ok {
		slog.Error("duplicate register workload %s", name)
	}

	// 2. Verify Config
	err := w.paramsVerify(viper.Sub(name))
	if err != nil {
		slog.Error("workload parse param error: %s", name)
	}

	t := viper.GetString(VarClientType)
	loadStorage := storage.LoadStorage(t)

	// 3. Prepare Context
	ctx := context.WithValue(context.Background(), VarConfigName, viper.Sub(name))
	ctx = context.WithValue(ctx, VarClient, loadStorage)
	ctx = context.WithValue(ctx, VarClientNum, 100)
	ctx = context.WithValue(ctx, VarRequestNum, 100)

	// 4. Run Workload Action
	if mode == ModePrepare {
		err := w.prepareAction(ctx, loadStorage)
		if err != nil {
			return
		}
	} else {
		wg := &sync.WaitGroup{}
		for i := 0; i < ctx.Value(VarClientNum).(int); i++ {
			wg.Add(1)
			go w.RunOneThread(ctx, wg)
		}
		wg.Wait()
	}
}

func registerWorkLoad(name string, w *WorkLoad) {
	_, ok := workLoadTable[name]
	if ok {
		slog.Error("duplicate register workload %s", name)
	}
	s := 0
	for _, a := range w.action {
		s += a.weight
	}
	if s != MaxDiceNum {
		panic("register workload " + name + " failed, please check your config")
	}
	workLoadTable[name] = w
}
