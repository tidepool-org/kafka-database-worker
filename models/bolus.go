package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Bolus struct {
	tableName struct{} `pg:"bolus"`

	*Base                    `pg:",inherit"`

	Normal         float64   `mapstructure:"normal" pg:"normal"`

	SubType        string    `mapstructure:"subType" pg:"sub_type"`
}

func DecodeBolus(data interface{}) *Bolus {
	var bolus = Bolus{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
		Result: &bolus,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
		} else {
			return &bolus
		}

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
