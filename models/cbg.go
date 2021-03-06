package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Cbg struct {
	Base                    `mapstructure:",squash"`

	Value          float64    `mapstructure:"value" pg:"value"  json:"value,omitempty"`

	Units          string    `mapstructure:"units" pg:"units"  json:"units,omitempty"`
}

func DecodeCbg(data interface{}) (*Cbg, mapstructure.Metadata, error) {
	var cbg = Cbg{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &cbg,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding cbg: ", err)
			return nil, metadata, err
		}

		if err := cbg.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &cbg, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}
