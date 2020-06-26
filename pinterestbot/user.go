package pinterestbot

import (
	"github.com/jezerdave/go-pinterest-bot/pinterestbot/resource"
	"github.com/jezerdave/go-pinterest-bot/pinterestbot/user"
)


type UserService interface {
	Info(SearchParameters) (*user.Response, error)
}

type userService struct {
	http *resource.Api
}

//Info
func (s userService) Info(param SearchParameters) (*user.Response, error) {

	request := resource.GetParameters(param)
	response := new(user.Response)

	request.Scope = ""
	_, _, err := s.http.Get(resource.UrlUserInfo, request, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewUserService(http *resource.Api) UserService {
	return userService{http: http}
}

