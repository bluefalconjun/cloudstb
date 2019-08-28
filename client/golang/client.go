package main

import (
	"fmt"
	"os"
)

import (
	//pkg will import from local $GOPATH root as path
	"github.com/bluefalconjun/cloudstb/client/golang/platform"
)

func main() {
	fmt.Println(os.Args)
	file, _ := os.Open(os.Args[1])
	platform.HwModInit(file)
}