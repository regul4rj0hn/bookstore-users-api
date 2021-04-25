package services

import (
	"github.com/regul4rj0hn/bookstore-users-api/models/users"
	"github.com/regul4rj0hn/bookstore-users-api/utils/crypto"
	"github.com/regul4rj0hn/bookstore-users-api/utils/dates"
	"github.com/regul4rj0hn/bookstore-users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	Create(users.User) (*users.User, *errors.Response)
	Get(int64) (*users.User, *errors.Response)
	Search(string) (users.Users, *errors.Response)
	Update(users.User, bool) (*users.User, *errors.Response)
	Delete(int64) *errors.Response
}

func (s *usersService) Create(user users.User) (*users.User, *errors.Response) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.CreatedOn = dates.GetNowDbFormat()
	user.Password = crypto.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) Get(userId int64) (*users.User, *errors.Response) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) Search(status string) (users.Users, *errors.Response) {
	users := &users.User{}
	return users.FindByStatus(status)
}

func (s *usersService) Update(user users.User, isPartial bool) (*users.User, *errors.Response) {
	current, err := UsersService.Get(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Edit(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) Delete(userId int64) *errors.Response {
	user := &users.User{Id: userId}
	return user.Delete()
}
