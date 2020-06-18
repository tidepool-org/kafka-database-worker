package models

import (
	"github.com/mitchellh/mapstructure"
	"fmt"

)

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
		switch baseModel.Type {
		case "upload":
			fmt.Println("Decoding upload")
			upload := DecodeUpload(data)
			return upload
		case "basal":
			fmt.Println("Decoding basal")
			basal := DecodeBasal(data)
			return basal
		default:
			fmt.Println("Currently not handling type: ", baseModel.Type)
		}
	} else {
		fmt.Println("Base Model not active")
	}
	return nil
}