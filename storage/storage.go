package storage

import (
	"context"
	"github.com/spf13/viper"
	"github.com/valkey-io/valkey-go"
	"log/slog"
)

type Creator interface {
	Create(v *viper.Viper) *Storage
}

type Storage interface {
	Ping(ctx context.Context) *valkey.ValkeyResult
}

var storageTable = map[string]Creator{}

func registerStorage(name string, s Creator) {
	_, ok := storageTable[name]
	if ok {
		slog.Error("duplicate register workload %s", name)
	}
	storageTable[name] = s
}

func LoadStorage(name string) *Storage {
	return storageTable[name].Create(viper.Sub(name))
}
