package service

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ppopgi-pang/ppopgipang-spine/auth/dto"
	userEntity "github.com/ppopgi-pang/ppopgipang-spine/users/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	jwtSecret        []byte
	jwtRefreshSecret []byte
	db               *gorm.DB
}

type AccessTokenPayload struct {
	UserID int64
	Role   string
}

type RefreshTokenPayload struct {
	UserID  int64
	TokenID string
}

type RotateRefreshTokenResult struct {
	Payload         *RefreshTokenPayload
	NewRefreshToken string
	User            userEntity.User
}

func NewAuthService(db *gorm.DB) *AuthService {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")

	if jwtSecret == "" || jwtRefreshSecret == "" {
		panic("JWT secrets이 필요합니다.")
	}

	return &AuthService{
		jwtSecret:        []byte(jwtSecret),
		jwtRefreshSecret: []byte(jwtRefreshSecret),
		db:               db,
	}
}

func (a *AuthService) SaveRefreshToken(ctx context.Context, userId int64, refreshToken string) error {
	db := a.db.WithContext(ctx)

	return db.Model(&userEntity.User{}).
		Where("id = ?", userId).
		Update("refreshToken", refreshToken).
		Error
}

func (a *AuthService) RotateRefreshToken(refreshToken string) (*RotateRefreshTokenResult, error) {
	payload, err := a.VerifyRefreshToken(refreshToken)
	if err != nil {
		log.Printf("verify refresh token error: %v", err)
		return nil, err
	}

	var user userEntity.User
	if err := a.db.Where("id = ?", payload.UserID).First(&user).Error; err != nil {
		log.Printf("여기1")
		return nil, err
	}

	if user.RefreshToken == nil || *user.RefreshToken != payload.TokenID {
		if user.RefreshToken == nil {
			log.Printf("db refresh token is nil for user %d", payload.UserID)
		} else {
			log.Printf("db tokenID=%s, payload tid=%s", *user.RefreshToken, payload.TokenID)
		}
		log.Printf("%+v", user)
		return nil, errors.New("refresh token reuse detected or revoked")
	}

	newTokenID := uuid.NewString()

	claims := jwt.MapClaims{
		"sub": payload.UserID,
		"tid": newTokenID,
		"typ": "refresh",
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	newSigned, err := token.SignedString(a.jwtRefreshSecret)
	if err != nil {
		log.Printf("여기3")
		return nil, err
	}

	// DB에 새 token id 저장 (rotate)
	user.RefreshToken = &newTokenID
	if err := a.db.Save(&user).Error; err != nil {
		log.Printf("여기4")
		return nil, err
	}

	return &RotateRefreshTokenResult{
		Payload: &RefreshTokenPayload{
			UserID:  payload.UserID,
			TokenID: newTokenID,
		},
		User:            user,
		NewRefreshToken: newSigned,
	}, nil
}

func (a *AuthService) IssueAccessToken(user *userEntity.User) string {
	role := "user"
	if user.IsAdmin {
		role = "admin"
	}

	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": role,
		"typ":  "access",
		"exp":  time.Now().Add(5 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.jwtSecret)
	if err != nil {
		panic(err)
	}

	return signed
}

func (a *AuthService) IssueRefreshToken(user *userEntity.User) string {
	tokenID := uuid.NewString()

	claims := jwt.MapClaims{
		"sub": user.ID,
		"tid": tokenID,
		"typ": "refresh",
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString(a.jwtRefreshSecret)
	if err != nil {
		panic(err)
	}

	user.RefreshToken = &tokenID

	return signed
}

func (a *AuthService) VerifyAccessToken(token string) (*AccessTokenPayload, error) {
	claims, err := a.parseToken(token, a.jwtSecret)
	if err != nil {
		return nil, err
	}

	if typ, _ := claims["typ"].(string); typ != "access" {
		return nil, jwt.ErrTokenMalformed
	}

	userID, err := toInt64(claims["sub"])
	if err != nil {
		return nil, err
	}

	role, _ := claims["role"].(string)
	if role == "" {
		role = "user"
	}

	return &AccessTokenPayload{
		UserID: userID,
		Role:   role,
	}, nil
}

func (a *AuthService) VerifyRefreshToken(token string) (*RefreshTokenPayload, error) {
	claims, err := a.parseToken(token, a.jwtRefreshSecret)
	if err != nil {
		return nil, err
	}

	if typ, _ := claims["typ"].(string); typ != "refresh" {
		return nil, jwt.ErrTokenMalformed
	}

	userID, err := toInt64(claims["sub"])
	if err != nil {
		return nil, err
	}

	tokenID, _ := claims["tid"].(string)
	if tokenID == "" {
		return nil, jwt.ErrTokenMalformed
	}

	return &RefreshTokenPayload{
		UserID:  userID,
		TokenID: tokenID,
	}, nil
}

func (a *AuthService) parseToken(token string, secret []byte) (jwt.MapClaims, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	if !parsed.Valid {
		return nil, jwt.ErrTokenMalformed
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}

func toInt64(value any) (int64, error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case string:
		if v == "" {
			return 0, errors.New("empty sub claim")
		}
		var n int64
		for i := 0; i < len(v); i++ {
			ch := v[i]
			if ch < '0' || ch > '9' {
				return 0, errors.New("invalid sub claim")
			}
			n = n*10 + int64(ch-'0')
		}
		return n, nil
	default:
		return 0, errors.New("invalid sub claim type")
	}
}

func (a *AuthService) CreateAdminUser(ctx context.Context, dto *dto.AdminUserRequest) error {
	db := a.db.WithContext(ctx)

	var existingAdmin userEntity.User
	result := db.Where("email = ?", dto.Email).First(&existingAdmin)
	if result.Error == nil {
		return gorm.ErrDuplicatedKey
	}
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(dto.Password),
		bcrypt.DefaultCost,
	)

	stringPassword := string(hashedPassword)

	if err != nil {
		return err
	}

	user := userEntity.User{
		Email:         dto.Email,
		Nickname:      dto.Email,
		AdminPassword: &stringPassword,
		IsAdmin:       true,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
