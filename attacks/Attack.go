package attacks

//Attack holds information on a attack
type Attack struct {
	ID uint16

	Name        string
	Description string

	Vip  bool
	Port string
        Duration string
	API  string
}
