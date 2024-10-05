// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: contract/protobuf/book.proto

package book

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author      string `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	ISBN        string `protobuf:"bytes,4,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
	Publisher   string `protobuf:"bytes,5,opt,name=Publisher,proto3" json:"Publisher,omitempty"`
	PublishDate string `protobuf:"bytes,6,opt,name=PublishDate,proto3" json:"PublishDate,omitempty"`
	Status      string `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Borrower    string `protobuf:"bytes,8,opt,name=borrower,proto3" json:"borrower,omitempty"`
	Reserver    string `protobuf:"bytes,9,opt,name=reserver,proto3" json:"reserver,omitempty"`
	BorrowDate  string `protobuf:"bytes,10,opt,name=borrowDate,proto3" json:"borrowDate,omitempty"`
	ReturnDate  string `protobuf:"bytes,11,opt,name=returnDate,proto3" json:"returnDate,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

func (x *Book) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *Book) GetPublishDate() string {
	if x != nil {
		return x.PublishDate
	}
	return ""
}

func (x *Book) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Book) GetBorrower() string {
	if x != nil {
		return x.Borrower
	}
	return ""
}

func (x *Book) GetReserver() string {
	if x != nil {
		return x.Reserver
	}
	return ""
}

func (x *Book) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *Book) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type GetAllBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Book `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetAllBooksResponse) Reset() {
	*x = GetAllBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooksResponse) ProtoMessage() {}

func (x *GetAllBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBooksResponse.ProtoReflect.Descriptor instead.
func (*GetAllBooksResponse) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_book_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllBooksResponse) GetItems() []*Book {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_book_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBookRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

type BorrowBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *BorrowBookRequest) Reset() {
	*x = BorrowBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookRequest) ProtoMessage() {}

func (x *BorrowBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookRequest.ProtoReflect.Descriptor instead.
func (*BorrowBookRequest) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_book_proto_rawDescGZIP(), []int{3}
}

func (x *BorrowBookRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

func (x *BorrowBookRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ReserveBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *ReserveBookRequest) Reset() {
	*x = ReserveBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReserveBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveBookRequest) ProtoMessage() {}

func (x *ReserveBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveBookRequest.ProtoReflect.Descriptor instead.
func (*ReserveBookRequest) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_book_proto_rawDescGZIP(), []int{4}
}

func (x *ReserveBookRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

func (x *ReserveBookRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ReturnBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
}

func (x *ReturnBookRequest) Reset() {
	*x = ReturnBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnBookRequest) ProtoMessage() {}

func (x *ReturnBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnBookRequest.ProtoReflect.Descriptor instead.
func (*ReturnBookRequest) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_book_proto_rawDescGZIP(), []int{5}
}

func (x *ReturnBookRequest) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

var File_contract_protobuf_book_proto protoreflect.FileDescriptor

var file_contract_protobuf_book_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x53, 0x42, 0x4e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f,
	0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x44, 0x22, 0x43, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x44, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42,
	0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2b, 0x0a,
	0x11, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x32, 0x96, 0x03, 0x0a, 0x0b, 0x42,
	0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0a, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x0a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x17, 0x5a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_protobuf_book_proto_rawDescOnce sync.Once
	file_contract_protobuf_book_proto_rawDescData = file_contract_protobuf_book_proto_rawDesc
)

func file_contract_protobuf_book_proto_rawDescGZIP() []byte {
	file_contract_protobuf_book_proto_rawDescOnce.Do(func() {
		file_contract_protobuf_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_protobuf_book_proto_rawDescData)
	})
	return file_contract_protobuf_book_proto_rawDescData
}

var file_contract_protobuf_book_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_contract_protobuf_book_proto_goTypes = []any{
	(*Book)(nil),                // 0: book.Book
	(*GetAllBooksResponse)(nil), // 1: book.GetAllBooksResponse
	(*DeleteBookRequest)(nil),   // 2: book.DeleteBookRequest
	(*BorrowBookRequest)(nil),   // 3: book.BorrowBookRequest
	(*ReserveBookRequest)(nil),  // 4: book.ReserveBookRequest
	(*ReturnBookRequest)(nil),   // 5: book.ReturnBookRequest
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_contract_protobuf_book_proto_depIdxs = []int32{
	0, // 0: book.GetAllBooksResponse.items:type_name -> book.Book
	0, // 1: book.BookService.CreateBook:input_type -> book.Book
	6, // 2: book.BookService.GetBooks:input_type -> google.protobuf.Empty
	0, // 3: book.BookService.UpdateBook:input_type -> book.Book
	2, // 4: book.BookService.DeleteBook:input_type -> book.DeleteBookRequest
	3, // 5: book.BookService.BorrowBook:input_type -> book.BorrowBookRequest
	4, // 6: book.BookService.ReserveBook:input_type -> book.ReserveBookRequest
	5, // 7: book.BookService.ReturnBook:input_type -> book.ReturnBookRequest
	0, // 8: book.BookService.CreateBook:output_type -> book.Book
	1, // 9: book.BookService.GetBooks:output_type -> book.GetAllBooksResponse
	0, // 10: book.BookService.UpdateBook:output_type -> book.Book
	6, // 11: book.BookService.DeleteBook:output_type -> google.protobuf.Empty
	6, // 12: book.BookService.BorrowBook:output_type -> google.protobuf.Empty
	6, // 13: book.BookService.ReserveBook:output_type -> google.protobuf.Empty
	6, // 14: book.BookService.ReturnBook:output_type -> google.protobuf.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_protobuf_book_proto_init() }
func file_contract_protobuf_book_proto_init() {
	if File_contract_protobuf_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_protobuf_book_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_protobuf_book_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllBooksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_protobuf_book_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_protobuf_book_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_protobuf_book_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReserveBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_contract_protobuf_book_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReturnBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_protobuf_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_protobuf_book_proto_goTypes,
		DependencyIndexes: file_contract_protobuf_book_proto_depIdxs,
		MessageInfos:      file_contract_protobuf_book_proto_msgTypes,
	}.Build()
	File_contract_protobuf_book_proto = out.File
	file_contract_protobuf_book_proto_rawDesc = nil
	file_contract_protobuf_book_proto_goTypes = nil
	file_contract_protobuf_book_proto_depIdxs = nil
}
