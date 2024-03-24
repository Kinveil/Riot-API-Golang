package summoner_spell

import (
	"fmt"
	"io"
	"encoding/json"
)

type ID int
type String string
type PrettyString string

const (
	SummonerBoost                    ID = 1
	SummonerExhaust                  ID = 3
	SummonerFlash                    ID = 4
	SummonerHaste                    ID = 6
	SummonerHeal                     ID = 7
	SummonerSmite                    ID = 11
	SummonerTeleport                 ID = 12
	SummonerMana                     ID = 13
	SummonerDot                      ID = 14
	SummonerBarrier                  ID = 21
	SummonerPoroRecall               ID = 30
	SummonerPoroThrow                ID = 31
	SummonerSnowball                 ID = 32
	SummonerSnowURFSnowball_Mark     ID = 39
	SummonerSnowball_Mark            ID = 40
	SummonerSnowURF                  ID = 41
	SummonerSnowURFCheck             ID = 42
	Summoner_UltBookPlaceholder      ID = 54
	Summoner_UltBookSmitePlaceholder ID = 55
)

var stringToIDMap = map[String]ID{
	"SummonerBoost":                    SummonerBoost,
	"SummonerExhaust":                  SummonerExhaust,
	"SummonerFlash":                    SummonerFlash,
	"SummonerHaste":                    SummonerHaste,
	"SummonerHeal":                     SummonerHeal,
	"SummonerSmite":                    SummonerSmite,
	"SummonerTeleport":                 SummonerTeleport,
	"SummonerMana":                     SummonerMana,
	"SummonerDot":                      SummonerDot,
	"SummonerBarrier":                  SummonerBarrier,
	"SummonerPoroRecall":               SummonerPoroRecall,
	"SummonerPoroThrow":                SummonerPoroThrow,
	"SummonerSnowball":                 SummonerSnowball,
	"SummonerSnowURFSnowball_Mark":     SummonerSnowURFSnowball_Mark,
	"SummonerSnowball_Mark":            SummonerSnowball_Mark,
	"SummonerSnowURF":                  SummonerSnowURF,
	"SummonerSnowURFCheck":             SummonerSnowURFCheck,
	"Summoner_UltBookPlaceholder":      Summoner_UltBookPlaceholder,
	"Summoner_UltBookSmitePlaceholder": Summoner_UltBookSmitePlaceholder,
}

var idToStringMap = map[ID]String{
	SummonerBoost:                    "SummonerBoost",
	SummonerExhaust:                  "SummonerExhaust",
	SummonerFlash:                    "SummonerFlash",
	SummonerHaste:                    "SummonerHaste",
	SummonerHeal:                     "SummonerHeal",
	SummonerSmite:                    "SummonerSmite",
	SummonerTeleport:                 "SummonerTeleport",
	SummonerMana:                     "SummonerMana",
	SummonerDot:                      "SummonerDot",
	SummonerBarrier:                  "SummonerBarrier",
	SummonerPoroRecall:               "SummonerPoroRecall",
	SummonerPoroThrow:                "SummonerPoroThrow",
	SummonerSnowball:                 "SummonerSnowball",
	SummonerSnowURFSnowball_Mark:     "SummonerSnowURFSnowball_Mark",
	SummonerSnowball_Mark:            "SummonerSnowball_Mark",
	SummonerSnowURF:                  "SummonerSnowURF",
	SummonerSnowURFCheck:             "SummonerSnowURFCheck",
	Summoner_UltBookPlaceholder:      "Summoner_UltBookPlaceholder",
	Summoner_UltBookSmitePlaceholder: "Summoner_UltBookSmitePlaceholder",
}

func (s ID) String() String {
	return idToStringMap[s]
}

func (s String) ID() ID {
	return stringToIDMap[s]
}

func (s *ID) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch id := v.(type) {
	case float64:
		*s = ID(id)
	case string:
		*s = String(id).ID()
	default:
		return fmt.Errorf("invalid spell id: %v", id)
	}

	return nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (s *ID) UnmarshalGQL(v interface{}) error {
	intValue, ok := v.(int)
	if !ok {
		return fmt.Errorf("rank must be an int")
	}

	*s = ID(intValue)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (s ID) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, s)
}
