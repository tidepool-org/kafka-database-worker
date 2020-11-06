package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"time"
)

type PumpSettings struct {
	Base                                             `mapstructure:",squash"`

	ActiveSchedule          string                      `mapstructure:"activeSchedule" pg:"active_schedule" json:"activeSchedule,omitempty"`

	BasalSchedules       map[string]interface{}      `mapstructure:"basalSchedules" pg:"basal_schedules" json:"basalSchedules,omitempty"`

	BgTargets             map[string]interface{}      `mapstructure:"bgTargets" pg:"bg_targets" json:"bgTargets,omitempty"`

	CarbRatios            map[string]interface{}      `mapstructure:"carbRatios" pg:"carb_ratios" json:"carbRatios,omitempty"`

	InsulinSensitivities   map[string]interface{}      `mapstructure:"insulinSensitivities" pg:"insulin_sensitivities" json:"insulinSensitivities,omitempty"`

	Units                map[string]interface{}      `mapstructure:"units" pg:"units" json:"units,omitempty"`

	Manufacturers        []string                     `mapstructure:"manufacturers" pg:"manufacturers,array" json:"manufacturers,omitempty"`
	Model                string                       `mapstructure:"model" pg:"model" json:"model,omitempty"`
	SerialNumber         string                       `mapstructure:"serialNumber" pg:"serial_number" json:"serialNumber,omitempty"`
}

type PumpSettings2 struct {
	tableName struct{}                              `pg:"pump_settings"`

	Base                                             `mapstructure:",squash"`

	ActiveSchedule          string                      `mapstructure:"activeSchedule" pg:"active_schedule" json:"activeSchedule,omitempty"`

	BasalSchedules       map[string]interface{}      `mapstructure:"basalSchedules" pg:"basal_schedules" json:"basalSchedules,omitempty"`

	BgTargets             []interface{}      `mapstructure:"bgTargets" pg:"bg_targets" json:"bgTargets,omitempty"`

	CarbRatios            []interface{}      `mapstructure:"carbRatio" pg:"carb_ratios" json:"carbRatios,omitempty"`

	InsulinSensitivities   []interface{}      `mapstructure:"insulinSensitivities" pg:"insulin_sensitivities" json:"insulinSensitivities,omitempty"`

	Units                map[string]interface{}      `mapstructure:"units" pg:"units" json:"units,omitempty"`

	Manufacturers        []string                     `mapstructure:"manufacturers" pg:"manufacturers,array" json:"manufacturers,omitempty"`
	Model                string                       `mapstructure:"model" pg:"model" json:"model,omitempty"`
	SerialNumber         string                       `mapstructure:"serialNumber" pg:"serial_number" json:"serialNumber,omitempty"`
}

func DecodePumpSettings(data interface{}) (*PumpSettings, mapstructure.Metadata, error) {
	var pumpSettings = PumpSettings{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &pumpSettings,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding pump settings: ", err)
			return nil, metadata, err
		}

		if err := pumpSettings.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &pumpSettings, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

func DecodePumpSettings2(data interface{}) (*PumpSettings2, mapstructure.Metadata, error) {
	var pumpSettings = PumpSettings2{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &pumpSettings,
		Metadata: &metadata,
	} ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding pump settings: ", err)
			return nil, metadata, err
		}

		if err := pumpSettings.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}
		pumpSettings.Type = "pumpSettings2"

		return &pumpSettings, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}
