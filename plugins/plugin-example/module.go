package example

import (
	"fmt"
	"github.com/globalpokecache/desktop/public/ifaces"
	"github.com/therecipe/qt/widgets"
)

func NewModule() ifaces.Module {
	return &Module{}
}

type Module struct {
	widget *widgets.QWidget
	moduleWidgets
}

func (m *Module) Title() string {
	return "Example"
}

func (m *Module) Namespace() string {
	return "example"
}

func (m *Module) OnFocus(c ifaces.ModuleContext) {
	// This method is called every time this plugin's tab is focused

	// Let me give an example of how to use local plugin API
	c.DoAction("example.save-info", []byte("Hello world"))
	iinfo, _ := c.DoAction("example.get-info", []byte("Hello world"))
	info := iinfo.([]byte)

	// Now lets change a UI widget using the info from the plugin API
	m.input.SetText(string(info))
}

func (m *Module) BeforeExit(c ifaces.ModuleContext) {
	// This method is called before the program exists
	fmt.Println("bye")
}
