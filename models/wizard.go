package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Wizard struct {
	Base                    `mapstructure:",squash"`

	Bolus          string    `mapstructure:"bolus" pg:"bolus"`
	Units          string    `mapstructure:"units" pg:"units"`

	Recommended    string    `mapstructure:"recommended" pg:"recommended"`

}

func DecodeWizard(data interface{}) *Wizard {
	var wizard = Wizard{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
		Result: &wizard,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
		} else {
			return &wizard
		}

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
