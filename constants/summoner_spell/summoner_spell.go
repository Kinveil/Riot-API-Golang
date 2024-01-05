package summoner_spell

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
