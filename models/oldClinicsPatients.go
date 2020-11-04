package models

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type OldClinicsPatients struct {
	OldClinicId          string    `mapstructure:"userId" pg:"old_clinic_id"`
	PatientId          string      `mapstructure:"sharerId" pg:"patient_id"`

}

func DecodeOldClinicsPatients(data interface{}) (*OldClinicsPatients, mapstructure.Metadata, error) {
	var patients = OldClinicsPatients{}
	var metadata = mapstructure.Metadata{}

	if decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &patients,
		Metadata: &metadata,
	   } ); err == nil {
		if err := decoder.Decode(data); err != nil {
			//fmt.Println("Error decoding clinis: ", err)
			return nil, metadata, err
		}

		if patients.OldClinicId == "" || patients.PatientId == "" {
			//fmt.Println("clinicID or patientID is null ")
			return nil, metadata, errors.New("clinicid or patientid is null")

		}

		return &patients, metadata, nil

	} else {
		fmt.Println("Can not create decoder: ", err)
		return nil, metadata, err
	}
}

func (u *OldClinicsPatients) GetType() string {
	return "oldClinicsPatients"
}

