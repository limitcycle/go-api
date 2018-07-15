package book

import (
	"context"
	model "go-api/models"
)

type BookUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*model.Book, string, error)
	GetByID(ctx context.Context, id int64) (*model.Book, error)
	GetByAuthor(ctx context.Context, author string) ([]*model.Book, error)
	GetByName(ctx context.Context, name string) ([]*model.Book, error)
	Store(ctx context.Context, book *model.Book) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) (*model.Book, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
