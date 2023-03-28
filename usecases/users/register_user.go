package users_case

import (
	"context"
	"net/http"
	"project2/helpers"
	"project2/repository"
	us "project2/repository/users"
	"project2/usecases"
	"regexp"
	"unicode"
)

func (x *usecase) RegisterUser(
	ctx context.Context, req usecases.RegisterUserRequest) (
	res usecases.RegisterUserResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	err = RequestValidation(req)
	if err != nil {
		return res, http.StatusBadRequest, err
	}

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	userPG := us.NewRepository(tx)

	hashPass := helpers.HashPass(req.Password)

	request := repository.RegisterUserRequest{
		UserName: req.UserName,
		Email:    req.Email,
		Password: hashPass,
	}

	_, httpcode, err = userPG.RegisterUser(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.RegisterUserResponse{
		UserName: req.UserName,
		Email:    req.Email,
	}

	return res, httpcode, err
}

func RequestValidation(req usecases.RegisterUserRequest) (err error) {
	emailValid := isEmailValid(req.Email)
	if !emailValid {
		return EmailNotValid
	}

	passwordValid := isPasswordValid(req.Password)
	if !passwordValid {
		return PasswordNotValid
	}

	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func isPasswordValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	/*
	* Password rules:
	* at least 7 letters
	* at least 1 number
	* at least 1 upper case
	* at least 1 special character
	 */

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
