package tier

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type String string
type PrettyString string

const (
	Challenger  String = "CHALLENGER"
	Grandmaster String = "GRANDMASTER"
	Master      String = "MASTER"
	Diamond     String = "DIAMOND"
	Emerald     String = "EMERALD"
	Platinum    String = "PLATINUM"
	Gold        String = "GOLD"
	Silver      String = "SILVER"
	Bronze      String = "BRONZE"
	Iron        String = "IRON"
)

var prettyStringToString = map[PrettyString]String{
	"Challenger":  Challenger,
	"Grandmaster": Grandmaster,
	"Master":      Master,
	"Diamond":     Diamond,
	"Emerald":     Emerald,
	"Platinum":    Platinum,
	"Gold":        Gold,
	"Silver":      Silver,
	"Bronze":      Bronze,
	"Iron":        Iron,
}

func (t String) PrettyString() PrettyString {
	return PrettyString(cases.Title(language.English).String(string(t)))
}

func (t PrettyString) String() String {
	return prettyStringToString[t]
}
