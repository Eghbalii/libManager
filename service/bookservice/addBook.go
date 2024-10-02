package bookservice

import (
	"github.com/eghbalii/libManager/entity"
	"github.com/eghbalii/libManager/param"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) AddBook(b param.AddBookRequest) (param.AddBookResponse, error) {
	book := entity.Book{
		ID:          primitive.NewObjectID(),
		Title:       b.Title,
		Author:      b.Author,
		ISBN:        b.ISBN,
		Publisher:   b.Publisher,
		PublishDate: b.PublishDate,
		Status:      entity.Available,
	}
	book, err := s.repo.AddBook(book)
	if err != nil {
		return param.AddBookResponse{}, err
	}

	return param.AddBookResponse{
		ID:          book.ID.Hex(),
		Title:       book.Title,
		Author:      book.Author,
		ISBN:        book.ISBN,
		Publisher:   book.Publisher,
		PublishDate: book.PublishDate,
		Status:      book.Status.String(),
	}, nil
}
