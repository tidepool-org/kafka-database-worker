package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type DeviceMeta struct {
	Base                                             `mapstructure:",squash"`

	Bolus             string                         `mapstructure:"bolus" pg:"bolus"`
	Units             string                         `mapstructure:"units" pg:"units"`

	Status             string                        `mapstructure:"status" pg:"status"`
	SubType            string                        `mapstructure:"subType" pg:"sub_type"`
	Duration           int64                         `mapstructure:"duration" pg:"duration"`

	ReasonMap    map[string]interface{}              `mapstructure:"reason" pg:"-"`
	ReasonJson   string                              `pg:"reason"`

}

func DecodeDeviceMeta(data interface{}) *DeviceMeta {
	var deviceMeta = DeviceMeta{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &deviceMeta,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
			return nil
		}

		reasonByteArray, err := json.Marshal(deviceMeta.ReasonMap)
		deviceMeta.ReasonJson = string(reasonByteArray)
		if err != nil {
			fmt.Println("Error encoding reason json: ", err)
			return nil
		}

		return &deviceMeta

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
