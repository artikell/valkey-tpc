package workload

import (
	"context"
	"log/slog"
	"math/rand"
	"sync"
)

const MAX_DICE_NUM = 1000

type actionRun func(ctx context.Context) error

type Action struct {
	name   string
	run    actionRun
	weight int
}

type PrepareAction struct {
	opt map[string]string
	run actionRun
}

type Option struct {
	RequestNum int
	ClientNum  int
}

type WorkLoad struct {
	prepareAction *PrepareAction
	action        []*Action
}

var workLoadTable = map[string]*WorkLoad{}

func registerWorkLoad(name string, w *WorkLoad) {
	_, ok := workLoadTable[name]
	if ok {
		slog.Error("duplicate register workload %s", name)
	}
	s := 0
	for _, a := range w.action {
		s += a.weight
	}
	if s != MAX_DICE_NUM {
		panic("register workload " + name + " failed, please check your config")
	}
	workLoadTable[name] = w
}

func GetOneAction(n int, load *WorkLoad) *Action {
	s := 0
	for _, a := range load.action {
		if n < s {
			return a
		} else {
			s += a.weight
		}
	}
	return nil
}

func (w *WorkLoad) RunOneThread(wg *sync.WaitGroup, opt *Option) {
	for i := 0; i < opt.RequestNum; i++ {
		n := rand.Intn(MAX_DICE_NUM)
		action := GetOneAction(n, w)
		err := action.run(context.TODO())
		if err != nil {
			slog.Error("Action run error", err.Error())
		}
	}
	wg.Done()
}

func RunWorkLoad(name string, opt *Option) {
	w, ok := workLoadTable[name]
	if ok {
		slog.Error("duplicate register workload %s", name)
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < opt.ClientNum; i++ {
		wg.Add(1)
		go w.RunOneThread(wg, opt)
	}
	wg.Wait()
}

func RunWorkLoadPrepareAction(name string, opt *Option) {
	w, ok := workLoadTable[name]
	if ok {
		slog.Error("duplicate register workload %s", name)
	}
	err := w.prepareAction.run(context.TODO())
	if err != nil {
		return
	}
}

func getOption(ctx context.Context) Option {
	return ctx.Value("option").(Option)
}
