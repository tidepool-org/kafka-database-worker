package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Bolus struct {
	Base                    `mapstructure:",squash"`


    Normal                 float64   `mapstructure:"normal" pg:"normal" json:"normal,omitempty"`
	ExpectedNormal         float64   `mapstructure:"expectedNormal" pg:"expected_normal" json:"expectedNormal,omitempty"`

	Duration                 float64   `mapstructure:"duration" pg:"duration" json:"duration,omitempty"`
	ExpectedDuration         float64   `mapstructure:"expectedDuration" pg:"expected_duration" json:"expectedDuration,omitempty"`

	Extended                 float64   `mapstructure:"extended" pg:"extended" json:"extended,omitempty"`
	ExpectedExtended         float64   `mapstructure:"expectedExtended" pg:"expected_extended" json:"expectedExtended,omitempty"`

	SubType                string    `mapstructure:"subType" pg:"sub_type" json:"subType,omitempty"`
}

func DecodeBolus(data interface{}) (*Bolus, error) {
	var bolus = Bolus{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &bolus,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding bolus: ", err)
			return nil, err
		}

		if err := bolus.DecodeBase(); err != nil {
			fmt.Println("Error decoding base json: ", err)
			return nil, err
		}

		return &bolus, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
