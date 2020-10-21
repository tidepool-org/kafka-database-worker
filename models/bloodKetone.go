package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type BloodKetone struct {
	Base                      `mapstructure:",squash"`

	Value          float64    `mapstructure:"value" pg:"value" json:"value"`
	Units          string    `mapstructure:"units" pg:"units" json:"units"`
}

func DecodeBloodKetone(data interface{}) (*BloodKetone, error)  {
	var bloodKetone = BloodKetone{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &bloodKetone,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding bloodKetone: ", err)
			return nil, err
		}

		if err := bloodKetone.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &bloodKetone, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, err
	}
}

