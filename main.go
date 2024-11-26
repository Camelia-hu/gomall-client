package main

import (
	"github.com/Camelia-hu/gomall-client/router"
	"github.com/Camelia-hu/gomall-client/rpc"
)

func main() {
	rpc.RpcInit()
	router.RouterInit()

}
