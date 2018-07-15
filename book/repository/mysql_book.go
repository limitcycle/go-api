package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-api/book"
	"go-api/models"

	"github.com/sirupsen/logrus"
)

type mysqlBookRepository struct {
	Conn *sql.DB
}

func NewMysqlBookRepository(Conn *sql.DB) book.BookRepository {
	return &mysqlBookRepository{Conn}
}

func (m *mysqlBookRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Book, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Book, 0)
	for rows.Next() {
		t := new(models.Book)
		rows.Scan(
			&t.ID,
			&t.Author,
			&t.Name,
			&t.Description,
			&t.Status,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlBookRepository) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Book, error) {
	query := `SELECT id, author, name, description, status
			   FROM book WHERE ID > ? LIMIT ?`
	return m.fetch(ctx, query, cursor, num)
}

func (m *mysqlBookRepository) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	query := `SELECT id, author, name, description, status
		       FROM book WHERE id = ?`

	list, err := m.fetch(ctx, query, id)

	if err != nil {
		return nil, err
	}

	a := &models.Book{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}
	return a, nil
}

func (m *mysqlBookRepository) GetByAuthor(ctx context.Context, author string) ([]*models.Book, error) {
	query := `SELECT id, author, name, description, status
				FROM book WHERE author LIKE ?`
	author = "%" + author + "%"
	list, err := m.fetch(ctx, query, author)
	if err != nil {
		return nil, err
	}

	if len(list) <= 0 {
		return nil, models.NOT_FOUND_ERROR
	}

	return list, nil
}

func (m *mysqlBookRepository) GetByName(ctx context.Context, name string) ([]*models.Book, error) {
	query := `SELECT id, author, name, description, status 
				FROM book WHERE name LIKE ?`
	name = "%" + name + "%"
	list, err := m.fetch(ctx, query, name)
	if err != nil {
		return nil, err
	}

	if len(list) <= 0 {
		return nil, models.NOT_FOUND_ERROR
	}

	return list, nil
}

func (m *mysqlBookRepository) Store(ctx context.Context, b *models.Book) (*models.Book, error) {
	query := `INSERT book SET id=? , author=? , name=? , description=? , status=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return new(models.Book), err
	}
	res, err := stmt.ExecContext(ctx, b.ID, b.Author, b.Name, b.Description, b.Status)
	if err != nil {
		return new(models.Book), err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird Behaviour. Total Affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}
	return b, nil
}

func (m *mysqlBookRepository) Update(ctx context.Context, b *models.Book) (*models.Book, error) {
	query := `UPDATE book SET author=? , name=? , description=? , status=? WHERE ID=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return new(models.Book), err
	}
	res, err := stmt.ExecContext(ctx, b.Author, b.Name, b.Description, b.Status, b.ID)
	if err != nil {
		return new(models.Book), err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird Behaviour. Total Affected: %d", affect)
		logrus.Error(err)
		return nil, err
	}
	return b, nil
}

func (m *mysqlBookRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM book WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAfected != 1 {
		err = fmt.Errorf("Weird Behaviour. Total Affected %d", rowsAfected)
		logrus.Error(err)
		return false, err
	}
	return true, nil
}
