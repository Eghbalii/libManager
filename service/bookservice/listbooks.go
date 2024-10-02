package bookservice

import (
	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) ListBooks() ([]entity.Book, error) {
	books, err := s.repo.ListAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s Service) GetBookByID(bookID primitive.ObjectID) (entity.Book, error) {
	book, err := s.repo.GetBookByID(bookID)
	if err != nil {
		return entity.Book{}, err
	}
	return book, nil
}
