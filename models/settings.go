package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Settings struct {
	Base                      `mapstructure:",squash"`

	Source               string                       `mapstructure:"source" pg:"source" json:"source,omitempty"`
	GroupId              string                      `mapstructure:"groupId" pg:"group_id" json:"groupId,omitempty"`
	ActiveSchedule       string                       `mapstructure:"activeSchedule" pg:"active_schedule" json:"activeSchedule,omitempty"`

	Units                map[string]interface{}      `mapstructure:"units" pg:"units" json:"units,omitempty"`

	BasalSchedules       map[string]interface{}      `mapstructure:"basalSchedules" pg:"basal_schedules" json:"basalSchedules,omitempty"`
	CarbRatio            []interface{}      `mapstructure:"carbRatio" pg:"carb_ratio" json:"carbRatio,omitempty"`
	InsulinSensitivity   []interface{}      `mapstructure:"insulinSensitivity" pg:"insulin_sensitivity" json:"insulinSensitivity,omitempty"`


}

func DecodeSettings(data interface{}) (*Settings, mapstructure.Metadata, error)  {
	var settings = Settings{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &settings,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding settings: ", err)
			return nil, metadata, err
		}

		if err := settings.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &settings, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

