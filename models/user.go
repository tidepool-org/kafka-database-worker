package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type User struct {

	Authenticated   bool    `mapstructure:"authenticated" pg:"authenticated"`

	UserId          string    `mapstructure:"userid" pg:"user_id"`

	Username         string    `mapstructure:"username" pg:"username"`
}

func DecodeUser(data interface{}) (*User, error) {
	var user = User{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &user,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding user: ", err)
			return nil, err
		}

		return &user, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, err
	}
}

func (u *User) GetType() string {
	return "user"
}

func (u *User) GetUserId() string {
	return u.UserId
}