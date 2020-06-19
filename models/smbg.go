package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Smbg struct {
	*Base                    `pg:",inherit"`

	Units          string    `mapstructure:"units" pg:"units"`

	value          float64    `mapstructure:"value" pg:"value"`
}

func DecodeSmbg(data interface{}) *Smbg {
	var smbg = Smbg{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
		Result: &smbg,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
		} else {
			return &smbg
		}

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
