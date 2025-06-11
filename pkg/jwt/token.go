package jwt

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	Secret []byte
}

func NewManager() *Manager {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is required")
	}
	return &Manager{Secret: []byte(secret)}
}

type CustomClaims struct {
	PenggunaID string `json:"pengguna_id"`
	Peran      string `json:"peran_id_str,omitempty"`
	IP         string `json:"ip,omitempty"`
	UserAgent  string `json:"ua,omitempty"`
	jwt.RegisteredClaims
}

func (m *Manager) GenerateAccessToken(userID, peran, ip, ua string) (string, error) {
	claims := CustomClaims{
		PenggunaID: userID,
		Peran:      peran,
		IP:         ip,
		UserAgent:  ua,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.Secret)
}

// GenerateRefreshToken membuat token acak sepanjang 32 byte (64 karakter hex)
func (m *Manager) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32) // 256-bit
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (m *Manager) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return m.Secret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
