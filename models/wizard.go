package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type Wizard struct {
	Base                                             `mapstructure:",squash"`

        Bolus             string                         `mapstructure:"bolus" pg:"bolus" json:"bolus,omitempty"`
        Units             string                         `mapstructure:"units" pg:"units" json:"units,omitempty"`

	Recommended    map[string]interface{}         `mapstructure:"recommended" pg:"recommended" json:"recommended"`

}

func DecodeWizard(data interface{}) (*Wizard, error) {
	var wizard = Wizard{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &wizard,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding wizard: ", err)
			return nil, err
		}

		if err := wizard.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &wizard, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
