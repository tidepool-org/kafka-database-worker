package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Users struct {

	Authenticated   bool    `mapstructure:"authenticated" pg:"authenticated"`

	UserId          string    `mapstructure:"userid" pg:"user_id"`

	Username         string    `mapstructure:"username" pg:"username"`
}

func DecodeUser(data interface{}) (*Users, error) {
	var users = Users{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &users,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding users: ", err)
			return nil, err
		}

		return &users, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, err
	}
}

func (u *Users) GetType() string {
	return "users"
}

func (u *Users) GetUserId() string {
	return u.UserId
}