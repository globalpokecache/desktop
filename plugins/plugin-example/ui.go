package example

import (
	"fmt"
	"github.com/globalpokecache/desktop/public/ifaces"
	"github.com/globalpokecache/desktop/public/models"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type moduleWidgets struct {
	input *widgets.QLineEdit
}

func (m *Module) StartUI(c ifaces.ModuleContext) {
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

func (m *Module) lookupWidgets() {
	// Getting input ref
	m.input = widgets.NewQLineEditFromPointer(m.widget.FindChild("input", core.Qt__FindChildrenRecursively).Pointer())
}

func (m *Module) UI() string {
	// Qt UI XML string
	return `<?xml version="1.0" encoding="UTF-8"?>
<ui version="4.0">
 <class>frame</class>
 <widget class="QFrame" name="frame">
    <widget class="QLineEdit" name="input"/>
 </widget>
 <resources/>
 <connections/>
</ui>`
}
