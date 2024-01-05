package queue_ranked

import "github.com/junioryono/Riot-API-Golang/constants/queue"

type ID queue.ID
type String string

const (
	RankedSolo5x5 ID = 420
	RankedFlexSR  ID = 440
	RankedFlexTT  ID = 470
)

var stringToIDMap = map[String]ID{
	"RANKED_SOLO_5X5": RankedSolo5x5,
	"RANKED_FLEX_SR":  RankedFlexSR,
	"RANKED_FLEX_TT":  RankedFlexTT,
}

var idToStringMap = map[ID]String{
	RankedSolo5x5: "RANKED_SOLO_5X5",
	RankedFlexSR:  "RANKED_FLEX_SR",
	RankedFlexTT:  "RANKED_FLEX_TT",
}

func (q ID) String() String {
	return idToStringMap[q]
}

func (q String) ID() ID {
	return stringToIDMap[q]
}

func (q String) PrettyString() queue.PrettyString {
	return queue.PrettyString(q)
}
