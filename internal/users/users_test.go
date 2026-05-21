package users

import "testing"

func TestHashPassword_ProducesNonEmptyHash(t *testing.T) {
	hash, err := HashPassword("mysecret")
	if err != nil {
		t.Fatalf("HashPassword returned unexpected error: %v", err)
	}
	if hash == "" {
		t.Fatal("HashPassword returned empty hash")
	}
}

func TestHashPassword_DifferentCallsProduceDifferentHashes(t *testing.T) {
	// bcrypt salts automatically, so two hashes for the same password differ.
	h1, err := HashPassword("mysecret")
	if err != nil {
		t.Fatal(err)
	}
	h2, err := HashPassword("mysecret")
	if err != nil {
		t.Fatal(err)
	}
	if h1 == h2 {
		t.Fatal("expected different hashes for the same password due to random salt")
	}
}

func TestCheckPasswordHash_CorrectPassword(t *testing.T) {
	password := "correcthorsebatterystaple"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	if !CheckPasswordHash(password, hash) {
		t.Fatal("CheckPasswordHash returned false for a correct password")
	}
}

func TestCheckPasswordHash_WrongPassword(t *testing.T) {
	hash, err := HashPassword("correctpassword")
	if err != nil {
		t.Fatal(err)
	}
	if CheckPasswordHash("wrongpassword", hash) {
		t.Fatal("CheckPasswordHash returned true for a wrong password")
	}
}

func TestCheckPasswordHash_EmptyPassword(t *testing.T) {
	hash, err := HashPassword("")
	if err != nil {
		t.Fatal(err)
	}
	if !CheckPasswordHash("", hash) {
		t.Fatal("CheckPasswordHash returned false for an empty password matched against its hash")
	}
	if CheckPasswordHash("notempty", hash) {
		t.Fatal("CheckPasswordHash returned true for non-empty password against empty-password hash")
	}
}
