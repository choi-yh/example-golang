package util

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateID() int {
	id, err := uuid.NewUUID()
	if err != nil {
		return 0
	}

	return id.ClockSequence()
}

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowPtr() *time.Time {
	now := time.Now().UTC()

	return &now
}

func EncodeHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
