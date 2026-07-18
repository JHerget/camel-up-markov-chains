package game

type Game struct {
	Board []Space `json:"board"`
	Rolls []Die
}

type Space struct {
	ID             int
	SpectatorValue int
	Camels         []Camel
}

type Camel string

const (
	BlackCamel  Camel = "black"
	BlueCamel   Camel = "blue"
	GreenCamel  Camel = "green"
	PurpleCamel Camel = "purple"
	RedCamel    Camel = "red"
	WhiteCamel  Camel = "white"
	YellowCamel Camel = "yellow"
)

type Die struct {
	Color DieColor
	Value int
}

type DieColor string

const (
	BlueDice   DieColor = "blue"
	GrayDice   DieColor = "gray"
	GreenDice  DieColor = "green"
	PurpleDice DieColor = "purple"
	RedDice    DieColor = "red"
	YellowDice DieColor = "yellow"
)
