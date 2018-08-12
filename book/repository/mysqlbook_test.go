package repository_test

import (
	"context"
	"go-api/book/mocks"
	"go-api/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func newMockBookRepository(t *testing.T) (*gomock.Controller, *mocks.MockBookRepository) {
	mockCtrl := gomock.NewController(t)

	return mockCtrl, mocks.NewMockBookRepository(mockCtrl)
}

func TestFetch(t *testing.T) {
	mocksCtrl, mockBookRepo := newMockBookRepository(t)
	defer mocksCtrl.Finish()
	// given
	books := make([]*models.Book, 0)
	testBook := &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "description",
		Status:      1,
	}
	books = append(books, testBook)
	// when&then
	mockBookRepo.EXPECT().Fetch(context.TODO(), "1", 1).Return(books, nil)
	// call
	mockBookRepo.Fetch(context.TODO(), "1", 1)
}

func TestGetByID(t *testing.T) {

	mockCtrl, mockBookRepo := newMockBookRepository(t)
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
	mockBookRepo.EXPECT().GetByID(context.TODO(), 1).Return(book, nil)
	// call
	mockBookRepo.GetByID(context.TODO(), 1)
}

func TestGetByAuthor(t *testing.T) {
	mockCtrl, mockBookRepo := newMockBookRepository(t)
	defer mockCtrl.Finish()
	// given
	author := "author"
	books := make([]*models.Book, 0)
	books = append(books, &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      0,
	})
	// when&then
	mockBookRepo.EXPECT().GetByAuthor(context.TODO(), author).Return(books, nil)
	// call
	mockBookRepo.GetByAuthor(context.TODO(), author)
}

func TestGetByName(t *testing.T) {
	mockCtrl, mockBookRepo := newMockBookRepository(t)
	defer mockCtrl.Finish()
	// given
	name := "name"
	books := make([]*models.Book, 0)
	books = append(books, &models.Book{
		ID:          1,
		Author:      "author",
		Name:        "name",
		Description: "desc",
		Status:      2,
	})
	// when&then
	mockBookRepo.EXPECT().GetByName(context.TODO(), name).Return(books, nil)
	// call
	mockBookRepo.GetByName(context.TODO(), name)
}

func TestStore(t *testing.T) {
	mockCtrl, mockBookRepo := newMockBookRepository(t)
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
	mockBookRepo.EXPECT().Store(context.TODO(), book).Return(book, nil)
	// call
	mockBookRepo.Store(context.TODO(), book)
}

func TestUpdate(t *testing.T) {
	mockCtrl, mockBookRepo := newMockBookRepository(t)
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
	mockBookRepo.EXPECT().Update(context.TODO(), book).Return(book, nil)
	// call
	mockBookRepo.Update(context.TODO(), book)
}

func TestDelete(t *testing.T) {
	mockCtrl, mockBookRepo := newMockBookRepository(t)
	defer mockCtrl.Finish()
	// given
	id := 1
	// when&then
	mockBookRepo.EXPECT().Delete(context.TODO(), id).Return(true, nil)
	// call
	mockBookRepo.Delete(context.TODO(), id)
}
