package main

func main() {
	mainModule := Module()
	mainModule.Init()
	mainService := mainModule.GetProvider("MainService").(*MainService)
	mainService.Start()
}
