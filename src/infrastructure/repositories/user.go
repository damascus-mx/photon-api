package infrastructure

import(
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func (u *UserRepository) GetByID() {
	
}
