package api42

import (
	"encoding/json"
	"io/ioutil"

	"github.com/nikunicke/hiveboard"
)

type UserService struct {
	user *hiveboard.User
}

func NewUserService() *UserService {
	return &UserService{
		user: nil,
	}
}

func (s *UserService) GetUser(url string) (*hiveboard.User, error) {
	var user *hiveboard.User
	response, err := hiveboard.Client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return user, nil
}
