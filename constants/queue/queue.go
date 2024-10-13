package queue

import (
	"fmt"
	"io"
)

type ID int16
type PrettyString string

const (
	RankedSolo5x5 ID = 420
	RankedFlexSR  ID = 440
	RankedFlexTT  ID = 470

	Snowdown                      ID = 73
	Hexakill                      ID = 75
	Nemesis                       ID = 310
	BlackMarketBrawlers           ID = 313
	DefinitelyNotDominion         ID = 317
	AllRandom                     ID = 325
	NormalDraft                   ID = 400
	NormalBlind                   ID = 430
	ARAM                          ID = 450
	BloodHuntAssassin             ID = 600
	DarkStarSingularity           ID = 610
	Clash                         ID = 700
	CoopVsAI                      ID = 850
	URF                           ID = 900
	Ascension                     ID = 910
	LegendOfThePoroKing           ID = 920
	NexusSiege                    ID = 940
	DoomBotsVoting                ID = 950
	DoomBotsStandard              ID = 960
	StarGuardianInvasionNormal    ID = 980
	StarGuardianInvasionOnslaught ID = 990
	ProjectHunters                ID = 1000
	ARURF                         ID = 1010
	OneForAll                     ID = 1020
	OdysseyExtractionIntro        ID = 1030
	OdysseyExtractionCadet        ID = 1040
	OdysseyExtractionCrewmember   ID = 1050
	OdysseyExtractionCaptain      ID = 1060
	OdysseyExtractionOnslaught    ID = 1070
	NexusBlitz                    ID = 1300
	Ultimates                     ID = 1400
	Arena                         ID = 1700
	Tutorial                      ID = 2020
)

var idToPrettyStringMap = map[ID]PrettyString{
	RankedSolo5x5:                 "Ranked Solo",
	RankedFlexSR:                  "Ranked Flex",
	RankedFlexTT:                  "Ranked Flex",
	Snowdown:                      "Snowdown",
	Hexakill:                      "Hexakill",
	Nemesis:                       "Nemesis",
	BlackMarketBrawlers:           "Black Market Brawlers",
	DefinitelyNotDominion:         "Definitely Not Dominion",
	AllRandom:                     "All Random",
	NormalDraft:                   "Normal Draft",
	NormalBlind:                   "Normal Blind",
	ARAM:                          "ARAM",
	BloodHuntAssassin:             "Blood Hunt Assassin",
	DarkStarSingularity:           "Dark Star: Singularity",
	Clash:                         "Clash",
	CoopVsAI:                      "Co-op vs. AI",
	URF:                           "URF",
	Ascension:                     "Ascension",
	LegendOfThePoroKing:           "Legend of the Poro King",
	NexusSiege:                    "Nexus Siege",
	DoomBotsVoting:                "Doom Bots Voting",
	DoomBotsStandard:              "Doom Bots Standard",
	StarGuardianInvasionNormal:    "Star Guardian Invasion: Normal",
	StarGuardianInvasionOnslaught: "Star Guardian Invasion: Onslaught",
	ProjectHunters:                "PROJECT: Hunters",
	ARURF:                         "ARURF",
	OneForAll:                     "One for All",
	OdysseyExtractionIntro:        "Odyssey Extraction: Intro",
	OdysseyExtractionCadet:        "Odyssey Extraction: Cadet",
	OdysseyExtractionCrewmember:   "Odyssey Extraction: Crewmember",
	OdysseyExtractionCaptain:      "Odyssey Extraction: Captain",
	OdysseyExtractionOnslaught:    "Odyssey Extraction: Onslaught",
	NexusBlitz:                    "Nexus Blitz",
	Ultimates:                     "Ultimates",
	Arena:                         "Arena",
	Tutorial:                      "Tutorial",
}

var prettyStringToIDMap = map[string]ID{
	"Ranked Solo":                       RankedSolo5x5,
	"Ranked Flex":                       RankedFlexSR,
	"Snowdown":                          Snowdown,
	"Hexakill":                          Hexakill,
	"Nemesis":                           Nemesis,
	"Black Market Brawlers":             BlackMarketBrawlers,
	"Definitely Not Dominion":           DefinitelyNotDominion,
	"All Random":                        AllRandom,
	"Normal Draft":                      NormalDraft,
	"Normal Blind":                      NormalBlind,
	"ARAM":                              ARAM,
	"Blood Hunt Assassin":               BloodHuntAssassin,
	"Dark Star: Singularity":            DarkStarSingularity,
	"Clash":                             Clash,
	"Co-op vs. AI":                      CoopVsAI,
	"URF":                               URF,
	"Ascension":                         Ascension,
	"Legend of the Poro King":           LegendOfThePoroKing,
	"Nexus Siege":                       NexusSiege,
	"Doom Bots Voting":                  DoomBotsVoting,
	"Doom Bots Standard":                DoomBotsStandard,
	"Star Guardian Invasion: Normal":    StarGuardianInvasionNormal,
	"Star Guardian Invasion: Onslaught": StarGuardianInvasionOnslaught,
	"PROJECT: Hunters":                  ProjectHunters,
	"ARURF":                             ARURF,
	"One for All":                       OneForAll,
	"Odyssey Extraction: Intro":         OdysseyExtractionIntro,
	"Odyssey Extraction: Cadet":         OdysseyExtractionCadet,
	"Odyssey Extraction: Crewmember":    OdysseyExtractionCrewmember,
	"Odyssey Extraction: Captain":       OdysseyExtractionCaptain,
	"Odyssey Extraction: Onslaught":     OdysseyExtractionOnslaught,
	"Nexus Blitz":                       NexusBlitz,
	"Ultimates":                         Ultimates,
	"Arena":                             Arena,
	"Tutorial":                          Tutorial,
}

func (q ID) PrettyString() PrettyString {
	return idToPrettyStringMap[q]
}

func (q PrettyString) ID() ID {
	return prettyStringToIDMap[string(q)]
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (q *ID) UnmarshalGQL(v interface{}) error {
	intValue, ok := v.(int)
	if !ok {
		return fmt.Errorf("rank must be an int")
	}

	*q = ID(intValue)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (q ID) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, q)
}
