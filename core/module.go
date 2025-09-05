package core

import "fmt"

var moduleSeen []string

type IModule interface {
	GetName() string
	// Get(module, provider string) IProvider
	Get(provider string) IProvider
	GetProvider(name string) IProvider
	AddProvider(p IProvider)
	GetModule(name string) IModule
	AddModule(dep IModule)
	Init() error
}

type Module struct {
	name    string
	depend  map[string]IModule
	provide map[string]IProvider
}

func NewModule(name string) *Module {
	return &Module{
		name:    name,
		depend:  make(map[string]IModule),
		provide: make(map[string]IProvider),
	}
}

func (m *Module) GetName() string {
	return m.name
}

func (m *Module) Get(provider string) IProvider {
	if p := m.GetProvider(provider); p != nil {
		return p
	}

	for _, d := range m.depend {
		if p := d.Get(provider); p != nil {
			return p
		}
	}

	return nil
}

func (m *Module) GetProvider(name string) IProvider {
	return m.provide[name]
	// if !ok {
	// 	return nil, fmt.Errorf("provider %s not found in module %s", name, m.name)
	// }
	// return provide, nil
}

func (m *Module) AddProvider(p IProvider) {
	m.provide[p.GetName()] = p
	p.AssignModule(m)
}

func (m *Module) GetModule(name string) IModule {
	return m.depend[name]
	// if !ok {
	// 	return nil, fmt.Errorf("dependency %s not found in module %s", name, m.name)
	// }
	// return deps, nil
}

func (m *Module) AddModule(dep IModule) {
	m.depend[dep.GetName()] = dep
}

func (m *Module) Init() error {
	for _, d := range m.depend {
		if inArray(d.GetName(), moduleSeen) {
			continue
		}

		moduleSeen = append(moduleSeen, d.GetName())

		if err := d.Init(); err != nil {
			return err
		}
	}

	fmt.Printf("Initializing %s\n", m.GetName())

	for _, p := range m.provide {
		if err := p.OnInit(); err != nil {
			return err
		}
	}

	return nil
}
