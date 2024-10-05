package mongosearch

import (
	"context"
	"fmt"

	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) SearchBookByAuthor(authorSearchTerm string) ([]entity.Book, error) {
	bookCollection := db.conn.Database("libManager").Collection("book")

	filter := bson.M{
		"author": bson.M{"$regex": fmt.Sprintf(".*%s.*", authorSearchTerm)},
	}

	cursor, err := bookCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var books []entity.Book
	if err = cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (db *DB) SearchBookByTitle(titleSearchTerm string) ([]entity.Book, error) {
	bookCollection := db.conn.Database("libManager").Collection("book")

	filter := bson.M{
		"title": bson.M{"$regex": fmt.Sprintf(".*%s.*", titleSearchTerm)},
	}
	cursor, err := bookCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var books []entity.Book
	if err = cursor.All(context.Background(), &books); err != nil {
		return nil, err
	}
	return books, nil
}
