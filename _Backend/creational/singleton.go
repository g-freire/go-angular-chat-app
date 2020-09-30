//package main
package tee

import (
	"fmt"
	"sync"
)

type singleton struct{
	count int
}
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func main() {
	var a = GetInstance()
	a.AddOne()
	a.AddOne()
	a.AddOne()
	fmt.Print(*a)
}


//PROTECTING SINGLETON
var once sync.Once

// type global
type singleton map[string]string

var (
	instance singleton
)

//init before main would also be a singleton, but not best approach - security, flexibility , etc
//func init() {
////	Connect()
//}

func NewClass() singleton {
	once.Do(func() { // <-- atomic, does not allow repeating
		instance = make(singleton) // <-- thread safe
	})

	//or
	//lock.Lock()
	//defer lock.Unlock()

	return instance
}