package models

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ClinicsPatients struct {
	ClinicId          string    `mapstructure:"clinicId" pg:"clinic_id"`
	PatientId          string    `mapstructure:"patientId" pg:"patient_id"`

	active   bool    `mapstructure:"active" pg:"active"`
}

func DecodeClinicsPatients(data interface{}) (*ClinicsPatients, mapstructure.Metadata, error) {
	var patients = ClinicsPatients{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &patients,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding clinis: ", err)
			return nil, metadata, err
		}

		if patients.ClinicId == "" || patients.PatientId == "" {
			//fmt.Println("clinicID or patientID is null ")
			return nil, metadata, errors.New("clinicid or patientid is null")

		}

		return &patients, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

func (u *ClinicsPatients) GetType() string {
	return "clinicsPatients"
}

