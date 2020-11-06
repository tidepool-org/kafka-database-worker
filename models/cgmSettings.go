package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type CgmSettings struct {
	Base                                             `mapstructure:",squash"`

        TransmitterId      string                        `mapstructure:"transmitterId" pg:"transmitter_id" json:"transmitterId,omitempty"`
        Units             string                         `mapstructure:"units" pg:"units" json:"units,omitempty"`

	LowAlerts    map[string]interface{}           `mapstructure:"lowAlerts" pg:"low_alerts" json:"lowAlerts,omitempty"`

	HighAlerts    map[string]interface{}           `mapstructure:"highAlerts" pg:"high_alerts" json:"highAlerts,omitempty"`

	RateOfChangeAlerts    map[string]interface{}   `mapstructure:"rateOfChangeAlerts" pg:"rate_of_change_alerts" json:"rateOfChangeAlerts,omitempty"`

	OutOfRangeAlerts    map[string]interface{}    `mapstructure:"outOfRangeAlerts" pg:"out_of_range_alerts" json:"outOfRangeAlerts,omitempty"`
}

func DecodeCgmSettings(data interface{}) (*CgmSettings, mapstructure.Metadata, error) {
	var cgmSettings = CgmSettings{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &cgmSettings,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding cgm settings: ", err)
			return nil, metadata, err
		}

		if err := cgmSettings.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &cgmSettings, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, nil
	}
}
