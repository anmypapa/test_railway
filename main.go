package main

import (
	_ "test_railway/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

