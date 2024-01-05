package rank

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
