package db

import (
	"context"
	"errors"
	"log"
	AppConfig "question_finder/app"
	"question_finder/login_register/types"
)

type LoginRegisterDBGateway struct {
	postgresConfig *AppConfig.PostgresConfig
}

func ConstructLoginRegisterDBGateway(appConfig *AppConfig.PostgresConfig) *LoginRegisterDBGateway {
	return &LoginRegisterDBGateway{
		postgresConfig: appConfig,
	}
}

func (g *LoginRegisterDBGateway) RegisterUserGateway(user types.User) string {
	ctx := context.Background()
	sql := registerUserSql()
	_, err := g.postgresConfig.DB.Exec(ctx, sql, user.Username, user.PasswordHash, user.Email, user.FirstName, user.LastName, user.CreatedAt)

	if err != nil {
		log.Printf("Insert failed: %v", err)
		return "Failed to register user"
	}
	return "Success"
}

func registerUserSql() string {
	return `INSERT INTO users (username, password_hash, email, first_name, last_name, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
}

func findUserByUsernameSql() string {
	return `SELECT id, username, password_hash FROM users WHERE username = $1`
}

func findUserByEmailSql() string {
	return `SELECT * FROM users WHERE email = $1`
}

func (loginRegisterGateway *LoginRegisterDBGateway) FindUserByUsername(username string) (types.User, error) {
	ctx := context.Background()
	sql := findUserByUsernameSql()
	rows, err := loginRegisterGateway.postgresConfig.DB.Query(ctx, sql, username)
	if err != nil {
		return types.User{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return types.User{}, errors.New("user not found")
	}
	var user types.User
	err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return types.User{}, errors.New("failed to scan user")
	}
	return user, nil
}
