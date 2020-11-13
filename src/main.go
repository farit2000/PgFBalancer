package main

import (
	"github.com/farit2000/PgFBalancer/src/proxy"
	"github.com/farit2000/PgFBalancer/src/web"
)

func main() {
	go proxy.StartProxy()
	web.StartWeb()
}