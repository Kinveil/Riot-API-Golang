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
