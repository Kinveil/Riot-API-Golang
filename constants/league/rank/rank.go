package rank

import (
	"fmt"
	"io"
)

type Int int
type String string

const (
	I Int = iota + 1
	II
	III
	IV
)

var intToStringMap = map[Int]String{
	I:   "I",
	II:  "II",
	III: "III",
	IV:  "IV",
}

var stringToIntMap = map[String]Int{
	"I":   I,
	"II":  II,
	"III": III,
	"IV":  IV,
}

func (r String) Int() Int {
	return stringToIntMap[r]
}

func (r Int) String() String {
	return intToStringMap[r]
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (r *Int) UnmarshalGQL(v interface{}) error {
	intValue, ok := v.(int)
	if !ok {
		return fmt.Errorf("rank must be an int")
	}

	*r = Int(intValue)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (r Int) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, r)
}
