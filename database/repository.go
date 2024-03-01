package database

import (
	"doctorlink/types"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}


func (p *PostgresRepository) SaveUser(u types.User) error {
    _, err := p.db.Exec("INSERT INTO users (id, phone, name, surname, gender, birthday, subregion, region) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
    u.Id, u.Phone, u.Name, u.Surname, u.Gender, u.Birthday, u.Subregion, u.Region)

    return err
}

func (p *PostgresRepository) AllUsersRepo() ([]types.User, error) {
    rows, _ := p.db.Query("SELECT * from users")
    users := []types.User{}
    
    for rows.Next() {
        user := types.User{}
        if err := rows.Scan(&user.Id, &user.Phone, &user.Name, &user.Surname, &user.Gender, &user.Birthday, &user.Subregion, &user.Region); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (p *PostgresRepository) getUserRepo(id uuid.UUID) (types.User, error) {
    users, err := p.AllUsersRepo()
    for i := 0; i < len(users); i++ {
       if users[i].Id == id {
           return users[i], err
       } 
    }

    return types.User{}, err
}

func (p *PostgresRepository) signInRepo(phone string) (types.User, error) {
    users, err := p.AllUsersRepo()
    for i := 0; i < len(users); i++ {
       if users[i].Phone == phone {
           return users[i], err
       } 
    }

    return types.User{}, err
}

func (p *PostgresRepository) generateSmsRepo(phone string) (int, error) {
    sms := types.Sms {
        Phone: phone,
        Sms: generateCode(),
    }
    _, err := p.db.Exec("INSERT INTO sms (phone, sms) VALUES ($1, $2)",
    sms.Phone, sms.Sms)

    return sms.Sms, err
}

func generateCode() int {
    rand.New(rand.NewSource(time.Now().UnixNano()))
    return rand.Intn(90000) + 10000
}

func (p *PostgresRepository) updateUserRepo(u types.User) (bool, error) {
    _, err := p.db.Exec("UPDATE users SET phone=$1, name=$2, surname=$3, gender=$4, birthday=$5, subregion=$6, region=$7 WHERE id=$1",
    u.Id, u.Phone, u.Name, u.Surname, u.Gender, u.Birthday, u.Subregion, u.Region)

    return true, err
}
