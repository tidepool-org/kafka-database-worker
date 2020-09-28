package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type PumpSettings struct {
	Base                                             `mapstructure:",squash"`

	ActiveSchedule          string                      `mapstructure:"activeSchedule" pg:"active_schedule"`

	BasalSchedules       interface{}      `mapstructure:"basalSchedules" pg:"basal_schedules"`

	BgTarget             []interface{}      `mapstructure:"bgTarget" pg:"bg_target"`

	CarbRatio            []interface{}      `mapstructure:"carbRatio" pg:"carb_ratio"`

	InsulinSensitivity   []interface{}      `mapstructure:"insulinSensitivity" pg:"insulin_sensitivity"`

	units                interface{}      `mapstructure:"units" pg:"units"`
}

func DecodePumpSettings(data interface{}) (*PumpSettings, error) {
	var pumpSettings = PumpSettings{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &pumpSettings,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding pump settings: ", err)
			return nil, err
		}

		if err := pumpSettings.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, err
		}

		return &pumpSettings, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, err
	}
}
