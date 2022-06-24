package main

import (
	"math"
)

var columnID int = 1
var floorRequestButtonID int = 1
var floor int = 1

type Battery struct {
	ID int
	status string
	columnsList []Column
	floorRequestButtonsList []FloorRequestButton
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	b := new(Battery)
	b.ID = _id
	b.status = "online"

	if _amountOfBasements > 0 {
		b.createBasementFloorRequestButtons(_amountOfBasements)
		b.createBasmentColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	b.createFloorRequestButtons(_amountOfFloors)
	b.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)

	return b
}

//Create the column for basement floor (floor < 0)
func (b *Battery) createBasmentColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	var servedFloors []int
	floor = -1

	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	column := NewColumn(columnID, "online", _amountOfBasements, _amountOfElevatorPerColumn, servedFloors, true)
	b.columnsList = append(b.columnsList, *column)
	columnID++
}

//Create floor that are greater then 0
func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevatorPerColumn int) {
	var amountOfFloorsPerColumn = int(math.Round(float64(_amountOfFloors / _amountOfColumns)))
	floor = 1

	for i := 0; i < _amountOfColumns; i++ {
		var servedFloors []int
		for j := 0; j < amountOfFloorsPerColumn; j++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := NewColumn(columnID, "online", _amountOfFloors, _amountOfElevatorPerColumn, servedFloors, false)
		b.columnsList = append(b.columnsList, *column)
		columnID++
	}
}

//Create floor request button
func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	var buttonFloor = 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(floorRequestButtonID, "OFF", buttonFloor, "up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
}

//create floor request button for basement
func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	var buttonFloor = -1
	for i := 0; i < amountOfBasements; i++ {
		basementFloorRequestButton := NewFloorRequestButton(floorRequestButtonID, "OFF", buttonFloor, "down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, *basementFloorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
}

//Contain element
func containsElement(requestedFloor int, servedFloorsList []int) bool {
	for _, floor := range servedFloorsList {
		if floor == requestedFloor {
			return true
		}
	}
	return false
}

//Choose the best column
func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, column := range b.columnsList {
		if containsElement(_requestedFloor, column.servedFloorsList) {
			return &column
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := *b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return &column, elevator
}
