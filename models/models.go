package models

type Coin struct {
	Id             int
	Name           string  // Название монеты
	Country        string  // Страна происхождения
	Year           int     // Год выпуска
	Denomination   string  // Номинал
	Material       string  // Материал (золото, серебро, медь, сплав)
	Weight         float64 // Вес в граммах
	Diameter       float64 // Диаметр в миллиметрах
	Thickness      float64 // Толщина в миллиметрах
	Condition      string  // Состояние (например, UNC, XF, VG и т.д.)
	MintMark       string  // Монетный двор (например, "M" для Москвы)
	HistoricalInfo string  // Историческая справка о монете
	Value          float64 // Оценочная стоимость монеты
}

type UpdateCoin struct {
	Name           string  // Название монеты
	Country        string  // Страна происхождения
	Year           int     // Год выпуска
	Denomination   string  // Номинал
	Material       string  // Материал (золото, серебро, медь, сплав)
	Weight         float64 // Вес в граммах
	Diameter       float64 // Диаметр в миллиметрах
	Thickness      float64 // Толщина в миллиметрах
	Condition      string  // Состояние (например, UNC, XF, VG и т.д.)
	MintMark       string  // Монетный двор (например, "M" для Москвы)
	HistoricalInfo string  // Историческая справка о монете
	Value          float64 // Оценочная стоимость монеты
}
