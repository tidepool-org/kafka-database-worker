package models

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ClinicsClinicians struct {
	ClinicId          string    `mapstructure:"clinicid" pg:"clinic_id"`
	ClinicianId          string    `mapstructure:"clinicianid" pg:"clinician_id"`

	active   bool    `mapstructure:"active" pg:"active"`
}

func DecodeClinicsClinicians(data interface{}) (*ClinicsClinicians, mapstructure.Metadata, error) {
	var clinicsClinicians = ClinicsClinicians{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &clinicsClinicians,
		Metadata: &metadata,
	} ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding clinis: ", err)
			return nil, metadata, err
		}

		if clinicsClinicians.ClinicId == "" || clinicsClinicians.ClinicianId == "" {
			//fmt.Println("clinicID or cliniciansId is null ")
			return nil, metadata, errors.New("clinicid or cliniciansid is null")

		}

		return &clinicsClinicians, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

func (u *ClinicsClinicians) GetType() string {
	return "clinicsClinicians"
}
