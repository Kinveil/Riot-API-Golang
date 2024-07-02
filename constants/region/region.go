package region

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Kinveil/Riot-API-Golang/constants/continent"
)

type Region string

const (
	// Brazil
	BR1 Region = "BR1"

	// Europe East
	EUN1 Region = "EUN1"

	// Europe West
	EUW1 Region = "EUW1"

	// Japan
	JP1 Region = "JP1"

	// Korea
	KR Region = "KR"

	// Latin America North
	LA1 Region = "LA1"

	// Latin America South
	LA2 Region = "LA2"

	// North America
	NA1 Region = "NA1"

	// Oceania
	OC1 Region = "OC1"

	// Philippines
	PH2 Region = "PH2"

	// Russia
	RU Region = "RU"

	// Singapore
	SG2 Region = "SG2"

	// Thailand
	TH2 Region = "TH2"

	// Turkey
	TR1 Region = "TR1"

	// Taiwan
	TW2 Region = "TW2"

	// Vietnam
	VN2 Region = "VN2"
)

func (r Region) String() string {
	return string(r)
}

var stringToRegion = map[string]Region{
	"BR":  BR1,
	"EUN": EUN1,
	"EUW": EUW1,
	"JP":  JP1,
	"KR":  KR,
	"LA1": LA1,
	"LAN": LA1,
	"LA2": LA2,
	"LAS": LA2,
	"NA":  NA1,
	"OC":  OC1,
	"PH":  PH2,
	"RU":  RU,
	"SG":  SG2,
	"TH":  TH2,
	"TR":  TR1,
	"TW":  TW2,
	"VN":  VN2,
}

func FromString(rgn string) Region {
	// Capitalize the string
	rgn = strings.ToUpper(rgn)

	// Special handling for Latin America regions
	if strings.HasPrefix(rgn, "LA") {
		if region, ok := stringToRegion[rgn]; ok {
			return region
		}
	} else {
		// Remove all numbers from the string for other regions
		rgn = strings.Map(func(r rune) rune {
			if unicode.IsNumber(r) {
				return -1
			}
			return r
		}, rgn)
	}

	if region, ok := stringToRegion[rgn]; ok {
		return region
	}

	panic(fmt.Sprintf("region %s is invalid", rgn))
}

var regionToHost = map[Region]string{
	BR1:  "https://br1.api.riotgames.com",
	EUN1: "https://eun1.api.riotgames.com",
	EUW1: "https://euw1.api.riotgames.com",
	JP1:  "https://jp1.api.riotgames.com",
	KR:   "https://kr.api.riotgames.com",
	LA1:  "https://la1.api.riotgames.com",
	LA2:  "https://la2.api.riotgames.com",
	NA1:  "https://na1.api.riotgames.com",
	OC1:  "https://oc1.api.riotgames.com",
	PH2:  "https://ph2.api.riotgames.com",
	RU:   "https://ru.api.riotgames.com",
	SG2:  "https://sg2.api.riotgames.com",
	TH2:  "https://th2.api.riotgames.com",
	TR1:  "https://tr1.api.riotgames.com",
	TW2:  "https://tw2.api.riotgames.com",
	VN2:  "https://vn2.api.riotgames.com",
}

// Returns the full hostname corresponding to the region.
func (r Region) Host() string {
	if host, ok := regionToHost[r]; ok {
		return host
	}

	panic(fmt.Sprintf("region %s does not have a configured host", r))
}

var regionToContinentMatchV5 = map[Region]continent.Continent{
	BR1:  continent.AMERICAS,
	EUN1: continent.EUROPE,
	EUW1: continent.EUROPE,
	JP1:  continent.ASIA,
	KR:   continent.ASIA,
	LA1:  continent.AMERICAS,
	LA2:  continent.AMERICAS,
	NA1:  continent.AMERICAS,
	OC1:  continent.SEA,
	PH2:  continent.SEA,
	RU:   continent.EUROPE,
	SG2:  continent.SEA,
	TH2:  continent.SEA,
	TR1:  continent.EUROPE,
	TW2:  continent.ASIA,
	VN2:  continent.SEA,
}

// Returns the continent that the region is in
func (r Region) ContinentMatchV5() continent.Continent {
	if continent, ok := regionToContinentMatchV5[r]; ok {
		return continent
	}

	panic(fmt.Sprintf("region %s does not have a configured continent", r))
}

// Map the nearest region to the continent
var regionToContinentAccountV1 = map[Region]continent.Continent{
	BR1:  continent.AMERICAS,
	EUN1: continent.EUROPE,
	EUW1: continent.EUROPE,
	JP1:  continent.ASIA,
	KR:   continent.ASIA,
	LA1:  continent.AMERICAS,
	LA2:  continent.AMERICAS,
	NA1:  continent.AMERICAS,
	OC1:  continent.AMERICAS,
	PH2:  continent.AMERICAS,
	RU:   continent.ASIA,
	SG2:  continent.ASIA,
	TH2:  continent.ASIA,
	TR1:  continent.EUROPE,
	TW2:  continent.ASIA,
	VN2:  continent.ASIA,
}

// Returns the continent that the region is in
func (r Region) ContinentAccountV1() continent.Continent {
	if continent, ok := regionToContinentAccountV1[r]; ok {
		return continent
	}

	panic(fmt.Sprintf("region %s does not have a configured continent", r))
}
