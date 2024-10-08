package bookservice

import (
	"github.com/eghbalii/libManager/entity"
	"github.com/eghbalii/libManager/param"
)

func (s Service) UpdateBook(b param.UpdateBookRequest) (param.UpdateBookResponse, error) {
	book, err := s.repo.GetBookByISBN(b.ISBN)
	if err != nil {
		return param.UpdateBookResponse{}, err
	}
	if b.Title != "" {
		book.Title = b.Title
	}
	if b.Author != "" {
		book.Author = b.Author
	}
	if b.ISBN != "" {
		book.ISBN = b.ISBN
	}
	if b.Publisher != "" {
		book.Publisher = b.Publisher
	}
	if b.PublishDate != "" {
		book.PublishDate = b.PublishDate
	}
	if b.Status != "" {
		book.Status = entity.MapToBookStatusEntity(b.Status)
	}

	book, err = s.repo.UpdateBookByID(book.ID, book)
	if err != nil {
		return param.UpdateBookResponse{}, err
	}

	return param.UpdateBookResponse{
		ID:          book.ID.Hex(),
		Title:       book.Title,
		Author:      book.Author,
		ISBN:        book.ISBN,
		Publisher:   book.Publisher,
		PublishDate: book.PublishDate,
		Status:      book.Status.String(),
	}, nil
}
