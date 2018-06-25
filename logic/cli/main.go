package main

import (
	"fmt"
	"os"
	"flag"
	"runtime"
	"time"
)

var (
	confFile string		// 配置文件路径
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./logic.json", "where logic.json is.")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	var (
		err error
	)

	// 初始化环境
	initArgs()
	initEnv()

	goto ERR

	for {
		time.Sleep(1 * time.Second)
	}

	os.Exit(0)

ERR:
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
