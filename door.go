package main

type Door struct {
	ID int
	status string
}

func NewDoor(_id int, _status string) *Door {
		door := new(Door)
		door.ID = _id
		door.status = ""
		
		return door
}
