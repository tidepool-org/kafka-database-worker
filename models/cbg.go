package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Cbg struct {
	*Base                    `pg:",inherit"`

	SubType        string    `mapstructure:"subType" pg:"sub_type"`

	value          float64    `mapstructure:"value" pg:"value"`
}

func DecodeCbg(data interface{}) *Cbg {
	var cbg = Cbg{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
		Result: &cbg,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
		} else {
			return &cbg
		}

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
