package core

import "fmt"

type OnInitFunc func(p IProvider) error

type IProvider interface {
	GetName() string
	GetModule() IModule
	AssignModule(m IModule)
	OnInit() error
	// OnInit(onInit OnInitFunc)
	// Init() error
}

type Provider struct {
	name   string
	module IModule
	// onInit OnInitFunc
}

func NewProvider(name string) *Provider {
	return &Provider{
		name: name,
	}
}

func (p *Provider) GetName() string {
	return p.name
}

func (p *Provider) GetModule() IModule {
	return p.module
}

func (p *Provider) AssignModule(m IModule) {
	p.module = m
}

// func (p *Provider) OnInit(onInit OnInitFunc) {
// 	p.onInit = onInit
// }

func (p *Provider) OnInit() error {
	fmt.Printf("Initializing %s\n", p.GetName())
	return nil
}

// func (p *Provider) Init() error {
// 	// Prevent circular dependency
// 	fmt.Printf("Initializing %s\n", p.GetName())
// 	return p.OnInit()
// }

// func (p *Provider) Init() error {
// 	fmt.Printf("Initializing %s\n", p.GetName())
// 	if p.onInit == nil {
// 		return nil
// 	}
// 	return p.onInit(p)
// 	// return p.OnInit()
// }
