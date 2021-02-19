package elevator

type Elevator struct {
	Name         string
	Direction    int
	CurrentFloor int
	Building     *Building
}
