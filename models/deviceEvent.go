package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type DeviceEvent struct {
	Base                                         `mapstructure:",squash"`

	SubType      string                          `mapstructure:"subType" pg:"sub_type"`
	Units        string                          `mapstructure:"units" pg:"units"`

	Value        float64                         `mapstructure:"value" pg:"value"`

	duration     int64                           `mapstructure:"duration" pg:"duration"`

	Reason       map[string]interface{}          `mapstructure:"reason" pg:"reason"`
	//ReasonJson   string                          `pg:"reason"`

	PrimeTarget  string                          `mapstructure:"primeTarget" pg:"prime_target"`
	Volume       float64                         `mapstructure:"volume" pg:"volume"`
}

func DecodeDeviceEvent(data interface{}) (*DeviceEvent, error) {
	var deviceEvent = DeviceEvent{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &deviceEvent,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding device event: ", err)
			return nil, err
		}

		if err := deviceEvent.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}


		/*nutritionByteArray, err := json.Marshal(deviceEvent.ReasonMap)
		deviceEvent.ReasonJson = string(nutritionByteArray)
		if err != nil {
			fmt.Println("Error encoding nutrition json: ", err)
			return nil, err
		}*/

		return &deviceEvent, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
