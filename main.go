package main

func main() {
	mainModule := Module()
	mainModule.Init()
	mainModule.Get("MainService").(*MainService).Start()
}
