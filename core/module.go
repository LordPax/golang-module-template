package core

import "fmt"

var moduleSeen []string

type IModule interface {
	GetName() string
	Get(provider string) IProvider
	AddProvider(p IProvider)
	AddModule(dep IModule)
	Init() error
	getProvider(provider string) IProvider
}

type Module struct {
	name     string
	module   map[string]IModule
	provider map[string]IProvider
}

func NewModule(name string) *Module {
	return &Module{
		name:     name,
		module:   make(map[string]IModule),
		provider: make(map[string]IProvider),
	}
}

func (m *Module) GetName() string {
	return m.name
}

func (m *Module) Get(provider string) IProvider {
	if p := m.getProvider(provider); p != nil {
		return p
	}

	for _, d := range m.module {
		if p := d.getProvider(provider); p != nil {
			return p
		}
	}

	panic(fmt.Sprintf("Provider %s not found", provider))
}

func (m *Module) getProvider(provider string) IProvider {
	return m.provider[provider]
}

func (m *Module) AddProvider(p IProvider) {
	m.provider[p.GetName()] = p
	p.AssignModule(m)
}

func (m *Module) AddModule(dep IModule) {
	m.module[dep.GetName()] = dep
}

func (m *Module) Init() error {
	for _, d := range m.module {
		if inArray(d.GetName(), moduleSeen) {
			continue
		}

		moduleSeen = append(moduleSeen, d.GetName())

		if err := d.Init(); err != nil {
			return err
		}
	}

	fmt.Printf("Initializing %s\n", m.GetName())

	for _, p := range m.provider {
		if err := p.OnInit(); err != nil {
			return err
		}
	}

	return nil
}
