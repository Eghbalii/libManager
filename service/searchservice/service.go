package searchservice

import "github.com/eghbalii/libManager/entity"

type Repository interface {
	SearchBookByTitle(title string) ([]entity.Book, error)
	SearchBookByAuthor(author string) ([]entity.Book, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) SearchBookByTitle(title string) ([]entity.Book, error) {
	return s.repo.SearchBookByTitle(title)
}

func (s Service) SearchBookByAuthor(author string) ([]entity.Book, error) {
	return s.repo.SearchBookByAuthor(author)
}
