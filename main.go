package main

func main() {
	go getApi()
	go initRedis()
	initApiServer()
}
