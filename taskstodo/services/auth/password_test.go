package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to be not empty")
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Errorf("expected password to match hash")
	}
	if ComparePasswords(hash, []byte("notpassword")) {
		t.Errorf("expected password to not match hash")
	}
}

func TestComparePasswords_IncorrectPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Fatalf("error hashing password: %v", err)
	}

	if ComparePasswords(hash, []byte("wrongpassword")) {
		t.Errorf("expected password to not match hash")
	}
}

func TestHashPassword_Error(t *testing.T) {
	_, err := HashPassword("")
	if err == nil {
		t.Error("expected error for empty password")
	}
}
