package statemachinesystem

import "fmt"

type statemachine struct {
	stateCurrent  state
	statePrevious state
	statePossible []state
}

type state struct {
	id   int
	name string
}

var sm statemachine

func Initialise() {

	firstState := state{id: 0, name: "one"}
	secondState := state{id: 1, name: "two"}
	thirdState := state{id: 2, name: "three"}

	sm.stateCurrent = firstState
	sm.statePrevious = firstState
	sm.statePossible = []state{firstState, secondState, thirdState}

	fmt.Printf("Initialise as FIRST : %v   PREV : %v \n", sm.stateCurrent, sm.statePrevious)

}

func GetCurrentState() state {

	return sm.stateCurrent
}

func ChangeCurrentState() {

	for index, el := range sm.statePossible {
		if el.id == sm.stateCurrent.id {
			sm.statePrevious = sm.stateCurrent
			// sm.stateCurrent = sm.statePossible[index%len(sm.statePossible)]

			sm.stateCurrent = sm.statePossible[index+1]

			fmt.Printf("Index: %v and current State is now %v", index, sm.stateCurrent)
			break
		}
	}
}
