package repository

import (
	"github.com/jmoiron/sqlx"
	"orderManager/internal/domain"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (p *Repository) CreateNewOrder(o domain.Order) (int, error) {
	//var id int
//
//
	//q := fmt.Sprintf("INSERT INTO order (fPas, lPas, track, status) values ($1, $2, $3, $4) RETURNING id")
	//row := r.db.QueryRow(q, string(o.FirstLat) + " ", user.Username, user.Password, user.Permissions)
	//err := row.Scan(&id)
	//if err != nil {
	//	return 0, err
	//}
	return 0, nil
}


