package models

import (
	"github.com/mitchellh/mapstructure"
	"fmt"

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
		default:
			fmt.Println("Currently not handling type: ", baseModel.Type)
		}
	} else {
		Inactive += 1
	}
	return nil
}