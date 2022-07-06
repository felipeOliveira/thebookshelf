package catalog

type catalogRepository interface {
	GetCategories() ([]Category, error)
	GetBooks(categoryId string) ([]Book, error)
	GetBookDetail(bookId string) (*Book, error)
}
