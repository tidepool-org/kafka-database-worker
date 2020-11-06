package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type DosingDecision struct {
	Base                      `mapstructure:",squash"`

	InsulinOnBoard                map[string]interface{}      `mapstructure:"insulinOnBoard" pg:"insulin_on_board" json:"insulinOnBoard,omitempty"`
	CarbohydratesOnBoard          map[string]interface{}      `mapstructure:"carbohydratesOnBoard" pg:"carbohydrates_on_board" json:"carbohydratesOnBoard,omitempty"`

	BloodGlucoseTargetSchedule    []interface{}      `mapstructure:"bloodGlucoseTargetSchedule" pg:"blood_glucose_target_schedule" json:"bloodGlucoseTargetSchedule,omitempty"`
	BloodGlucoseForecast          []interface{}      `mapstructure:"bloodGlucoseForecast" pg:"blood_glucose_forecast" json:"bloodGlucoseForecast,omitempty"`
	BloodGlucoseForecastIncludingPendingInsulin          []interface{}      `mapstructure:"bloodGlucoseForecastIncludingPendingInsulin" pg:"blood_glucose_forecast_including_pending_insulin" json:"bloodGlucoseForecastIncludingPendingInsulin,omitempty"`

	RecommendedBasal              map[string]interface{}      `mapstructure:"recommendedBasal" pg:"recommended_basal" json:"recommendedBasal,omitempty"`
	RecommendedBolus              map[string]interface{}      `mapstructure:"recommendedBolus" pg:"recommended_bolus" json:"recommendedBolus,omitempty"`
	Units                         map[string]interface{}      `mapstructure:"units" pg:"units" json:"units,omitempty"`

}

func DecodeDosingDecision(data interface{}) (*DosingDecision, mapstructure.Metadata, error)  {
	var dosingDecision = DosingDecision{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &dosingDecision,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding dosingDecision: ", err)
			return nil, metadata, err
		}

		if err := dosingDecision.DecodeBase(); err != nil {
			//fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}

		return &dosingDecision, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

