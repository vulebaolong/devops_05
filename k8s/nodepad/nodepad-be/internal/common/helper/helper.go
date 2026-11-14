package helper

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampToRFC3339(timestamp *timestamppb.Timestamp) string {
	if timestamp == nil {
		return ""
	}

	return timestamp.AsTime().UTC().Format(time.RFC3339)
}

func TimeToRFC3339(value time.Time) string {
	return TimestampToRFC3339(timestamppb.New(value))
}

func OptionalTimeToRFC3339(value time.Time) *string {
	if value.IsZero() {
		return nil
	}
	formatted := TimeToRFC3339(value)
	return &formatted
}

func GenerateAccessToken(userID, secret string, expiresIn int64) (string, error) {
	now := time.Now().Unix()
	claims := jwt.MapClaims{
		"userId": userID,
		"iat":    now,
		"exp":    now + expiresIn,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func ParseUserIDFromToken(rawToken, secret string) (*uuid.UUID, error) {
	rawToken = strings.TrimSpace(strings.TrimPrefix(rawToken, "Bearer "))
	if rawToken == "" {
		return nil, nil
	}
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid access token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid access token claims")
	}
	value, ok := claims["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("missing user id")
	}
	id, err := uuid.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}
	return &id, nil
}

func UserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	value, ok := ctx.Value("userId").(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("missing user id")
	}
	return uuid.Parse(value)
}
