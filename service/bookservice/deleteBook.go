package bookservice

import "go.mongodb.org/mongo-driver/bson/primitive"

func (s Service) DeleteBook(bookID primitive.ObjectID) error {
	err := s.repo.DeleteBookByID(bookID)
	if err != nil {
		return err
	}

	return nil
}
