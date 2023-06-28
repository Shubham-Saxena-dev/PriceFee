package registry

import (
	"fmt"
	"sync"
)

var (
	iocContainer Container
	once         sync.Once
)

type ResolutionFactory func(container Container) (interface{}, error)

type Container interface {
	RegisterFunc(interface{}, ResolutionFactory) error
	RegisterInstance(interface{}, interface{}) error
	Resolve(interface{}) interface{}
}

type ResolutionLifetime int

const (
	ResolutionLifetimePerCall ResolutionLifetime = iota
	ResolutionLifetimeSingleton
)

type simpleIocContainer struct {
	registrations map[interface{}]*containerRegistration
	mutex         sync.RWMutex
}

type containerRegistration struct {
	name               interface{}
	registrationType   registrationType
	instance           interface{}
	factoryFunc        ResolutionFactory
	lifetimeManagement ResolutionLifetime
}

type registrationType int

const (
	registrationTypeFunc registrationType = iota
	registrationTypeInstance
)

func NewSimpleIocContainer() Container {
	once.Do(func() {
		iocContainer = &simpleIocContainer{
			registrations: make(map[interface{}]*containerRegistration),
		}
	})
	return iocContainer
}

func (c *simpleIocContainer) RegisterFunc(name interface{}, factory ResolutionFactory) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, present := c.registrations[name]; present {
		return nil
	}

	data := &containerRegistration{
		name:               name,
		registrationType:   registrationTypeFunc,
		factoryFunc:        factory,
		lifetimeManagement: ResolutionLifetimePerCall,
	}
	c.registrations[name] = data
	return nil
}

func (c *simpleIocContainer) RegisterInstance(name interface{}, instance interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, present := c.registrations[name]; present {
		return nil
	}

	data := &containerRegistration{
		name:               name,
		registrationType:   registrationTypeInstance,
		instance:           instance,
		lifetimeManagement: ResolutionLifetimeSingleton,
	}
	c.registrations[name] = data
	return nil
}

func (c *simpleIocContainer) Resolve(name interface{}) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var registration *containerRegistration
	var present bool
	if registration, present = c.registrations[name]; !present {
		panic(fmt.Errorf(`there are no items registered in the container with the name "%s"`, name))
	}

	switch registration.registrationType {
	case registrationTypeInstance:
		return registration.instance

	case registrationTypeFunc:
		return c.resolveFromFunc(registration)

	default:
		panic(fmt.Errorf(`unexpected registration type "%v"`, registration.registrationType))
	}
}

func (c *simpleIocContainer) resolveFromFunc(registration *containerRegistration) interface{} {
	if registration.lifetimeManagement == ResolutionLifetimePerCall {
		return c.invokeResolutionFactory(registration.factoryFunc)
	}
	panic(fmt.Errorf(`unexpected registration lifetime "%v"`, registration.lifetimeManagement))
}

func (c *simpleIocContainer) invokeResolutionFactory(factory ResolutionFactory) interface{} {
	instance, err := factory(c)
	if err != nil {
		panic(err)
	}
	return instance
}

func GetContainer() Container {
	return iocContainer
}
