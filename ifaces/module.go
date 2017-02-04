package ifaces

import (
	"github.com/Sirupsen/logrus"
	"github.com/globalpokecache/desktop/public/models"
	"github.com/therecipe/qt/widgets"
)

type ModuleContext interface {
	StoragePut(string, []byte) error
	StorageGet(string) ([]byte, error)

	Config() models.Config
	Widget() *widgets.QWidget
	Log() *logrus.Entry

	AddAction(string, ActionHandler)
	DoAction(string, ...interface{}) (interface{}, error)
	ActionArguments(string, int, interface{})
}

type Module interface {
	Title() string
	Namespace() string
	UI() string

	StartAPI(ModuleContext)
	StartUI(ModuleContext)
	OnFocus(ModuleContext)
	BeforeExit(ModuleContext)
}
