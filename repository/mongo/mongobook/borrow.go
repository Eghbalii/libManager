package mongobook

import (
	"context"
	"time"

	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d DB) BorrowBook(bookID primitive.ObjectID, borrowerID primitive.ObjectID) error {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"_id": bookID}
	update := bson.M{"$set": bson.M{"borrower": borrowerID, "status": entity.Borrowed,
		"borrow_date": time.Now().Format("2006-01-02"), "return_date": time.Now().AddDate(0, 0, 7).Format("2006-01-02")}}
	_, err := bookCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d DB) ReturnBook(bookID primitive.ObjectID) error {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"_id": bookID}
	book, err := d.GetBookByID(bookID)
	if err != nil {
		return err
	}
	var update bson.M
	if book.Status == entity.Borrowed {
		update = bson.M{"$set": bson.M{"borrower": primitive.NilObjectID, "status": entity.Available,
			"borrow_date": "", "return_date": ""}}
	} else if book.Status == entity.Reserved {
		update = bson.M{"$set": bson.M{"reserver": primitive.NilObjectID, "status": entity.Borrowed,
			"borrow_date": time.Now().Format("2006-01-02"), "return_date": time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
			"borrower": book.Reserver}}
	}
	_, err = bookCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d DB) ReserveBook(bookID primitive.ObjectID, reserverID primitive.ObjectID) error {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"_id": bookID}
	update := bson.M{"$set": bson.M{"reserver": reserverID, "status": entity.Reserved}}
	_, err := bookCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
