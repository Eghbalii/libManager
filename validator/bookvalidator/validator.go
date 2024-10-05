package bookvalidator

import (
	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	IsISBNUnique(isbn string) bool
	GetBookByID(bookID primitive.ObjectID) (entity.Book, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
