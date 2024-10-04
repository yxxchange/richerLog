package main

import (
	"github.com/yxxchange/richerLog/log"
	"github.com/yxxchange/richerLog/test"
)

func main() {
	log.CustomBuilder().Build()
	log.UseDefault()
	log.Infof("Hello, world.")
	test.TestLog()
}
