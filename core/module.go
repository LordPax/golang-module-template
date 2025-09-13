package core

import "fmt"

var moduleSeen []string

// IModule defines the interface for a module that can manage providers and dependencies.
type IModule interface {
	GetName() string
	Get(provider string) IProvider
	AddProvider(p IProvider)
	AddModule(dep IModule)
	Init() error
	getProvider(provider string) IProvider
}

// Module is a struct that implements the IModule interface, managing providers and dependencies.
type Module struct {
	name     string
	module   map[string]IModule
	provider map[string]IProvider
}

// NewModule creates a new instance of Module with the given name.
func NewModule(name string) *Module {
	return &Module{
		name:     name,
		module:   make(map[string]IModule),
		provider: make(map[string]IProvider),
	}
}

// GetName returns the name of the module.
func (m *Module) GetName() string {
	return m.name
}

// Get retrieves a provider by its name, searching within the module and its dependencies.
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

// getProvider is a helper method that retrieves a provider by its name within the current module.
func (m *Module) getProvider(provider string) IProvider {
	return m.provider[provider]
}

// AddProvider adds a provider to the module and assigns the module to the provider.
func (m *Module) AddProvider(p IProvider) {
	m.provider[p.GetName()] = p
	p.AssignModule(m)
}

// AddModule adds a dependency module to the current module.
func (m *Module) AddModule(dep IModule) {
	m.module[dep.GetName()] = dep
}

// Init initializes the module and its dependencies, ensuring each module is initialized only once.
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

	fmt.Printf("==> Initializing %s\n", m.GetName())

	for _, p := range m.provider {
		if err := p.OnInit(); err != nil {
			return err
		}
	}

	return nil
}
