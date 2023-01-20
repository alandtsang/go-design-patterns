package singleton

import (
	"fmt"
	"sync"
)

var (
	once           sync.Once
	singleInstance *Single
)

// Single is a singleton structure.
type Single struct {
}

// New creates the single instance and return it when the first time the function is called.
// The second time the function is called, the single instance is returned.
func New() *Single {
	once.Do(func() {
		fmt.Println("Creating single instance")
		singleInstance = &Single{}
	})

	return singleInstance
}
