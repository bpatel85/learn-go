package elevator

type Building struct {
	MinFloor int
	MaxFloor int

	Elevators []*Elevator
	Persons   []*Person
}
