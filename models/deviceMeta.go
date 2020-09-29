package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type DeviceMeta struct {
	Base                                             `mapstructure:",squash"`


        Status             string                        `mapstructure:"status" pg:"status" json:"status,omitempty"`
        SubType            string                        `mapstructure:"subType" pg:"sub_type" json:"subType,omitempty"`
        Duration           int64                         `mapstructure:"duration" pg:"duration" json:"duration,omitempty"`

	Reason    map[string]interface{}              `mapstructure:"reason" pg:"reason" json:"reason"`

}

func DecodeDeviceMeta(data interface{}) (*DeviceMeta, error) {
	var deviceMeta = DeviceMeta{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &deviceMeta,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding device meta: ", err)
			return nil, err
		}

		if err := deviceMeta.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}


		return &deviceMeta, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, nil
	}
}
