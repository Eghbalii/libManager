package bookvalidator

import (
	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (v Validator) ValidateBorrowBookRequest(bookID, userID string) (bool, error) {
	bookObjectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return false, err
	}
	book, err := v.repo.GetBookByID(bookObjectID)
	if err != nil {
		return false, err
	}

	if book.Status == entity.Available {
		return true, nil
	}
	return false, nil
}

func (v Validator) ValidateReturnBookRequest(bookID, userID string) (bool, error) {
	bookObjectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return false, err
	}

	book, err := v.repo.GetBookByID(bookObjectID)
	if err != nil {
		return false, err
	}

	if book.Status == entity.Borrowed || book.Status == entity.Reserved {
		return true, nil
	}
	return false, nil
}

func (v Validator) ValidateReserveBookRequest(bookID, userID string) (bool, error) {
	bookObjectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return false, err
	}
	book, err := v.repo.GetBookByID(bookObjectID)
	if err != nil {
		return false, err
	}

	if book.Status == entity.Borrowed && book.Borrower.Hex() != userID {
		return true, nil
	}
	return false, nil
}
