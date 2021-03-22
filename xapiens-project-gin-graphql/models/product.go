package models

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}

// inisialisasi data (macem databasenya)
var ProductData = []Product{
	{
		ID:    1,
		Name:  "Kopi luwak",
		Info:  "Nyaman di lambung nggak bikin kembung",
		Price: 1500,
	},
	{
		ID:    2,
		Name:  "Indomie",
		Info:  "Mie paling mantap",
		Price: 2500,
	},
	{
		ID:    3,
		Name:  "Coki coki",
		Info:  "Coklat asli",
		Price: 500,
	},
}
