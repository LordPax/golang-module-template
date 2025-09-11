package main

func main() {
	mainModule := Module()
	if err := mainModule.Init(); err != nil {
		panic(err)
	}
	mainModule.Get("MainService").(*MainService).Start()
}
