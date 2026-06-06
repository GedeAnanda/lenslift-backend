package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/model"
	"github.com/nanda/lenslift-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo    *repository.UserRepository
	profileRepo *repository.ProfileRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo:    repository.NewUserRepository(),
		profileRepo: repository.NewProfileRepository(),
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email sudah terdaftar")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hash),
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	profile := &model.Profile{
		UserID:   user.ID,
		FullName: req.FullName,
	}
	if err := s.profileRepo.Create(profile); err != nil {
		return nil, err
	}

	token, err := generateToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken: token,
		User: dto.UserProfile{
			ID:       user.ID.String(),
			FullName: req.FullName,
			Email:    user.Email,
		},
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	token, err := generateToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	profile, _ := s.profileRepo.FindByUserID(user.ID.String())
	fullName := ""
	if profile != nil {
		fullName = profile.FullName
	}

	return &dto.AuthResponse{
		AccessToken: token,
		User: dto.UserProfile{
			ID:       user.ID.String(),
			FullName: fullName,
			Email:    user.Email,
		},
	}, nil
}

func generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}