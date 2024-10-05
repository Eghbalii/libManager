package bookservice

import "go.mongodb.org/mongo-driver/bson/primitive"

func (s Service) BorrowBook(bookID primitive.ObjectID, borrowerID primitive.ObjectID) error {
	err := s.repo.BorrowBook(bookID, borrowerID)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) ReturnBook(bookID primitive.ObjectID) error {
	err := s.repo.ReturnBook(bookID)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) ReserveBook(bookID primitive.ObjectID, reserverID primitive.ObjectID) error {
	err := s.repo.ReserveBook(bookID, reserverID)
	if err != nil {
		return err
	}

	return nil
}
