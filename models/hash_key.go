package models

type HashKey struct {
	Key string
	HashKeyStatus
}

type HashKeyStatus struct {
	Active bool
}
