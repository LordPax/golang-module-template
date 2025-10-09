package main

import (
	"fmt"
	"golang-api/core"
	"golang-api/database"
	"golang-api/gin"
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type MainService struct {
	*core.Provider
	databaseService *database.DatabaseService
	ginService      *gin.GinService
}

func NewMainService(module *MainModule) *MainService {
	return &MainService{
		Provider:        core.NewProvider("MainService"),
		databaseService: module.Get("DatabaseService").(*database.DatabaseService),
		ginService:      module.Get("GinService").(*gin.GinService),
	}
}

func (ms *MainService) Start() {
	defer ms.databaseService.Close()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "h", "help":
			ms.help()
			return
		case "l", "list":
			ms.list()
			return
		case "c", "call":
			ms.call()
			return
		case "graph":
			ms.graph()
			return
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			return
		}
	}

	ms.ginService.Run()

	return
}

func (ms *MainService) help() {
	fmt.Printf("Usage: %s <command> [args...] \n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("  h, help      Show this help message")
	fmt.Println("  l, list      List all events")
	fmt.Println("  c, call      Call an event")
}

func (ms *MainService) list() {
	listEvent := ms.GetModule().List()
	for _, event := range listEvent {
		fmt.Println(event)
	}
}

func (ms *MainService) call() {
	if err := ms.GetModule().Call(os.Args[2]); err != nil {
		panic(err)
	}
	if !core.CalledEvent {
		fmt.Printf("Event %s not found\n", os.Args[2])
	}
}

func (ms *MainService) graph() {
	fmt.Println("Generating graph.dot file ...")
	file, _ := os.Create("example/graph.dot")
	defer file.Close()

	g := ms.GetModule().Graph()
	graph.TopologicalSort(g)

	draw.DOT(g, file, draw.GraphAttribute("label", "Module Dependency Graph"))
}
