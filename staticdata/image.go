package staticdata

type Image struct {
	Full   string  `json:"full"`
	Group  string  `json:"group"`
	Sprite string  `json:"sprite"`
	H      float64 `json:"h"`
	W      float64 `json:"w"`
	Y      float64 `json:"y"`
	X      float64 `json:"x"`
}
