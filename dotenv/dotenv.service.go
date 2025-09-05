package dotenv

import (
	"bufio"
	"fmt"
	"golang-api/core"
	"os"
)

type DotenvService struct {
	*core.Provider
	values   map[string]string
	filepath string
}

func NewDotenvService(filepath string) *DotenvService {
	return &DotenvService{
		Provider: core.NewProvider("DotenvService"),
		values:   make(map[string]string),
		filepath: filepath,
	}
}

func (ds *DotenvService) OnInit() error {
	fmt.Printf("Initializing %s\n", ds.GetName())
	return ds.Load()
}

func (ds *DotenvService) Load() error {
	fmt.Printf("Loading environment variables from %s\n", ds.filepath)
	file, err := os.Open(ds.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

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
