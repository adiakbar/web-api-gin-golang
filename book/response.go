package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Discount    int    `json:"rating"`
	Rating      int    `json:"discount"`
}