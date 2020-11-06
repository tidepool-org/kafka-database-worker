package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type PhysicalActivity struct {
	Base `mapstructure:",squash"`

	Duration map[string]interface{} `mapstructure:"duration" pg:"duration" json:"duration"`

	Distance map[string]interface{} `mapstructure:"distance" pg:"distance" json:"distance"`

	Energy map[string]interface{} `mapstructure:"energy" pg:"energy" json:"energy"`

	Name string `mapstructure:"name" pg:"name" json:"name"`
}

func DecodePhysicalActivity(data interface{}) (*PhysicalActivity, mapstructure.Metadata, error) {
	var physicalActivity = PhysicalActivity{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &physicalActivity,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding physical activity: ", err)
			return nil, metadata, err
		}

		if err := physicalActivity.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &physicalActivity, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, nil
	}
}
