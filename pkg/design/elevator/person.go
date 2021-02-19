package elevator

type Person struct {
	Name            string
	Source          int
	Destination     int
	CurrentElevator *Elevator
}
