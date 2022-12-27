package usecase

import (
	"regexp"
	"strings"

	"github.com/matheus-gondim/go-password-validation/internal/domain/entities"
)

type Verify struct{}

func (v Verify) Exec(verifyPassword entities.VerifyPassword) (entities.Verified, error) {
	res := entities.Verified{
		Verify:  true,
		NoMatch: []string{},
	}

	for _, rule := range verifyPassword.Rules {
		switch rule.Rule {
		case "minSize":
			if !v.minSize(verifyPassword.Password, rule.Value) {
				res.NoMatch = append(res.NoMatch, "minSize")
			}
		case "minUppercase":
			isMinUppercase, err := v.minUppercase(verifyPassword.Password, rule.Value)
			if err != nil {
				return res, err
			}
			if !isMinUppercase {
				res.NoMatch = append(res.NoMatch, "minUppercase")
			}
		case "minLowercase":
			isMinLowercase, err := v.minLowercase(verifyPassword.Password, rule.Value)
			if err != nil {
				return res, err
			}
			if !isMinLowercase {
				res.NoMatch = append(res.NoMatch, "minLowercase")
			}
		case "minDigit":
			isMinDigit, err := v.minDigit(verifyPassword.Password, rule.Value)
			if err != nil {
				return res, err
			}
			if !isMinDigit {
				res.NoMatch = append(res.NoMatch, "minDigit")
			}
		case "minSpecialChars":
			if !v.minSpecialChars(verifyPassword.Password, rule.Value) {
				res.NoMatch = append(res.NoMatch, "minSpecialChars")
			}
		case "noRepeted":
			if !v.noRepeted(verifyPassword.Password) {
				res.NoMatch = append(res.NoMatch, "noRepeted")
			}
		}
	}

	if len(res.NoMatch) > 0 {
		res.Verify = false
	}

	return res, nil
}

func (v Verify) minSize(str string, size int) bool {
	return len(str) >= size
}

func (v Verify) minUppercase(str string, amount int) (bool, error) {
	r, err := regexp.Compile("[A-Z]")
	if err != nil {
		return false, err
	}
	upperLetters := r.FindAllString(str, len(str))
	return len(upperLetters) >= amount, nil
}

func (v Verify) minLowercase(str string, amount int) (bool, error) {
	r, err := regexp.Compile("[a-z]")
	if err != nil {
		return false, err
	}
	lowerLetters := r.FindAllString(str, len(str))
	return len(lowerLetters) >= amount, nil
}

func (v Verify) minDigit(str string, amount int) (bool, error) {
	r, err := regexp.Compile("[0-9]")
	if err != nil {
		return false, err
	}
	digits := r.FindAllString(str, len(str))
	return len(digits) >= amount, nil
}

func (v Verify) minSpecialChars(str string, amount int) bool {
	amountSpecialChar := 0
	specialChar := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '\\', '/', '{', '}', '[', ']'}
	for _, sc := range specialChar {
		if strings.Contains(str, string(sc)) {
			amountSpecialChar++
		}
	}
	return amountSpecialChar >= amount
}

func (v Verify) noRepeted(str string) bool {
	for ix := 0; ix < len(str)-1; ix++ {
		if str[ix] == str[ix+1] {
			return false
		}
	}
	return true
}
