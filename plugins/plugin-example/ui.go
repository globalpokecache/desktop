package example

import (
	"github.com/globalpokecache/desktop/public/ifaces"
	"github.com/globalpokecache/desktop/public/models"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (m *Module) lookupWidgets() {
	// Getting input ref
	m.input = widgets.NewQTextEditFromPointer(m.widget.FindChild("input", core.Qt__FindChildrenRecursively).Pointer())
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
