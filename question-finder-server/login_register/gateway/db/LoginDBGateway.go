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

func (g *LoginRegisterDBGateway) RegisterUserGateway(user types.User) (string, error) {
	ctx := context.Background()
	sql := registerUserSql()
	_, err := g.postgresConfig.DB.Exec(ctx, sql, user.ID, user.Username, user.PasswordHash, user.Email, user.FirstName, user.LastName, user.CreatedAt)

	if err != nil {
		log.Println("Login DB Gateway failed", err)
		return "Could not register user", err
	}
	return "Success", nil
}

func registerUserSql() string {
	return `INSERT INTO public.user (id, username, password_hash, email, first_name, last_name, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
}

func findUserByUsernameSql() string {
	return `SELECT id, username, password_hash FROM public.user WHERE username = $1`
}

func findUserByEmailSql() string {
	return `SELECT * FROM public.user WHERE email = $1`
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

func (loginRegisterGateway *LoginRegisterDBGateway) FindUserByEmail(email string) (bool, error) {
	ctx := context.Background()
	sql := findUserByEmailSql()
	rows, err := loginRegisterGateway.postgresConfig.DB.Query(ctx, sql, email)
	if err != nil {
		log.Println("Find User By Email failed", err)
		return true, err
	}
	defer rows.Close()

	if !rows.Next() {
		log.Println("User not found")
		return true, nil
	}

	return false, nil
}
