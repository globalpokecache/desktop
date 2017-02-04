package models

type Worker struct {
	ID int64
	WorkerStatus
	AccountInUse string
}

type WorkerStatus struct {
	Active bool
}
