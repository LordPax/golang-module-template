package code_test

import (
	"fmt"
	"golang-api/code"
	"time"

	"github.com/jaswdr/faker/v2"
)

var fake = faker.New()

func CreateCode(userId, email string) *code.Code {
	return &code.Code{
		ID:        fake.UUID().V4(),
		UserID:    userId,
		Email:     email,
		Code:      fmt.Sprintf("%05d", fake.IntBetween(10000, 99999)),
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}
}

func CreateManyCodes(n int) []*code.Code {
	codes := make([]*code.Code, n)
	for i := 0; i < n; i++ {
		userId := fake.UUID().V4()
		email := fake.Internet().Email()
		codes[i] = CreateCode(userId, email)
	}
	return codes
}
