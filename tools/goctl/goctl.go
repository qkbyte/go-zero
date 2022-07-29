package main

import (
	"github.com/qkbyte/go-zero/core/load"
	"github.com/qkbyte/go-zero/core/logx"
	"github.com/qkbyte/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
