package book

type BookRequest struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required,number"`
	SubTitle string `json:"sub_title"`
	Rating   int    `json:"rating" binding:"required"`
	Discount int    `json:"discount" binding:"required"`
}
