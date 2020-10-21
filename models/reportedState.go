package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type ReportedState struct {
	Base                      `mapstructure:",squash"`

	States             map[string]interface{}      `mapstructure:"states" pg:"states" json:"states,omitempty"`
}

func DecodeReportedState(data interface{}) (*ReportedState, error)  {
	var reportedState = ReportedState{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &reportedState,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding reportedState: ", err)
			return nil, err
		}

		if err := reportedState.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &reportedState, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, err
	}
}

