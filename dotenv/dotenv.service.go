package dotenv

import (
	"bufio"
	"fmt"
	"golang-api/core"
	"os"
)

type IDotenvService interface {
	core.IProvider
	Load() error
	Get(key string) string
	Set(key, value string)
}

type DotenvService struct {
	*core.Provider
	values   map[string]string
	filepath string
}

func NewDotenvService(module core.IModule) *DotenvService {
	return &DotenvService{
		Provider: core.NewProvider("DotenvService"),
		values:   make(map[string]string),
		filepath: ".env",
	}
}

func NewDotenvServiceWithPath(module core.IModule, path string) *DotenvService {
	return &DotenvService{
		Provider: core.NewProvider("DotenvService"),
		values:   make(map[string]string),
		filepath: path,
	}
}

func (ds *DotenvService) OnInit() error {
	return ds.Load()
}

func (ds *DotenvService) readFromFile() (*os.File, error) {
	file, err := os.Open(ds.filepath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Loading environment variables from %s\n", ds.filepath)
	return file, nil
}

func (ds *DotenvService) readFromEnv() {
	for _, env := range os.Environ() {
		key, value := parseLine(env)
		if key == "" {
			continue
		}
		ds.Set(key, value)
	}
}

func (ds *DotenvService) Load() error {
	file, err := ds.readFromFile()
	if err != nil {
		fmt.Println("Loading environment variables from OS environment")
		ds.readFromEnv()
		return nil
		// fmt.Println("Loading environment variables from standard input")
		// file = os.Stdin
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key, value := parseLine(line)
		if key == "" {
			continue
		}
		ds.Set(key, value)
	}

	return scanner.Err()
}

func (ds *DotenvService) Get(key string) string {
	return ds.values[key]
}

func (ds *DotenvService) Set(key, value string) {
	ds.values[key] = value
}
