package main

import "server/start"

func main() {
	start.InitConfig()
	start.ConnectMysql()
	start.ServerStart()
}
