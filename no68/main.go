package main

import (
	"fmt"
	"sync"
)

type partyVisitor struct {
	mtx  sync.RWMutex
	name string
	age  int
}

func (pv *partyVisitor) isAllowedEntrance(hasBday bool) error {
	pv.mtx.Lock()
	defer pv.mtx.Unlock()
	if hasBday {
		pv.age++
	}

	if pv.age < 21 {
		// causes a deadlock, since '%v' calles '(pv *partyVisitor) String()', which requires the mutex, which is hold by this method itself.
		return fmt.Errorf("%v: is not of age", pv)
	}

	return nil
}

func (pv *partyVisitor) String() string {
	pv.mtx.RLock()
	defer pv.mtx.RUnlock()
	return fmt.Sprintf("name: %s, age: %d", pv.name, pv.age)
}

func main() {
	partyV := partyVisitor{
		mtx:  sync.RWMutex{},
		name: "allen",
		age:  19,
	}

	print(partyV.isAllowedEntrance(true))
}
