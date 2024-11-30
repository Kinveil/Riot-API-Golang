package continent

import (
	"strings"
)

type Continent string

const (
	AMERICAS Continent = "AMERICAS"
	ASIA     Continent = "ASIA"
	EUROPE   Continent = "EUROPE"
	SEA      Continent = "SEA"
)

func (c Continent) String() string {
	return string(c)
}

var stringToContinent = map[string]Continent{
	"AMERICAS": AMERICAS,
	"ASIA":     ASIA,
	"EUROPE":   EUROPE,
	"SEA":      SEA,
}

func FromString(cntnt string) (Continent, bool) {
	cntnt = strings.ToUpper(cntnt)
	continent, ok := stringToContinent[cntnt]
	return continent, ok
}

var continentToHost = map[Continent]string{
	AMERICAS: "https://americas.api.riotgames.com",
	ASIA:     "https://asia.api.riotgames.com",
	EUROPE:   "https://europe.api.riotgames.com",
	SEA:      "https://sea.api.riotgames.com",
}

// Returns the full hostname corresponding to the region.
func (c Continent) Host() string {
	return continentToHost[c]
}
