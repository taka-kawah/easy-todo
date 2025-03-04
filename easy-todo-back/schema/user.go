package schema

import (
	"errors"
	"fmt"

	"github.com/dlclark/regexp2"
	"golang.org/x/crypto/bcrypt"
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

type validatedPassword []byte

func newValidatedPassword(password string) (validatedPassword, error) {
	pattern := `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,}$`
	regexp, err := regexp2.Compile(pattern, 0) //option=none
	if err != nil {
		return nil, fmt.Errorf("provided pattern didn't work: %w", err)
	}
	isValidPassword, err := regexp.MatchString(password)
	if err != nil {
		return nil, fmt.Errorf("timeout occured during validating: %w", err)
	}
	if !isValidPassword {
		return nil, errors.New("provided password didn't match the pattern")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to generate hashed password %w", err)
	}
	return validatedPassword(hashed), nil
}

type User struct {
	gorm.Model
	Email    validatedEmail
	Password validatedPassword
}

type UserDriver struct {
	User
	db *gorm.DB
}

func NewUserDriver(db *gorm.DB, email string, password string) (*UserDriver, error) {
	db.AutoMigrate(&User{})

	ve, err := newValidatedEmail(email)
	if err != nil {
		return &UserDriver{}, err
	}
	_, err = findUserFromEmail(db, ve)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &UserDriver{}, errors.New("provided email is already used")
	}
	vp, err := newValidatedPassword(password)
	if err != nil {
		return &UserDriver{}, err
	}
	newUserDriver := UserDriver{User: User{Email: ve, Password: vp}, db: db}
	if err := newUserDriver.createUser(); err != nil {
		return &UserDriver{}, err
	}
	return &newUserDriver, nil
}

func (ud *UserDriver) createUser() error {
	if err := ud.db.Create(ud.User).Error; err != nil {
		return fmt.Errorf("failed to register new account: %w", err)
	}
	return nil
}

func findUserFromEmail(db *gorm.DB, ve validatedEmail) (*User, error) {
	var user User
	res := db.Where("email = ?", ve).First(&user)

	if res.Error != nil {
		return &User{}, fmt.Errorf("failed to find user: %w", res.Error)
	}
	return &user, nil
}

func FindUserDriverFromEmail(db *gorm.DB, email string) (*UserDriver, error) {
	ve, err := newValidatedEmail(email)
	if err != nil {
		return &UserDriver{}, err
	}
	var user User
	res := db.Where("email = ?", ve).First(&user)
	ud := UserDriver{db: db, User: user}

	if res.Error != nil {
		return &UserDriver{}, fmt.Errorf("failed to find user: %w", res.Error)
	}
	return &ud, nil
}

// func Login(db *gorm.DB, email string, password string) (*UserDriver, error) {
// 	ve, err := newValidatedEmail(email)
// 	if err != nil {
// 		return &UserDriver{}, err
// 	}
// 	foundUser, err := findUserFromEmail(db, ve)
// 	if err != nil {
// 		return &UserDriver{}, err
// 	}
// 	vp, err := newValidatedPassword(password)
// 	if err != nil {
// 		return &UserDriver{}, err
// 	}
// 	if err := bcrypt.CompareHashAndPassword(foundUser.Password, vp); err != nil {
// 		return &UserDriver{}, err
// 	}
// 	newUserDriver := UserDriver{User: *foundUser, db: db}
// 	// token, err := newUserDriver.createToken()
// 	// if err != nil{
// 	// 	return &UserDriver{}, nil
// 	// }
// 	return &newUserDriver, nil
// }

// func (ud *UserDriver) createToken()(string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": ud.User.ID,
// 		"email": ud.User.Email,
// 		"exp": time.Now().Add(time.Hour).Unix(),
// 	})
// 	tokenStirng, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
// 	if err != nil{
// 		return "", fmt.Errorf("failed to get token: %w", err)
// 	}
// 	return tokenStirng, nil
// }

// func (ud *UserDriver) decodeToken(){

// }
