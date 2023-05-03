package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(book BookRequest) (Book, error) {
	newBook := Book{
		Title:       book.Title,
		Price:       book.Price,
		Description: book.SubTitle,
		Discount:    book.Discount,
		Rating:      book.Rating,
	}
	return s.repository.Create(newBook)
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	book.Title = bookRequest.Title
	book.Price = bookRequest.Price
	book.Description = bookRequest.SubTitle
	book.Discount = bookRequest.Discount
	book.Rating = bookRequest.Rating
	return s.repository.Update(book)
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
