package continent

import "fmt"

type Continent string

const (
	// Serves NA, BR, LAN and LAS
	AMERICAS Continent = "AMERICAS"

	// Serves KR and JP
	ASIA Continent = "ASIA"

	// Serves EUNE, EUW, TR and RU
	EUROPE Continent = "EUROPE"

	// Serves OCE, PH2, SG2, TH2, TW2 and VN2
	SEA Continent = "SEA"
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
