package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type DeviceEvent struct {
	Base                                         `mapstructure:",squash"`

        SubType      string                          `mapstructure:"subType" pg:"sub_type" json:"subType,omitempty"`
        Units        string                          `mapstructure:"units" pg:"units" json:"units,omitempty"`

        Value        float64                         `mapstructure:"value" pg:"value" json:"value,omitempty"`

        duration     int64                           `mapstructure:"duration" pg:"duration" json:"duration,omitempty"`

	Reason       map[string]interface{}          `mapstructure:"reason" pg:"reason" json:"reason,omitempty"`

	PrimeTarget  string                          `mapstructure:"primeTarget" pg:"prime_target" json:"primeTarget,omitempty"`
	Volume       float64                         `mapstructure:"volume" pg:"volume" json:"volume,omitempty"`
}

func DecodeDeviceEvent(data interface{}) (*DeviceEvent, mapstructure.Metadata, error) {
	var deviceEvent = DeviceEvent{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &deviceEvent,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding device event: ", err)
			return nil, metadata, err
		}

		if err := deviceEvent.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}


		/*nutritionByteArray, err := json.Marshal(deviceEvent.ReasonMap)
		deviceEvent.ReasonJson = string(nutritionByteArray)
		if err != nil {
			fmt.Println("Error encoding nutrition json: ", err)
			return nil, err
		}*/

		return &deviceEvent, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, nil
	}
}
