package jwt
package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken_ReturnsNonEmptyString(t *testing.T) {
	token, err := GenerateToken("testuser")
	if err != nil {
		t.Fatalf("GenerateToken returned unexpected error: %v", err)
	}
	if token == "" {
		t.Fatal("GenerateToken returned empty token")
	}
}

func TestParseToken_RoundTrip(t *testing.T) {
	username := "testuser"
	token, err := GenerateToken(username)
	if err != nil {
		t.Fatalf("GenerateToken returned unexpected error: %v", err)
	}

	got, err := ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken returned unexpected error: %v", err)
	}
	if got != username {
		t.Fatalf("ParseToken returned %q, want %q", got, username)
	}
}

func TestParseToken_InvalidToken(t *testing.T) {
	_, err := ParseToken("this.is.not.a.valid.token")
	if err == nil {
		t.Fatal("ParseToken expected an error for invalid token, got nil")
	}
}

func TestParseToken_ExpiredToken(t *testing.T) {
	// Build a token that is already expired.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "testuser"
	claims["exp"] = time.Now().Add(-time.Hour).Unix() // expired 1h ago
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		t.Fatalf("failed to sign expired token: %v", err)
	}

	_, err = ParseToken(tokenString)
	if err == nil {
		t.Fatal("ParseToken expected an error for expired token, got nil")
	}
}

func TestGenerateToken_DifferentUsersProduceDifferentTokens(t *testing.T) {
	t1, err := GenerateToken("alice")
	if err != nil {
		t.Fatal(err)
	}
	t2, err := GenerateToken("bob")
	if err != nil {
		t.Fatal(err)
	}
	if t1 == t2 {
		t.Fatal("tokens for different users should not be equal")
	}
}
