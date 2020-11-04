package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Upload struct {
	Base                    `mapstructure:",squash"`

	DataSetType          string    `mapstructure:"dataSetType" pg:"data_set_type"`
	DataState            string    `mapstructure:"_dataState" pg:"data_state"`

	DeviceManufacturers  []string  `mapstructure:"deviceManufacturers" pg:"device_manufacturers,array" json:"deviceManufacturers,omitempty"`
	DeviceModel          string    `mapstructure:"deviceModel" pg:"device_model" json:"deviceModel,omitempty"`
	DeviceSerialNumber   string    `mapstructure:"deviceSerialNumber" pg:"device_serial_number" json:"deviceSerialNumber,omitempty"`

	DeviceTags           []string    `mapstructure:"deviceTags" pg:"device_tags,array" json:"deviceTags,omitempty"`

	State                string    `mapstructure:"_state" pg:"state"`
	Version              string    `mapstructure:"version" pg:"version"`
}

func DecodeUpload(data interface{}) (*Upload, mapstructure.Metadata, error) {
	var upload = Upload{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: StringToTimeHookFuncTimezoneOptional(time.RFC3339),
		Result: &upload,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding upload: ", err)
			return nil, metadata, err
		}

		if err := upload.DecodeBase(); err != nil {
			fmt.Println("Error encoding base json: ", err)
			return nil, metadata, err
		}


		return &upload, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, nil
	}
}
