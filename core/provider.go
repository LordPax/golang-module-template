package core

import "fmt"

type OnInitFunc func(p IProvider) error

type IProvider interface {
	GetName() string
	GetModule() IModule
	AssignModule(m IModule)
	OnInit() error
}

type Provider struct {
	name   string
	module IModule
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

func (p *Provider) OnInit() error {
	fmt.Printf("Initializing %s\n", p.GetName())
	return nil
}
