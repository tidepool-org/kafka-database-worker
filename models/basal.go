package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Basal struct {
	Base                      `mapstructure:",squash"`

        DeliveryType      string   `mapstructure:"deliveryType,omitempty" pg:"delivery_type" json:"deliveryType,omitempty"`
        Duration          int64    `mapstructure:"duration,omitempty" pg:"duration" json:"duration,omitempty"`
        ExpectedDuration  int64    `mapstructure:"expectedDuration,omitempty" pg:"expected_duration" json:"expectedDuration,omitempty"`
        Rate              float64  `mapstructure:"rate,omitempty" pg:"rate" json:"rate,omitempty"`
        Percent           float64  `mapstructure:"percent,omitempty" pg:"percent" json:"percent,omitempty"`
        ScheduleName      string   `mapstructure:"scheduleName,omitempty" pg:"schedule_name" json:"scheduleName,omitempty"`
}

func DecodeBasal(data interface{}) (*Basal, mapstructure.Metadata, error)  {
	var basal = Basal{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &basal,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding basal: ", err)
			return nil, metadata, err
		}

		if err := basal.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &basal, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

