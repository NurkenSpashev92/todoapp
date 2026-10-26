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

func NewUser(id int, version int, fullName string, PhoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: PhoneNumber,
	}
}

func NewUserUninitialized(fullName string, PhoneNumber *string) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		fullName,
		PhoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLen := len([]rune(u.FullName))
	if fullNameLen < 3 || fullNameLen > 100 {
		return fmt.Errorf(
			"invalide `FullName` len: %d %w",
			fullNameLen,
			core_errors.ErrInvlidArgument,
		)
	}

	if u.PhoneNumber != nil {
		PhoneNumberLength := len([]rune(*u.PhoneNumber))
		if PhoneNumberLength < 10 || PhoneNumberLength > 15 {
			return fmt.Errorf(
				"invalide `PhoneNumber` len: %d %w",
				PhoneNumberLength,
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

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(fullName, phoneNumber Nullable[string]) UserPatch {
	return UserPatch{FullName: fullName, PhoneNumber: phoneNumber}
}

func (u *UserPatch) Validate() error {
	if u.FullName.Set && u.FullName.Value == nil {
		return fmt.Errorf(
			"FullName can't be patched to NULL: %w",
			core_errors.ErrInvlidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
