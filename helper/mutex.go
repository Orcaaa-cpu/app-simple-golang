package helper

import "sync"

var (
	Lock = sync.Mutex{}
)
