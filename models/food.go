package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type Food struct {
	Base                                           `mapstructure:",squash"`

	Nutrition    map[string]interface{}         `mapstructure:"nutrition" pg:"nutrition" json:"nutrition"`

}

func DecodeFood(data interface{}) (*Food, error) {
	var food = Food{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &food,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding food: ", err)
			return nil, err
		}

		if err := food.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &food, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
