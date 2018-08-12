package usecase_test

import (
	"context"
	"go-api/book/mocks"
	"go-api/models"
	"testing"

	"github.com/golang/mock/gomock"
)

// mock object
func newMockUsecase(t *testing.T) (*gomock.Controller, *mocks.MockBookUsecase) {
	ctrl := gomock.NewController(t)

	return ctrl, mocks.NewMockBookUsecase(ctrl)
}

func TestGetByID(t *testing.T) {
	mockCtrl, mockUcase := newMockUsecase(t)
	defer mockCtrl.Finish()
	// given
	book := &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      0,
	}
	// when&then
	mockUcase.EXPECT().GetByID(context.TODO(), 1).Return(book, nil)
	// call
	mockUcase.GetByID(context.TODO(), 1)
}

func TestFetch(t *testing.T) {
	mockCtrl, mockUcase := newMockUsecase(t)
	defer mockCtrl.Finish()
	// given
	book := &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      0,
	}
	books := make([]*models.Book, 0)
	books = append(books, book)
	// when&then
	mockUcase.EXPECT().Fetch(context.TODO(), "1", 1).Return(books, "1", nil)
	// call
	mockUcase.Fetch(context.TODO(), "1", 1)
}

func TestGetByAuthor(t *testing.T) {
	mockCtrl, mockUcase := newMockUsecase(t)
	defer mockCtrl.Finish()
	// given
	book := &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      0,
	}
	books := make([]*models.Book, 0)
	books = append(books, book)
	// when&then
	mockUcase.EXPECT().GetByAuthor(context.TODO(), "author").Return(books, nil)
	// call
	mockUcase.GetByAuthor(context.TODO(), "author")
}

func TestStore(t *testing.T) {
	mockCtrl, mockUcase := newMockUsecase(t)
	defer mockCtrl.Finish()
	// given
	book := &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      2,
	}
	// when&then
	mockUcase.EXPECT().Store(context.TODO(), book).Return(book, nil)
	// call
	mockUcase.Store(context.TODO(), book)
}

func TestUpdate(t *testing.T) {
	mockCtrl, mockUcase := newMockUsecase(t)
	defer mockCtrl.Finish()
	// given
	book := &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      1,
	}

	// when&then
	mockUcase.EXPECT().Update(context.TODO(), book).Return(book, nil)

	// call
	mockUcase.Update(context.TODO(), book)
}

func TestDelete(t *testing.T) {
	mockCtrl, mockUcase := newMockUsecase(t)
	defer mockCtrl.Finish()
	// given
	id := 1
	// when&then
	mockUcase.EXPECT().Delete(context.TODO(), id).Return(true, nil).Times(1)
	// call
	mockUcase.Delete(context.TODO(), id)
}
