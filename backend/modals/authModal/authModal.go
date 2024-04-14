package authmodal

import (
	"errors"
	"time"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Id       string `json:"id,omitempty"`
	User     string `json:"user"`
	Token    string `json:"token"`
	LoggedIn bool   `json:"loggedin"`
}

// [user] -> id of the user
func New(user, token string) *Auth {
	return &Auth{
		User:  user,
		Token: token,
	}
}

// login
func Login(a *surreal.AppRepository, email, password string) (*jwt.Token, error) {

	user, err := surreal.Find[usermodal.User](a,
		"select * from users where email = $email limit 1; ",
		map[string]interface{}{
			"email": email},
		[]usermodal.User{})

	if nil != err {
		return nil, err
	}

	if 0 == len(user) {
		return nil, errors.New("no such user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))

	if nil != err {
		return nil, errors.New("wrong email or password")

	}

	// TODO: add the secrete to tomal file
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user[0].Id,

		// token is valied for a 30 days
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	return token, err
}

// singup
func SignUp(a *surreal.AppRepository, name, email, password, phone, role string) error {

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := usermodal.New(
		name,
		email,
		phone,
		string(hash),
		role,
	)

	err := user.Create(a)

	if nil != err {
		return err
	}

	return nil
}

// signout
