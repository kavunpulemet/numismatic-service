package models

type UpdateCoin struct {
	Name           string  `bson:"name" json:"name"`
	Country        string  `bson:"country" json:"country"`
	Year           int     `bson:"year" json:"year"`
	Denomination   string  `bson:"denomination" json:"denomination"`
	Material       string  `bson:"material" json:"material"`
	Weight         float64 `bson:"weight" json:"weight"`
	Diameter       float64 `bson:"diameter" json:"diameter"`
	Thickness      float64 `bson:"thickness" json:"thickness"`
	Condition      string  `bson:"condition" json:"condition"`
	MintMark       string  `bson:"mintMark" json:"mintMark"`
	HistoricalInfo string  `bson:"historicalInfo" json:"historicalInfo"`
	Value          float64 `bson:"value" json:"value"`
}
