package models

type Coin struct {
	Id             string  `bson:"_id" json:"id"`
	Name           string  `bson:"name" json:"name"`                     // Название монеты
	Country        string  `bson:"country" json:"country"`               // Страна происхождения
	Year           int     `bson:"year" json:"year"`                     // Год выпуска
	Denomination   string  `bson:"denomination" json:"denomination"`     // Номинал
	Material       string  `bson:"material" json:"material"`             // Материал (золото, серебро, медь, сплав)
	Weight         float64 `bson:"weight" json:"weight"`                 // Вес в граммах
	Diameter       float64 `bson:"diameter" json:"diameter"`             // Диаметр в миллиметрах
	Thickness      float64 `bson:"thickness" json:"thickness"`           // Толщина в миллиметрах
	Condition      string  `bson:"condition" json:"condition"`           // Состояние (например, UNC, XF, VG и т.д.)
	MintMark       string  `bson:"mintMark" json:"mintMark"`             // Монетный двор (например, "M" для Москвы)
	HistoricalInfo string  `bson:"historicalInfo" json:"historicalInfo"` // Историческая справка о монете
	Value          float64 `bson:"value" json:"value"`                   // Оценочная стоимость монеты
}
