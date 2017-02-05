package example

import (
	"github.com/globalpokecache/desktop/public/ifaces"
)

func (m *Module) StartAPI(c ifaces.ModuleContext) {

	// Adding the example action 'save-info' to this module API
	// This action can be used by this or other plugins
	// See module.go to check how to use it
	// If the action name first letter is uppercase means it should be available
	// in the HTTP API
	c.AddAction("SaveInfo", func(args ...interface{}) (interface{}, error) {
		c.StoragePut("info", args[0].([]byte))
		return nil, nil
	})
	// This method defines the input type for the Web API
	// Example: here we are saying it should receive two strings as argument
	c.ActionArguments("save-info", 2, string(""))

	c.AddAction("GetInfo", func(args ...interface{}) (interface{}, error) {
		return c.StorageGet("info")
	})
	c.ActionArguments("get-info", 1, new(string))

	// This is a native only action
	c.AddAction("native-action", func(args ...interface{}) (interface{}, error) {
		return nil, nil
	})
}
