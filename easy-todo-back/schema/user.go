package schema

import (
	"errors"
	"fmt"

	"github.com/dlclark/regexp2"
	"gorm.io/gorm"
)

type validatedEmail string

func newValidatedEmail(email string) (validatedEmail, error) {
	pattern := `^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`
	regexp, err := regexp2.Compile(pattern, 0) //option=none
	if err != nil {
		return "", fmt.Errorf("provided pattern didn't work: %w", err)
	}
	isValidEmail, err := regexp.MatchString(email)
	if err != nil {
		return "", fmt.Errorf("timeout occured during validating: %w", err)
	}
	if !isValidEmail {
		return "", errors.New("provided email didn't match the pattern")
	}
	return validatedEmail(email), nil
}

type validatedPassword string

func newValidatedPassword(password string) (validatedPassword, error) {
	pattern := `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,}$`
	regexp, err := regexp2.Compile(pattern, 0) //option=none
	if err != nil {
		return "", fmt.Errorf("provided pattern didn't work: %w", err)
	}
	isValidPassword, err := regexp.MatchString(password)
	if err != nil {
		return "", fmt.Errorf("timeout occured during validating: %w", err)
	}
	if !isValidPassword {
		return "", errors.New("provided password didn't match the pattern")
	}
	return validatedPassword(password), nil
}

type User struct {
	gorm.Model
	Email    validatedEmail
	password validatedPassword
}

type UserDriver struct {
	User
	db *gorm.DB
}

func NewUserDriver(db *gorm.DB, email string, password string) (*UserDriver, error) {
	ve, err := newValidatedEmail(email)
	if err != nil {
		return &UserDriver{}, err
	}
	vp, err := newValidatedPassword(password)
	if err != nil {
		return &UserDriver{}, err
	}
	NewUserDriver := UserDriver{User: User{Email: ve, password: vp}, db: db}
	if err := NewUserDriver.createUser(); err != nil {
		return &UserDriver{}, err
	}
	return &NewUserDriver, nil
}

func (ud *UserDriver) createUser() error {
	if err := ud.db.Create(ud.User).Error; err != nil {
		return fmt.Errorf("failed to register new account: %w", err)
	}
	return nil
}
