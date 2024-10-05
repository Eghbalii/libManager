package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/eghbalii/libManager/param"
	"github.com/eghbalii/libManager/service/bookservice"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	book.UnimplementedBookServiceServer
	svc bookservice.Service
}

func New(svc bookservice.Service) *Server {
	return &Server{
		UnimplementedBookServiceServer: book.UnimplementedBookServiceServer{},
		svc:                            svc,
	}
}

func (s Server) Start() {
	listener, err := net.Listen("tcp", ":8086")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	book.RegisterBookServiceServer(grpcServer, &s)

	log.Println("grpc server started at :8086")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve grpc server: ", err)
	}

}

func (b *Server) CreateBook(ctx context.Context, req *book.Book) (*book.Book, error) {
	resp, err := b.svc.AddBook(param.AddBookRequest{
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		Publisher:   req.Publisher,
		PublishDate: req.PublishDate,
	})
	if err != nil {
		return nil, err
	}

	return &book.Book{
		ID:          resp.ID,
		Title:       resp.Title,
		Author:      resp.Author,
		ISBN:        resp.ISBN,
		Publisher:   resp.Publisher,
		PublishDate: resp.PublishDate,
		Status:      resp.Status,
	}, nil
}

func (b *Server) GetBooks(ctx context.Context, empty *emptypb.Empty) (*book.GetAllBooksResponse, error) {
	resp, err := b.svc.ListBooks()
	if err != nil {
		return &book.GetAllBooksResponse{}, err
	}

	books := make([]*book.Book, 0)
	for _, b := range resp {
		books = append(books, &book.Book{
			ID:          b.ID.Hex(),
			Title:       b.Title,
			Author:      b.Author,
			ISBN:        b.ISBN,
			Publisher:   b.Publisher,
			PublishDate: b.PublishDate,
			Status:      b.Status.String(),
			Borrower:    b.Borrower.Hex(),
			Reserver:    b.Reserver.Hex(),
			BorrowDate:  b.BorrowDate,
			ReturnDate:  b.ReturnDate,
		})
	}

	return &book.GetAllBooksResponse{Items: books}, nil
}

func (b *Server) UpdateBook(ctx context.Context, req *book.Book) (*book.Book, error) {
	resp, err := b.svc.UpdateBook(param.UpdateBookRequest{
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		Publisher:   req.Publisher,
		PublishDate: req.PublishDate,
	})

	if err != nil {
		return nil, err
	}

	return &book.Book{
		ID:          resp.ID,
		Title:       resp.Title,
		Author:      resp.Author,
		ISBN:        resp.ISBN,
		Publisher:   resp.Publisher,
		PublishDate: resp.PublishDate,
		Status:      resp.Status,
	}, nil
}

func (b *Server) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*emptypb.Empty, error) {
	bookID, err := primitive.ObjectIDFromHex(req.BookID)
	if err != nil {
		return nil, err
	}

	err = b.svc.DeleteBook(bookID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (b *Server) BorrowBook(ctx context.Context, req *book.BorrowBookRequest) (*emptypb.Empty, error) {
	bookID, err := primitive.ObjectIDFromHex(req.BookID)
	if err != nil {
		return nil, err
	}

	borrowerID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, err
	}

	err = b.svc.BorrowBook(bookID, borrowerID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (b *Server) ReturnBook(ctx context.Context, req *book.ReturnBookRequest) (*emptypb.Empty, error) {
	bookID, err := primitive.ObjectIDFromHex(req.BookID)
	if err != nil {
		return nil, err
	}

	err = b.svc.ReturnBook(bookID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (b *Server) ReserveBook(ctx context.Context, req *book.ReserveBookRequest) (*emptypb.Empty, error) {
	bookID, err := primitive.ObjectIDFromHex(req.BookID)
	if err != nil {
		return nil, err
	}

	reserverID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return nil, err
	}

	err = b.svc.ReserveBook(bookID, reserverID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
