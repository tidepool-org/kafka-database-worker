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

func DecodeFood(data interface{}) (*Food, mapstructure.Metadata, error) {
	var food = Food{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &food,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding food: ", err)
			return nil, metadata, err
		}

		if err := food.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &food, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, nil
	}
}
