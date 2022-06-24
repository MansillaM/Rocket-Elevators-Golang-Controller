package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID int
	status string
	floor int
	direction string
}

func NewFloorRequestButton(_id int, _status string, _floor int, _direction string) *FloorRequestButton {
	floorRequestButton := new(FloorRequestButton)
	floorRequestButton.ID = _id
	floorRequestButton.status = _status
	floorRequestButton.floor = _floor
	floorRequestButton.direction = _direction

	return floorRequestButton
}
