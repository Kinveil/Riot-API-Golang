package region

import (
	"fmt"
	"strings"

	"github.com/junioryono/Riot-API-Golang/constants/continent"
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

// Returns the full hostname corresponding to the region.
func (r Region) Host() string {
	switch r {
	case BR1:
		return "https://br1.api.riotgames.com"
	case EUN1:
		return "https://eun1.api.riotgames.com"
	case EUW1:
		return "https://euw1.api.riotgames.com"
	case JP1:
		return "https://jp1.api.riotgames.com"
	case KR:
		return "https://kr.api.riotgames.com"
	case LA1:
		return "https://la1.api.riotgames.com"
	case LA2:
		return "https://la2.api.riotgames.com"
	case NA1:
		return "https://na1.api.riotgames.com"
	case OC1:
		return "https://oc1.api.riotgames.com"
	case PH2:
		return "https://ph2.api.riotgames.com"
	case RU:
		return "https://ru.api.riotgames.com"
	case SG2:
		return "https://sg2.api.riotgames.com"
	case TH2:
		return "https://th2.api.riotgames.com"
	case TR1:
		return "https://tr1.api.riotgames.com"
	case TW2:
		return "https://tw2.api.riotgames.com"
	case VN2:
		return "https://vn2.api.riotgames.com"
	default:
		panic(fmt.Sprintf("region %s does not have a configured host", r))
	}
}

// Returns the continent that the region is in. (Match v5)
func (r Region) Continent() continent.Continent {
	switch r {
	case BR1:
		return continent.AMERICAS
	case EUN1:
		return continent.EUROPE
	case EUW1:
		return continent.EUROPE
	case JP1:
		return continent.ASIA
	case KR:
		return continent.ASIA
	case LA1:
		return continent.AMERICAS
	case LA2:
		return continent.AMERICAS
	case NA1:
		return continent.AMERICAS
	case OC1:
		return continent.SEA
	case PH2:
		return continent.SEA
	case RU:
		return continent.EUROPE
	case SG2:
		return continent.SEA
	case TH2:
		return continent.SEA
	case TR1:
		return continent.EUROPE
	case TW2:
		return continent.ASIA
	case VN2:
		return continent.SEA
	default:
		panic(fmt.Sprintf("region %s does not have a configured host", r))
	}
}

func (r Region) String() string {
	return string(r)
}

func FormatString(rgn string) Region {
	rgn = strings.ToUpper(rgn)
	switch rgn {
	case "BR":
		return "BR1"
	case "EUN":
		fallthrough
	case "EUNE":
		return "EUN1"
	case "EUW":
		return "EUW1"
	case "JP":
		return "JP1"
	case "KR":
		return "KR"
	case "LAN":
		return "LA1"
	case "LAS":
		return "LA2"
	case "NA":
		return "NA1"
	case "OCE":
		return "OC1"
	case "RU":
		return "RU"
	case "TR":
		return "TR1"
	default:
		panic(fmt.Sprintf("region %s is invalid", rgn))
	}
}
