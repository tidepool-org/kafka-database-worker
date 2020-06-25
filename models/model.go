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

func DecodeModel(data interface{}) (*string, interface{}) {
	var baseModel BaseModel
	if err := mapstructure.Decode(data, &baseModel); err != nil {
		fmt.Println("Problem decoding base model")
		return nil, nil
	}
	var model interface{}
	if baseModel.Active {
		Active += 1
		switch baseModel.Type {
		case "upload":
			model = DecodeUpload(data)
			//return upload
		case "basal":
			model = DecodeBasal(data)
		case "bolus":
			model = DecodeBolus(data)
		case "cbg":
			model = DecodeCbg(data)
		case "smbg":
			model = DecodeSmbg(data)
		case "wizard":
			model = DecodeWizard(data)
		case "food":
			model = DecodeFood(data)
		case "deviceEvent":
			model = DecodeDeviceEvent(data)
		case "pumpSettings":
			model = DecodePumpSettings(data)
		case "physicalActivity":
			model = DecodePhysicalActivity(data)
		default:
			fmt.Println("Currently not handling type: ", baseModel.Type)
			return nil, nil
		}
	} else {
		Inactive += 1
		return nil, nil
	}
	return &baseModel.Type, model
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