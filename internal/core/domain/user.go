package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/nurkenspashev92/todoapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(id int, version int, fullName string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalide `FullName` len: %d %w",
			fullNameLength,
			core_errors.ErrInvlidArgument,
		)
	}

	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf(
				"invalide `PhoneNumber` len: %d %w",
				phoneNumberLength,
				core_errors.ErrInvlidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalide `PhoneNumber` format: %w",
				core_errors.ErrInvlidArgument,
			)
		}
	}

	return nil
}
