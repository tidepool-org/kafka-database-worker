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

func DecodeModel(data interface{}) (Model, error) {
	var baseModel BaseModel
	if err := mapstructure.Decode(data, &baseModel); err != nil {
		fmt.Println("Problem decoding base model", err)
		return nil, err
	}
	if baseModel.Active {
		Active += 1
		switch baseModel.Type {
		case "upload":
			upload, err := DecodeUpload(data)
			return upload, err
		case "basal":
			basal, err := DecodeBasal(data)
			return basal, err
		case "bolus":
			bolus, err := DecodeBolus(data)
			return bolus, err
		case "cbg":
			cbg, err := DecodeCbg(data)
			return cbg, err
		case "smbg":
			smbg, err := DecodeSmbg(data)
			return smbg, err
		case "wizard":
			wizard, err := DecodeWizard(data)
			return wizard, err
		case "food":
			food, err := DecodeFood(data)
			return food, err
		case "deviceEvent":
			deviceEvent, err := DecodeDeviceEvent(data)
			return deviceEvent, err
		case "pumpSettings":
			pumpSettings, err := DecodePumpSettings(data)
			return pumpSettings, err
		case "physicalActivity":
			physicalActivity, err := DecodePhysicalActivity(data)
			return physicalActivity, err
		case "cgmSettings":
			cgmSettings, err := DecodeCgmSettings(data)
			return cgmSettings, err
		case "deviceMeta":
			deviceMeta, err := DecodeDeviceMeta(data)
			return deviceMeta, err
		default:
			fmt.Println("Currently not handling type: ", baseModel.Type)
		}
	} else {
		Inactive += 1
	}
	return nil, fmt.Errorf("Inactive")
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