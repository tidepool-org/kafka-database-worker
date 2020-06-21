package models

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type PumpSettings struct {
	Base                                             `mapstructure:",squash"`

	ActiveSchedule          string                      `mapstructure:"activeSchedule" pg:"active_schedule"`

	BasalSchedulesMap       map[string]interface{}      `mapstructure:"nutrition" pg:"-"`
	BasalSchedulesJson      string                      `pg:"nutrition"`

	BgTargetMap             map[string][]interface{}      `mapstructure:"bgTarget" pg:"-"`
	BgTargetJson            string                      `pg:"bg_target"`

	CarbRatioMap            map[string][]interface{}      `mapstructure:"carbRatio" pg:"-"`
	CarbRatioJson           string                      `pg:"carb_ratio"`

	InsulinSensitivityMap   map[string][]interface{}      `mapstructure:"insulinSensitivity" pg:"-"`
	InsulinSensitivityJson  string                      `pg:"insulin_sensitivity"`

	unitsMap                map[string]interface{}      `mapstructure:"units" pg:"-"`
	unitsJson               string                      `pg:"units"`
}

func DecodePumpSettings(data interface{}) *PumpSettings {
	var pumpSettings = PumpSettings{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &pumpSettings,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
			return nil
		}

		basalSchedulesByteArray, err := json.Marshal(pumpSettings.BasalSchedulesMap)
		pumpSettings.BasalSchedulesJson = string(basalSchedulesByteArray)
		if err != nil {
			fmt.Println("Error encoding Basal Schedules json: ", err)
			return nil
		}

		bgTargetMapByteArray, err := json.Marshal(pumpSettings.BgTargetMap)
		pumpSettings.BgTargetJson = string(bgTargetMapByteArray)
		if err != nil {
			fmt.Println("Error encoding Bg Target json: ", err)
			return nil
		}

		carbRatioMapByteArray, err := json.Marshal(pumpSettings.CarbRatioMap)
		pumpSettings.CarbRatioJson = string(carbRatioMapByteArray)
		if err != nil {
			fmt.Println("Error encoding carb ration json: ", err)
			return nil
		}

		insulinSensitivityByteArray, err := json.Marshal(pumpSettings.InsulinSensitivityMap)
		pumpSettings.InsulinSensitivityJson = string(insulinSensitivityByteArray)
		if err != nil {
			fmt.Println("Error encoding insulin sensitivity json: ", err)
			return nil
		}

		unitsMapByteArray, err := json.Marshal(pumpSettings.unitsMap)
		pumpSettings.unitsJson = string(unitsMapByteArray)
		if err != nil {
			fmt.Println("Error encoding units json: ", err)
			return nil
		}

		return &pumpSettings

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
