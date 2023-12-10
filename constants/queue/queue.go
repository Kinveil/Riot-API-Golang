package queue

import (
	"encoding/json"
	"fmt"
)

type Queue int

const (
	RankedSolo5x5 Queue = 420
	RankedFlexSR  Queue = 440
	RankedFlexTT  Queue = 470

	Snowdown                      Queue = 73
	Hexakill                      Queue = 75
	Nemesis                       Queue = 310
	BlackMarketBrawlers           Queue = 313
	DefinitelyNotDominion         Queue = 317
	AllRandom                     Queue = 325
	NormalDraft                   Queue = 400
	NormalBlind                   Queue = 430
	ARAM                          Queue = 450
	BloodHuntAssassin             Queue = 600
	DarkStarSingularity           Queue = 610
	Clash                         Queue = 700
	CoopVsAI                      Queue = 850
	URF                           Queue = 900
	Ascension                     Queue = 910
	LegendOfThePoroKing           Queue = 920
	NexusSiege                    Queue = 940
	DoomBotsVoting                Queue = 950
	DoomBotsStandard              Queue = 960
	StarGuardianInvasionNormal    Queue = 980
	StarGuardianInvasionOnslaught Queue = 990
	ProjectHunters                Queue = 1000
	ARURF                         Queue = 1010
	OneForAll                     Queue = 1020
	OdysseyExtractionIntro        Queue = 1030
	OdysseyExtractionCadet        Queue = 1040
	OdysseyExtractionCrewmember   Queue = 1050
	OdysseyExtractionCaptain      Queue = 1060
	OdysseyExtractionOnslaught    Queue = 1070
	NexusBlitz                    Queue = 1300
	Ultimates                     Queue = 1400
	Arena                         Queue = 1700
	Tutorial                      Queue = 2020
)

var queueMap = map[string]Queue{
	"RANKED_SOLO_5X5": RankedSolo5x5,
	"RANKED_FLEX_SR":  RankedFlexSR,
	"RANKED_FLEX_TT":  RankedFlexTT,
}

var queueIDMap = map[Queue]string{
	RankedSolo5x5: "RANKED_SOLO_5X5",
	RankedFlexSR:  "RANKED_FLEX_SR",
	RankedFlexTT:  "RANKED_FLEX_TT",
}

func (q *Queue) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if v, ok := queueMap[s]; ok {
		*q = v
		return nil
	}

	var i int
	if err := json.Unmarshal(b, &i); err == nil {
		*q = Queue(i)
		return nil
	}

	return fmt.Errorf("invalid queue %q", s)
}

func (q Queue) String() string {
	if s, ok := queueIDMap[q]; ok {
		return s
	}
	return fmt.Sprintf("%d", q)
}
