package main

import (
	"sort"
)

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
	direction             string
}

func NewElevator(_id int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	e := new(Elevator)
	e.ID = _id
	e.status = _status
	e.floorRequestsList = make([]int, 0)
	e.completedRequestsList = make([]int, 0)
	e.amountOfFloors = _amountOfFloors
	e.currentFloor = _currentFloor
	e.door = Door{1, ""}
	e.direction = ""

	return e
}

//Make elevator move to desire direction
func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		var destination int = e.floorRequestsList[0]
		e.status = "moving"
		e.sortFloorList()
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		if !contains(e.completedRequestsList, destination) {
			e.completedRequestsList = append(e.completedRequestsList, destination)
		}
		e.floorRequestsList = e.floorRequestsList[1:]

	}
	e.status = "idle"

}

//sort floors in numerical or reverse
func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Sort(sort.IntSlice(e.floorRequestsList))
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}

//open and close doors
func (e *Elevator) operateDoors() {
	if e.status == "stopped" || e.status == "idle" {
		e.door.status = "open"

		if len(e.floorRequestsList) < 1 {
			e.direction = ""
			e.status = "idle"
		}
	}
}

//add requested floor in a request list
func (e *Elevator) addNewRequest(userPosition int) {
	if !contains(e.floorRequestsList, userPosition) {
		e.floorRequestsList = append([]int{userPosition}, e.floorRequestsList...)
	}
	if e.currentFloor < userPosition {
		e.direction = "up"
	}
	if e.currentFloor > userPosition {
		e.direction = "down"
	}
}
