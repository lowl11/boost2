package storage

import "time"

type ConnectionBuilder struct {
	maxOpenConnections int
	maxIdleConnections int

	maxConnectionLifetime *time.Duration
	maxIdleLifetime       *time.Duration

	ignoreSchema bool
}

func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{}
}

func (builder *ConnectionBuilder) SetMaxOpenConnections(maxOpenConnections int) *ConnectionBuilder {
	builder.maxOpenConnections = maxOpenConnections
	return builder
}

func (builder *ConnectionBuilder) SetMaxIdleConnections(maxIdleConnections int) *ConnectionBuilder {
	builder.maxIdleConnections = maxIdleConnections
	return builder
}

func (builder *ConnectionBuilder) SetMaxConnectionLifetime(lifetime time.Duration) *ConnectionBuilder {
	builder.maxConnectionLifetime = &lifetime
	return builder
}

func (builder *ConnectionBuilder) getMaxConnectionLifetime() time.Duration {
	if builder.maxConnectionLifetime != nil {
		return *builder.maxConnectionLifetime
	}

	return time.Second * 10
}

func (builder *ConnectionBuilder) SetMaxIdleLifetime(lifetime time.Duration) *ConnectionBuilder {
	builder.maxIdleLifetime = &lifetime
	return builder
}

func (builder *ConnectionBuilder) IgnoreSchema() *ConnectionBuilder {
	builder.ignoreSchema = true
	return builder
}

func (builder *ConnectionBuilder) isIgnoreSchema() bool {
	return builder.ignoreSchema
}

func (builder *ConnectionBuilder) getMaxIdleLifetime() time.Duration {
	if builder.maxIdleLifetime != nil {
		return *builder.maxIdleLifetime
	}

	return time.Second * 10
}
