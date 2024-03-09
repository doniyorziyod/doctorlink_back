package database

import (
	"database/sql"
	"doctorlink/types"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (p *PostgresRepository) SaveUser(u types.User) error {
	_, err := p.db.Exec("INSERT INTO users (id, username, name, surname, gender, birthday, subregion, region) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		u.Id, u.Username, u.Name, u.Surname, u.Gender, u.Birthday, u.Subregion, u.Region)

	return err
}

func (p *PostgresRepository) AllUsersRepo() ([]types.User, error) {
	rows, _ := p.db.Query("SELECT * from users")
	users := []types.User{}

	for rows.Next() {
		user := types.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Gender, &user.Birthday, &user.Subregion, &user.Region); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *PostgresRepository) GetUserRepo(id uuid.UUID) (types.User, error) {
	users, err := p.AllUsersRepo()
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			return users[i], err
		}
	}

	return types.User{}, err
}

func (p *PostgresRepository) SignInRepo(username string) (uuid.UUID, error) {
	users, err := p.AllUsersRepo()
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			return users[i].Id, err
		}
	}

	return uuid.Nil, err
}

func (p *PostgresRepository) GenerateSmsRepo(username string) (int, error) {
	sms := types.Sms{
		Username: username,
		Sms:      p.GenerateCode(),
	}
	_, err := p.db.Exec("INSERT INTO sms (username, sms) VALUES ($1, $2)",
		sms.Username, sms.Sms)

	return sms.Sms, err
}

func (p *PostgresRepository) GenerateCode() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(90000) + 10000
}

func (p *PostgresRepository) UpdateUserRepo(u types.User) (bool, error) {
	_, err := p.db.Exec("UPDATE users SET phone=$1, name=$2, surname=$3, gender=$4, birthday=$5, subregion=$6, region=$7 WHERE id=$1",
		u.Id, u.Username, u.Name, u.Surname, u.Gender, u.Birthday, u.Subregion, u.Region)

	return true, err
}

func (p *PostgresRepository) SaveBotUser(u types.BotUser) error {
	_, err := p.db.Exec("INSERT INTO bot (id, firstname, lastname, isforum, username, languagecode, isbot, ispremium, addedtomenu, usernames, customemojistatus, canjoingroups, canreadmessages, supportsinline) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		u.ID, u.FirstName, u.LastName, u.IsForum, u.Username, u.LanguageCode, u.IsBot, u.IsPremium, u.AddedToMenu, nil, u.CustomEmojiStatus, u.CanJoinGroups, u.CanReadMessages, u.SupportsInline)

	return err
}
