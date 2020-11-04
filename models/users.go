package models

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Users struct {

	Authenticated   bool    `mapstructure:"authenticated" pg:"authenticated"`

	UserId          string    `mapstructure:"userid" pg:"user_id"`

	Username         string    `mapstructure:"username" pg:"username"`
}

func DecodeUser(data interface{}) (*Users, mapstructure.Metadata, error) {
	var users = Users{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &users,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding users: ", err)
			return nil, metadata, err
		}

		if users.UserId == "" || users.Username == "" {
			//fmt.Println("Username or userid is null ")
			return nil, metadata, errors.New("Username or userid is null")

		}

		return &users, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

func (u *Users) GetType() string {
	return "users"
}

func (u *Users) GetUserId() string {
	return u.UserId
}