package rank

import (
	"fmt"
	"io"
)

type Int int16
type String string

const (
	I   String = "I"
	II  String = "II"
	III String = "III"
	IV  String = "IV"
)

var intToStringMap = map[Int]String{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
}

var stringToIntMap = map[String]Int{
	"I":   1,
	"II":  2,
	"III": 3,
	"IV":  4,
}

func (r String) Int() Int {
	return stringToIntMap[r]
}

func (r Int) String() String {
	return intToStringMap[r]
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (r *Int) UnmarshalGQL(v interface{}) error {
	intValue, ok := v.(int16)
	if !ok {
		return fmt.Errorf("rank must be an int16")
	}

	*r = Int(intValue)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (r Int) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, r)
}
