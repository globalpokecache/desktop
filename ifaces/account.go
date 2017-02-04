package ifaces

import (
	"github.com/muxgo/desktop/public/models"
)

type AccountControl interface {
	Add(*models.Account) error
	Remove(*models.Account) error
}
