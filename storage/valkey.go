package storage

import (
	"context"
	"github.com/spf13/viper"
	"github.com/valkey-io/valkey-go"
)

const valkeyStorageName = "valkey"

type ValkeyClient struct {
}

func (r *ValkeyClient) Ping(ctx context.Context) *valkey.ValkeyResult {
	return nil
}

// ValkeyClientCreator - create client function
type ValkeyClientCreator struct {
}

func (c *ValkeyClientCreator) Create(v *viper.Viper) *Storage {
	return nil
}

func init() {
	registerStorage(valkeyStorageName, &ValkeyClientCreator{})
}
