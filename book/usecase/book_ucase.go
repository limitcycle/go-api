package usecase

import (
	"context"
	"go-api/book"
	"go-api/models"
	"strconv"
	"time"
)

type bookUsecase struct {
	bookRepos      book.BookRepository
	contextTimeOut time.Duration
}

type bookChannel struct {
	Book  *models.Book
	Error error
}

func NewBookUsecase(repos book.BookRepository, timeout time.Duration) book.BookUsecase {
	return &bookUsecase{
		bookRepos:      repos,
		contextTimeOut: timeout,
	}
}

func (buc *bookUsecase) Fetch(c context.Context, cursor string, num int64) ([]*models.Book, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()

	listBook, err := buc.bookRepos.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""
	if size := len(listBook); size == int(num) {
		lastID := listBook[num-1].ID
		nextCursor = strconv.Itoa(int(lastID))
	}
	return listBook, nextCursor, err
}

func (buc *bookUsecase) GetByID(c context.Context, id int64) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()

	res, err := buc.bookRepos.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (buc *bookUsecase) GetByAuthor(c context.Context, author string) ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()

	res, err := buc.bookRepos.GetByAuthor(ctx, author)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (buc *bookUsecase) GetByName(c context.Context, name string) ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()

	res, err := buc.bookRepos.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (buc *bookUsecase) Store(c context.Context, b *models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()

	existID, _ := buc.GetByID(ctx, b.ID)
	if existID != nil {
		return nil, models.CONFLIT_ERROR
	}
	currentBook, err := buc.bookRepos.Store(ctx, b)
	if err != nil {
		return nil, err
	}
	return currentBook, err
}

func (buc *bookUsecase) Update(c context.Context, b *models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()

	return buc.bookRepos.Update(ctx, b)
}

func (buc *bookUsecase) Delete(c context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(c, buc.contextTimeOut)
	defer cancel()
	existedBook, _ := buc.bookRepos.GetByID(ctx, id)
	if existedBook == nil {
		return false, models.NOT_FOUND_ERROR
	}
	return buc.bookRepos.Delete(ctx, id)
}
