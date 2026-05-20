package user

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	query := `
		select id,
		       name,
		       email,
		       phone
		from users
		where email = $1
	`

	row := r.db.QueryRow(query, email)

	var u User

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) FindByPhone(phone string) (*User, error) {
	query := `
		select id,
		       name,
		       email,
		       phone
		from users
		where phone = $1
	`

	row := r.db.QueryRow(query, phone)

	var u User

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) Create(req RegisterRequest) (*User, error) {
	query := `
		insert into users(name, email, phone)
		values ($1, $2, $3)
		returning id, name, email, phone
	`

	row := r.db.QueryRow(query, req.Name, req.Email, req.Phone)

	var u User

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Repository) List() ([]*User, error) {
	query := `
		select id,
			   name,
			   email,
			   phone
		from users
		order by id desc
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
