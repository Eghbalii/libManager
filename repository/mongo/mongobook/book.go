package mongobook

import (
	"context"

	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d DB) AddBook(b entity.Book) (entity.Book, error) {
	bookCollection := d.conn.Database("libManager").Collection("book")

	// Create index
	_, err := bookCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"title":  "text",
				"author": "text",
			},
			Options: options.Index().SetUnique(true),
		},
	)

	_, err = bookCollection.InsertOne(context.Background(), b)
	if err != nil {
		return entity.Book{}, err
	}
	return b, nil
}

func (d DB) IsISBNUnique(isbn string) bool {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"isbn": isbn}
	result := bookCollection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return true
	}
	return false
}

func (d DB) GetBookByISBN(isbn string) (entity.Book, error) {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"isbn": isbn}
	result := bookCollection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return entity.Book{}, err
	}
	var book entity.Book
	err := result.Decode(&book)
	if err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (d DB) GetBookByID(bookID primitive.ObjectID) (entity.Book, error) {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"_id": bookID}
	result := bookCollection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return entity.Book{}, err
	}
	var book entity.Book
	err := result.Decode(&book)
	if err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (d DB) DeleteBookByID(bookID primitive.ObjectID) error {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"_id": bookID}
	_, err := bookCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (d DB) UpdateBookByID(bookID primitive.ObjectID, b entity.Book) (entity.Book, error) {
	bookCollection := d.conn.Database("libManager").Collection("book")
	filter := bson.M{"_id": bookID}

	_, err := bookCollection.UpdateOne(context.Background(), filter, bson.M{"$set": b})
	if err != nil {
		return entity.Book{}, err
	}
	return b, nil
}

func (d DB) ListAllBooks() ([]entity.Book, error) {
	bookCollection := d.conn.Database("libManager").Collection("book")
	cursor, err := bookCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var books []entity.Book
	if err = cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}
	return books, nil
}
