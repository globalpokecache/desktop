package example

import (
	"fmt"
	"github.com/globalpokecache/desktop/public/ifaces"
	"github.com/globalpokecache/desktop/public/models"
	"github.com/therecipe/qt/widgets"
)

func NewModule() ifaces.Module {
	return &Module{}
}

type Module struct {
	widget *widgets.QWidget
	input  *widgets.QLineEdit
}

func (m *Module) Title() string {
	return "Example"
}

func (m *Module) Namespace() string {
	return "example"
}

func (m *Module) Start(c ifaces.ModuleContext) {
	// This methods is called when the plugin is started and the UI is ready

	// Save plugin's widget reference
	m.widget = c.Widget()

	// Lookup for widgets in the UI
	m.lookupWidgets()

	// Let me show how to use other plugins API
	// Lets fetch all accounts available in the account module
	iaccs, _ := c.DoAction("accounts.Fetch")
	if iaccs == nil {
		// Failed to load accounts
	} else {
		// Laoded accounts list
		// But we need to identify the interface{}
		accs := iaccs.([]*models.Account)
		// Now do what you want with accounts
		fmt.Println(accs)
	}
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
