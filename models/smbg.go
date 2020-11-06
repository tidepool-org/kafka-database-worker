package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Smbg struct {
	Base                    `mapstructure:",squash"`

	Units          string    `mapstructure:"units" pg:"units" json:"units"`

	Value          float64    `mapstructure:"value" pg:"value" json:"value"`
}

func DecodeSmbg(data interface{}) (*Smbg, mapstructure.Metadata, error) {
	var smbg = Smbg{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &smbg,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding smbg: ", err)
			return nil, metadata, err
		}

		if err := smbg.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}


		return &smbg, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, nil
	}
}
