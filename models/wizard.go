package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Wizard struct {
	Base                                             `mapstructure:",squash"`

	Bolus             string                         `mapstructure:"bolus" pg:"bolus"`
	Units             string                         `mapstructure:"units" pg:"units"`

	RecommendedMap    map[string]interface{}         `mapstructure:"recommended" pg:"-"`
	RecommendedJson   string                         `pg:"recommended"`

}

func DecodeWizard(data interface{}) *Wizard {
	var wizard = Wizard{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &wizard,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
			return nil
		}

		recommendedByteArray, err := json.Marshal(wizard.RecommendedMap)
		wizard.RecommendedJson = string(recommendedByteArray)
		if err != nil {
			fmt.Println("Error encoding recommended json: ", err)
			return nil
		}

		return &wizard

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}