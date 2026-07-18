package game

type Game struct {
	Board []Space `json:"board"`
	Rolls []Die   `json:"rolls"`
}

type Space struct {
	ID             int     `json:"id"`
	SpectatorValue int     `json:"spectatorValue"`
	Camels         []Camel `json:"camels"`
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
	Color DieColor `json:"color"`
	Value int      `json:"value"`
}

type DieColor string

const (
	BlueDie   DieColor = "blue"
	GrayDie   DieColor = "gray"
	GreenDie  DieColor = "green"
	PurpleDie DieColor = "purple"
	RedDie    DieColor = "red"
	YellowDie DieColor = "yellow"
)
