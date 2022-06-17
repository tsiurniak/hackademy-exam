package users

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
)

func HandleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(err.Error()))
}

func writeResponse(w http.ResponseWriter, status int, response string) {
	w.WriteHeader(status)
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write([]byte(response))
}

type User struct {
	Id       uint
	Email    string
	Password string
}

type RegisterUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserParams RegisterUserParams

type ResetPasswordParams struct {
	Password string `json:"passwod"`
}

type UserRepository interface {
	Add(string, *UserStorage) error
	Get(string) (*UserStorage, error)
	Update(string, *UserStorage) error
	Delete(string) (*UserStorage, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(u UserRepository) *UserService {
	return &UserService{
		repository: u,
	}
}

func validateEmail(email string) error {
	// 1. Email is valid
	match, _ := regexp.Match(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, []byte(email))
	if !match {
		return errors.New("unvalid email address")
	}
	return nil
}

func validatePassword(password string) error {
	// 2. Password at least 8 symbols
	if len(password) < 8 {
		return errors.New("password too short")
	}
	return nil
}

func validateRegisterParams(params RegisterUserParams) error {
	err := validateEmail(params.Email)
	if err != nil {
		return err
	}

	err = validatePassword(params.Password)
	if err != nil {
		return err
	}

	return nil
}

func readParams(r *http.Request, object interface{}) error {
	err := json.NewDecoder(r.Body).Decode(object)
	if err != nil {
		log.Println(err)
		return errors.New("could not read params")
	}

	return nil
}

func (u *UserService) Register(w http.ResponseWriter, r *http.Request) {

	params := RegisterUserParams{}
	err := readParams(r, &params)
	if err != nil {
		HandleError(err, w)
		return
	}

	if err := validateRegisterParams(params); err != nil {
		HandleError(err, w)
		return
	}

	passwordDigest := md5.New().Sum([]byte(params.Password))

	newUser := User{
		Email:    params.Email,
		Password: string(passwordDigest),
	}

	err = u.repository.Add(newUser.Email, NewUserStorageWithUser(newUser))
	if err != nil {
		HandleError(err, w)
		return
	}
	writeResponse(w, http.StatusCreated, "registered")
}
