package config

import (
	"github.com/sllt/tao/rest"
	"github.com/sllt/tao/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf
	Check zrpc.RpcClientConf
}
