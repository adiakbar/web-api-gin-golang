package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
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

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	return s.repository.Create(book)
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	updateBook, err := s.repository.Update(book)

	return updateBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)

	deleteBook, err := s.repository.Delete(book)

	return deleteBook, err
}
