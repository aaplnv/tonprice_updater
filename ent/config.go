// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	AEDQuote  []ent.Hook
	ARSQuote  []ent.Hook
	AUDQuote  []ent.Hook
	BHDQuote  []ent.Hook
	BRLQuote  []ent.Hook
	BTCQuote  []ent.Hook
	CADQuote  []ent.Hook
	CHFQuote  []ent.Hook
	CLPQuote  []ent.Hook
	CNYQuote  []ent.Hook
	CZKQuote  []ent.Hook
	EUROQuote []ent.Hook
	GBPQuote  []ent.Hook
	HKDQuote  []ent.Hook
	HUFQuote  []ent.Hook
	IDRQuote  []ent.Hook
	ILSQuote  []ent.Hook
	INRQuote  []ent.Hook
	JPYQuote  []ent.Hook
	MXNQuote  []ent.Hook
	NOKQuote  []ent.Hook
	NZDQuote  []ent.Hook
	PKRQuote  []ent.Hook
	PLNQuote  []ent.Hook
	RUBQuote  []ent.Hook
	SARQuote  []ent.Hook
	SEKQuote  []ent.Hook
	TRYQuote  []ent.Hook
	TWDQuote  []ent.Hook
	UAHQuote  []ent.Hook
	USDQuote  []ent.Hook
	ZARQuote  []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
