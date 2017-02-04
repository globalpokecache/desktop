package ifaces

type ActionHandler func(...interface{}) (interface{}, error)
