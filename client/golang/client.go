package main

import (
	"fmt"
	"os"
)

import (
	//pkg will import path from local $GOPATH
	"github.com/bluefalconjun/cloudstb/client/golang/platform"
)

func main() {
	fmt.Println(os.Args)
	file, _ := os.Open(os.Args[1])
	fmt.Println(file)
	//new to use the struct in package platform
	hwmods := new(platform.HwModule)
	//give the var but not &var to HwModInit and go compiler will automaticlly get var addr
	platform.HwModInit(hwmods)
}