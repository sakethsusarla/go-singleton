package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// Singleton is a struct that will have only one instance
type Singleton struct {
}

var instance *Singleton

// GetInstance is a function that will return the instance of Singleton
func GetInstance() *Singleton {
	if instance == nil {
		once.Do(
			func() {
				instance = &Singleton{}
				fmt.Println("Created instance of Singleton")
			},
		)
	} else {
		fmt.Println("Instance of Singleton already exists")
	}

	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetInstance()
	}

	fmt.Scanln()
}
