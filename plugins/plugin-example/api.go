package example

import (
	"github.com/muxgo/desktop/public/ifaces"
)

func (m *Module) API(c ifaces.ModuleContext) {

	// Adding the example action 'save-info' to this module API
	// This action can be used by this or other plugins
	// See module.go to check how to use it
	c.AddAction("save-info", func(args ...interface{}) (interface{}, error) {
		c.StoragePut("info", args[0].([]byte))
		return nil, nil
	})
	c.ActionArguments("save-info", new(string), new(string))

	c.AddAction("get-info", func(args ...interface{}) (interface{}, error) {
		info := c.StorageGet("info")
		return info, nil
	})
	c.ActionArguments("get-info", new(string))
}
