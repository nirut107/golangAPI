package repository


import (
	"database/sql"
	"go-backend/model"
)

type PostgresUserRepo struct {
	DB *sql.DB
}

func (r PostgresUserRepo) GetAll() ([]model.User, error) {
	rows, err := r.DB.Query(`SELECT id, name FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
		
	}
	return users, nil
}

func (r PostgresUserRepo) GetByID(id int) (model.User, error) {
	var u model.User
	err := r.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return u, nil
}

func (r PostgresUserRepo) Create(u model.User) (model.User, error) {
	err := r.DB.QueryRow(
		`INSERT INTO users (name) VALUES ($1) RETURNING id`,
		u.Name,
	).Scan(&u.ID)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r PostgresUserRepo) Delete(id int) error {
	result, err := r.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (r PostgresUserRepo) Update(u model.User) (model.User, error) {
	result, err := r.DB.Exec(
		`UPDATE users SET name = $1 WHERE id = $2`,
		u.Name,
		u.ID,
	)
	if err != nil {
		return model.User{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.User{}, err
	}
	if rowsAffected == 0 {
		return model.User{}, ErrUserNotFound
	}
	return u, nil
}

func NewUserRepoPostgres(db *sql.DB) UserRepository {
	return &PostgresUserRepo{DB: db}
}