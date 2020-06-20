package models

import (
	"github.com/mitchellh/mapstructure"
	"fmt"
	"strings"
	"time"
	"reflect"
)

var Active = 0
var Inactive = 0

type BaseModel struct {
	Type      string `mapstructure:"type"`
	Active    bool `mapstructure:"_active"`
}

func DecodeModel(data interface{}) interface{} {
	var baseModel BaseModel
	if err := mapstructure.Decode(data, &baseModel); err != nil {
		fmt.Println("Problem decoding base model")
		return nil
	}
	if baseModel.Active {
		Active += 1
		switch baseModel.Type {
		case "upload":
			upload := DecodeUpload(data)
			return upload
		case "basal":
			basal := DecodeBasal(data)
			return basal
		case "bolus":
			bolus := DecodeBolus(data)
			return bolus
		case "cbg":
			cbg := DecodeCbg(data)
			return cbg
		case "smbg":
			smbg := DecodeSmbg(data)
			return smbg
		case "wizard":
			wizard := DecodeWizard(data)
			return wizard
		case "food":
			food := DecodeFood(data)
			return food
		case "deviceEvent":
			food := DecodeDeviceEvent(data)
			return food
		case "pumpSetting":
			food := DecodePumpSettings(data)
			return food
		default:
			fmt.Println("Currently not handling type: ", baseModel.Type)
		}
	} else {
		Inactive += 1
	}
	return nil
}

// StringToTimeHookFuncTimezoneOptional returns a DecodeHookFunc that converts
// strings to time.Time.  If time does not have a timezone - appends a Z for UTC timezone
func StringToTimeHookFuncTimezoneOptional(layout string) mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		// Convert it by parsing
		s := data.(string)
		if !strings.Contains(s, "Z") && !strings.Contains(s, "+") {
			s += "Z"
		}
		return time.Parse(layout, s)
	}
}