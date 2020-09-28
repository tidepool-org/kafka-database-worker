package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type CgmSettings struct {
	Base                                             `mapstructure:",squash"`

	TransmitterId      string                        `mapstructure:"transmitterId" pg:"transmitter_id"`
	Units             string                         `mapstructure:"units" pg:"units"`

	LowAlerts    map[string]interface{}           `mapstructure:"lowAlerts" pg:"low_alerts"`

	HighAlerts    map[string]interface{}           `mapstructure:"highAlerts" pg:"high_alerts"`

	RateOfChangeAlerts    map[string]interface{}   `mapstructure:"rateOfChangeAlerts" pg:"rate_of_change_alerts"`

	OutOfRangeAlerts    map[string]interface{}    `mapstructure:"outOfRangeAlerts" pg:"out_of_range_alerts"`
}

func DecodeCgmSettings(data interface{}) (*CgmSettings, error) {
	var cgmSettings = CgmSettings{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &cgmSettings,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding cgm settings: ", err)
			return nil, err
		}

		if err := cgmSettings.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &cgmSettings, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
