package app

type QueryResponse struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	FilterCategory []GraphCategory `json:"filterCategory"`
	FilterProduct  []GraphProduct  `json:"filterProduct"`
}

type GraphCategory struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	ParentCategory *ParentCategory `json:"parentCategory"`
}

type ParentCategory struct {
	AdditionalInfo *AdditionalInfo `json:"additionalInfo"`
}

type AdditionalInfo struct {
	CategoryID int `json:"categoryId"`
}

type GraphProduct struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	Price           float64          `json:"price"`
	OldPrice        *float64         `json:"oldPrice"`
	ProductCategory *ProductCategory `json:"category"`
	Images          []Image          `json:"images"`
}

type ProductCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
}
