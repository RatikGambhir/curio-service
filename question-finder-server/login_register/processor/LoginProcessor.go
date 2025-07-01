package processor

import (
	"os"
	AppConfig "question_finder/app"
	dbGateway "question_finder/login_register/gateway/db"
	"question_finder/login_register/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginProcessor struct {
	LoginDBGateway *dbGateway.LoginRegisterDBGateway
}

func ConstructLoginProcessor(appConfig *AppConfig.PostgresConfig) *LoginProcessor {
	return &LoginProcessor{
		LoginDBGateway: dbGateway.ConstructLoginRegisterDBGateway(appConfig),
	}
}

func (loginProcessor *LoginProcessor) RegisterUser(req types.RegisterRequest) string {
	hashedPassword, err := genHashedPassword(req.Password)
	if err != nil {
		return "Password hashing failed"
	}
	user := types.User{
		Username:     req.Username,
		Email:        req.Email,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}
	return loginProcessor.LoginDBGateway.RegisterUserGateway(user)
}

func (loginProcessor *LoginProcessor) LoginUser(req types.LoginRequest) (string, error) {
	user, err := loginProcessor.LoginDBGateway.FindUserByUsername(req.Username)
	if err != nil {
		return "Failed to find user! email is incorrect", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "Invalid password", err
	}
	token, err := genToken(*user.ID)
	if err != nil {
		return "Failed to generate token", err
	}

	return token, nil
}

func genToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// func verifyToken(token string) (string, error) {
// 	token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return token.Claims.(jwt.MapClaims)["userID"].(string), nil
// }

func genHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
