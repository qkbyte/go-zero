package gateway

import (
	"time"

	"github.com/qkbyte/go-zero/rest"
	"github.com/qkbyte/go-zero/zrpc"
)

type (
	// GatewayConf is the configuration for gateway.
	GatewayConf struct {
		rest.RestConf
		Upstreams []Upstream
		Timeout   time.Duration `json:",default=5s"`
	}

	// RouteMapping is a mapping between a gateway route and an upstream rpc method.
	RouteMapping struct {
		// Method is the HTTP method, like GET, POST, PUT, DELETE.
		Method string
		// Path is the HTTP path.
		Path string
		// RpcPath is the gRPC rpc method, with format of package.service/method
		RpcPath string
	}

	// Upstream is the configuration for an upstream.
	Upstream struct {
		// Grpc is the target of the upstream.
		Grpc zrpc.RpcClientConf
		// ProtoSets is the file list of proto set, like [hello.pb]
		// if your proto file import another proto file, you need to write multi-file slice, like [hello.pb, common.pb]
		ProtoSets []string `json:",optional"`
		// Mapping is the mapping between gateway routes and Upstream rpc methods.
		// Keep it blank if annotations are added in rpc methods.
		Mapping []RouteMapping `json:",optional"`
	}
)
