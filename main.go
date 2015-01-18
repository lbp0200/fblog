package main

import (
	"github.com/hoisie/web"
)

func hello(val string) string { return "hello " + val }

func main() {
	web.Get("/(.*)", hello)
	web.RunFcgi("0.0.0.0:9999")
}
