package models

import (
	"github.com/mitchellh/mapstructure"
	"time"
	"fmt"
)

type Upload struct {
	CreatedTime    time.Time `mapstructure:"createdTime" pg:"created_time type:timestamptz"`
	DeviceId       string    `mapstructure:"deviceId" pg:"device_id"`
	Id             string    `mapstructure:"id" pg:"id"`

	Time           time.Time `mapstructure:"time" pg:"time"`
	Timezone       string    `mapstructure:"timezone" pg:"timezone"`

	UploadId       string    `mapstructure:"uploadId" pg:"upload_id"`
	UserId         string    `mapstructure:"_userId" pg:"user_id"`

	DataSetType    string    `mapstructure:"dataSetType" pg:"data_set_type"`
	DataState      string    `mapstructure:"_dataState" pg:"data_state"`

	DeviceSerialNumber string `mapstructure:"deviceSerialNumber" pg:"device_serial_number"`
	State          string    `mapstructure:"_state" pg:"state"`
	Version        string    `mapstructure:"version" pg:"version"`
	ModifiedTime   time.Time `mapstructure:"modifiedTime" pg:"modified_time"`
	Revision       int64     `mapstructure:"revision" pg:"revision"`
}

func DecodeUpload(data interface{}) *Upload {
	var upload = Upload{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
		Result: &upload,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			fmt.Println("Error decoding: ", err)
		} else {
			return &upload
		}

	} else {
		fmt.Println("Can not create decoder: ", err)
	}
	return nil
}
