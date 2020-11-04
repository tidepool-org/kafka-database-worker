package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Insulin struct {
	Base                `mapstructure:",squash"`

	Dose             map[string]interface{}      `mapstructure:"dose" pg:"dose" json:"dose,omitempty"`
}

func DecodeInsulin(data interface{}) (*Insulin, mapstructure.Metadata, error)  {
	var insulin = Insulin{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &insulin,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding insulin: ", err)
			return nil, metadata, err
		}

		if err := insulin.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &insulin, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

