package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type PhysicalActivity struct {
	Base                                           `mapstructure:",squash"`

	Duration    map[string]interface{}         `mapstructure:"duration" pg:"duration"`

	Distance    map[string]interface{}         `mapstructure:"distance" pg:"distance"`

	Energy      map[string]interface{}         `mapstructure:"energy" pg:"energy"`

	Name           string                         `mapstructure:"name" pg:"name"`
}

func DecodePhysicalActivity(data interface{}) (*PhysicalActivity, error) {
	var physicalActivity = PhysicalActivity{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &physicalActivity,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding physical activity: ", err)
			return nil, err
		}

		if err := physicalActivity.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &physicalActivity, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
