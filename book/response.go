package book

type BookResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	SubTitle string `json:"sub_title"`
	Rating   int    `json:"rating"`
	Discount int    `json:"discount"`
}
