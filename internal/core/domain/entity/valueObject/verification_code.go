package valueobject

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

type VerificationCode struct {
	Value string `json:"value,omitempty"`
}

var (
	lettersVerificationCode = []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	numbersVerificationCode = []rune(`0123456789`)
)

var regValidCode = regexp.MustCompile(`[a-zA-Z][a-zA-Z][a-zA-Z][0-9][0-9][0-9]`)

func (v *VerificationCode) Generate() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	code := ""

	for i := 0; i < 3; i++ {
		lettersIndex := generateRandom(0, len(lettersVerificationCode)-1)
		code = fmt.Sprintf("%s%s", code, string(lettersVerificationCode[lettersIndex]))

	}

	for i := 0; i < 3; i++ {
		numbersIndex := generateRandom(0, len(numbersVerificationCode)-1)
		code = fmt.Sprintf("%s%s", code, string(numbersVerificationCode[numbersIndex]))
	}

	v.Value = code
}

func (c *VerificationCode) Validate() error {
	valueMatch := regValidCode.FindStringSubmatch(c.Value)

	if valueMatch == nil {
		return fmt.Errorf("verificationCode %s is not valid", c.Value)
	}
	return nil
}

func generateRandom(min, max int) int {
	return min + rand.Intn(max-min)
}
