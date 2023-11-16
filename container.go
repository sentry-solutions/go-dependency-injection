package di

import (
	"context"
	"fmt"
	"sync"

	commons "github.com/sentry-solutions/sentry-go-commons"
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

// WithValue adds a key-value pair to the container's context.
func (c *Container) WithValue(key interface{}, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.context = context.WithValue(c.context, key, value)
}

// Resolve retrieves a value from the container's context based on the provided key.
func (c *Container) Resolve(key interface{}) (interface{}, commons.LogMessage) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value := c.context.Value(key)
	if value == nil {
		return nil, commons.LogMessage{
			Severity:    commons.LOG_SEVERITY_ERROR,
			Description: fmt.Sprintf("Invalid DI container Key: %v", key),
			StackTrace:  nil,
		}
	}

	return value, commons.LogMessage{
		Severity:    commons.LOG_SEVERITY_ERROR,
		Description: "Resolved valid DI keypair",
		StackTrace:  nil,
	}
}
