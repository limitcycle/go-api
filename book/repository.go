package book

import (
	"context"
	model "go-api/models"
)

type BookRepository interface {
	Fetch(ctx context.Context, cursor string, num int) ([]*model.Book, error)
	GetByID(ctx context.Context, id int) (*model.Book, error)
	GetByAuthor(ctx context.Context, author string) ([]*model.Book, error)
	GetByName(ctx context.Context, name string) ([]*model.Book, error)
	Store(ctx context.Context, b *model.Book) (*model.Book, error)
	Update(ctx context.Context, b *model.Book) (*model.Book, error)
	Delete(ctx context.Context, id int) (bool, error)
}
