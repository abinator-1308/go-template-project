package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/devlibx/gox-base/errors"
	"github.com/harishb2k/go-template-project/pkg/common/objects"
	db "github.com/harishb2k/go-template-project/pkg/database/mysql/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func (u *UserRepository) Persist(ctx context.Context, user *objects.User) error {
	q := db.New(u.db)
	_, err := q.PersistUser(ctx, db.PersistUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Property: user.Key,
	})
	return err
}

func (u *UserRepository) Get(ctx context.Context, user *objects.User) (*objects.User, error) {
	q := db.New(u.db)
	if result, err := q.GetUser(ctx, db.GetUserParams{
		ID:       user.ID,
		Property: user.Key,
	}); err != nil {
		return nil, err
	} else {
		return &objects.User{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}
}

func NewUserRepository(_db *sql.DB) (*UserRepository, error) {
	ud := &UserRepository{
		db: _db,
	}
	return ud, nil
}

func NewMySQLDb(config *MySQLConfig) (*sql.DB, error) {
	config.SetupDefaults()
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Db))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open SQL db")
	}
	return _db, nil
}
