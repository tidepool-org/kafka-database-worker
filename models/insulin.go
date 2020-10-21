package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Insulin struct {
	Base                `mapstructure:",squash"`

	Insulin             map[string]interface{}      `mapstructure:"insulin" pg:"insulin" json:"insulin,omitempty"`
}

func DecodeInsulin(data interface{}) (*Insulin, error)  {
	var insulin = Insulin{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &insulin,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding insulin: ", err)
			return nil, err
		}

		if err := insulin.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &insulin, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, err
	}
}

