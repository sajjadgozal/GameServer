package auth

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"sajjadgozal/gameserver/internal/models"
)

type AuthService struct {
	secret string
	db     *gorm.DB
}

type MyClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	TokenId  string `json:"token_id"`
	jwt.StandardClaims
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		secret: "asdf1234",
		db:     db,
	}
}

func (s *AuthService) CreateAccount(name string, email string, password string) models.Account {

	account := models.Account{}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password")
		return account
	}

	account.Name = name
	account.Password = string(hashedPassword)
	account.Email = email

	s.db.Create(&account)

	return account
}

func (s *AuthService) GetAccountByEmail(email string) (models.Account, error) {
	account := models.Account{}
	result := s.db.Where("email = ?", email).First(&account)
	if result.Error != nil {
		return account, result.Error
	}
	return account, nil
}

func (s *AuthService) VerifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GenerateJWT(user models.Account) (string, error) {
	// Set the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour) // 1 day

	// Create custom claims
	claims := &MyClaims{
		UserID:   int(user.ID),
		Username: user.Name,
		TokenId:  user.OAuth,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "sajjad",
			Audience:  "sajjad",
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Id:        strconv.Itoa(int(user.ID)),
			Subject:   "sajjad",
		},
	}

	// Create the token with the claims and the signing method
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte(s.secret) // Replace with your secret key
	tokenString, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) VerifyJWT(tokenString string) (*MyClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil // Replace with your secret key
	})
	// Check for errors during parsing
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
