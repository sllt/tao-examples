package config

import (
	"github.com/sllt/tao/core/stores/cache"
	"github.com/sllt/tao/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
}
