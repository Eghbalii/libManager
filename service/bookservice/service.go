package bookservice

import (
	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	AddBook(b entity.Book) (entity.Book, error)
	GetBookByISBN(isbn string) (entity.Book, error)
	GetBookByID(bookID primitive.ObjectID) (entity.Book, error)
	DeleteBookByID(bookID primitive.ObjectID) error
	UpdateBookByID(bookID primitive.ObjectID, b entity.Book) (entity.Book, error)
	ListAllBooks() ([]entity.Book, error)
	BorrowBook(bookID primitive.ObjectID, borrowerID primitive.ObjectID) error
	ReturnBook(bookID primitive.ObjectID) error
	ReserveBook(bookID primitive.ObjectID, reserverID primitive.ObjectID) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
