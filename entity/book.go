package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	ISBN        string             `bson:"isbn"`
	Publisher   string             `bson:"publisher"`
	PublishDate string             `bson:"publish_date"`
	Status      BookStatus         `bson:"status"`
}

type BookStatus uint8

const (
	Available BookStatus = iota + 1
	Borrowed
	Reserved
)

const (
	AvailableStr = "available"
	BorrowedStr  = "borrowed"
	ReservedStr  = "reserved"
)

func (b BookStatus) String() string {
	switch b {
	case Available:
		return AvailableStr
	case Borrowed:
		return BorrowedStr
	case Reserved:
		return ReservedStr
	default:
		return ""
	}
}

func MapToBookStatusEntity(statusStr string) BookStatus {
	switch statusStr {
	case AvailableStr:
		return Available
	case BorrowedStr:
		return Borrowed
	case ReservedStr:
		return Reserved
	default:
		return 0
	}
}
