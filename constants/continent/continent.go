package continent

import (
	"fmt"
	"strings"
)

type Continent string

const (
	AMERICAS Continent = "AMERICAS"
	ASIA     Continent = "ASIA"
	EUROPE   Continent = "EUROPE"
	SEA      Continent = "SEA"
)

// Returns the full hostname corresponding to the region.
func (c Continent) Host() string {
	switch c {
	case AMERICAS:
		return "https://americas.api.riotgames.com"
	case ASIA:
		return "https://asia.api.riotgames.com"
	case EUROPE:
		return "https://europe.api.riotgames.com"
	case SEA:
		return "https://sea.api.riotgames.com"
	default:
		panic(fmt.Sprintf("region %s does not have a configured host", c))
	}
}

func (c Continent) String() string {
	return string(c)
}

func FormatString(cntnt string) Continent {
	cntnt = strings.ToUpper(cntnt)
	switch cntnt {
	case "AMERICAS":
		return AMERICAS
	case "ASIA":
		return ASIA
	case "EUROPE":
		return EUROPE
	case "SEA":
		return SEA
	default:
		panic(fmt.Sprintf("continent %s is invalid", cntnt))
	}
}
