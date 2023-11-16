package di

import (
	"context"
	"fmt"
	"sync"
)

// Container represents a dependency injection container.
type Container struct {
	context context.Context
	mu      sync.Mutex
}

// NewContainer creates a new dependency injection container.
func NewContainer() *Container {
	return &Container{
		context: context.Background(),
	}
}

type Key string

// WithValue adds a key-value pair to the container's context.
func (c *Container) WithValue(key Key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.context = context.WithValue(c.context, key, value)
}

// Resolve retrieves a value from the container's context based on the provided key.
func (c *Container) Resolve(key Key) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value := c.context.Value(key)
	if value == nil {
		return nil, fmt.Errorf("invalid DI container key: %s", key)
	}

	return value, nil
}
