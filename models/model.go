package models

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strings"
	"time"
)

const DeviceDataCollection = "deviceData"
const UsersCollection = "users"
const ClinicsCollection = "clinic"
const ClinicsCliniciansCollection = "clinicsClinicians"
const ClinicsPatientsCollection = "clinicsPatients"
const PermsCollection = "perms"

type BaseDeviceModel struct {
	Type      string `mapstructure:"type"`
	Active    bool `mapstructure:"_active"`
}

func DecodeModel(data interface{}, topic string) (Model, mapstructure.Metadata, error) {
	switch {
	case strings.HasSuffix(topic, DeviceDataCollection):
		return DecodeDeviceModel(data)
	default:
		return DecodeGeneralModel(data, topic)
	}
}

func DecodeGeneralModel(data interface{}, topic string) (Model, mapstructure.Metadata, error) {
	switch {
	case strings.HasSuffix(topic, UsersCollection):
		user, metadata, err := DecodeUser(data)
		return user, metadata, err
	case strings.HasSuffix(topic, ClinicsCollection):
		user, metadata, err := DecodeClinics(data)
		return user, metadata, err
	case strings.HasSuffix(topic, ClinicsCliniciansCollection):
		user, metadata, err := DecodeClinicsClinicians(data)
		return user, metadata, err
	case strings.HasSuffix(topic, ClinicsPatientsCollection):
		user, metadata, err := DecodeClinicsPatients(data)
		return user, metadata, err
	case strings.HasSuffix(topic, PermsCollection):
		user, metadata, err := DecodeOldClinicsPatients(data)
		return user, metadata, err
	}
	fmt.Println("Could not decode.  Do not have a database for topic: ", topic)
	return nil, mapstructure.Metadata{}, nil
}

func DecodeDeviceModel(data interface{}) (Model, mapstructure.Metadata, error) {
	var baseDeviceModel BaseDeviceModel
	if err := mapstructure.Decode(data, &baseDeviceModel); err != nil {
		fmt.Println("Problem decoding base model", err)
		return nil, mapstructure.Metadata{}, err
	}
	return DecodeDeviceModelWithType(data, baseDeviceModel.Type)
}

func DecodeDeviceModelWithType(data interface{}, modelType string) (Model, mapstructure.Metadata, error) {
	switch modelType {
	case "basal":
		basal, metadata, err := DecodeBasal(data)
		return basal, metadata, err
	case "bloodKetone":
		bloodKetone, metadata, err := DecodeBloodKetone(data)
		return bloodKetone, metadata, err
	case "bolus":
		bolus, metadata, err := DecodeBolus(data)
		return bolus, metadata, err
	case "cbg":
		cbg, metadata, err := DecodeCbg(data)
		return cbg, metadata, err
	case "cgmSettings":
		cgmSettings, metadata, err := DecodeCgmSettings(data)
		return cgmSettings, metadata, err
	case "deviceEvent":
		deviceEvent, metadata, err := DecodeDeviceEvent(data)
		return deviceEvent, metadata, err
	case "deviceMeta":
		deviceMeta, metadata, err := DecodeDeviceMeta(data)
		return deviceMeta, metadata, err
	case "dosingDecision":
		dosingDecision, metadata, err := DecodeDosingDecision(data)
		return dosingDecision, metadata, err
	case "food":
		food, metadata, err := DecodeFood(data)
		return food, metadata, err
	case "insulin":
		insulin, metadata, err := DecodeInsulin(data)
		return insulin, metadata, err
	case "physicalActivity":
		physicalActivity, metadata, err := DecodePhysicalActivity(data)
		return physicalActivity, metadata, err
	case "pumpSettings":
		pumpSettings, metadata, err := DecodePumpSettings(data)
		// XXX This is somewhat of a hack.  Seems like data model is not consistent
		if err != nil {
			pumpSettings, metadata, err := DecodePumpSettings2(data)
			return pumpSettings, metadata, err
		}
		return pumpSettings, metadata, err
	case "reportedState":
		reportedState, metadata, err := DecodeReportedState(data)
		return reportedState, metadata, err
	case "settings":
		settings, metadata, err := DecodeSettings(data)
		return settings, metadata, err
	case "smbg":
		smbg, metadata, err := DecodeSmbg(data)
		return smbg, metadata, err
	case "upload":
		upload, metadata, err := DecodeUpload(data)
		return upload, metadata, err
	case "wizard":
		wizard, metadata, err := DecodeWizard(data)
		return wizard, metadata, err
	default:
		fmt.Println("Currently not handling type: ", modelType)
	}
	return nil, mapstructure.Metadata{}, nil
}

// StringToTimeHookFunusers returns a DecodeHookFunc that converts
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
		endStr := s[len(s)-9:]
		if !strings.Contains(s, "Z") && !(strings.Contains(endStr, "+") || strings.Contains(endStr, "-")) {
			s += "Z"
		}
		return time.Parse(layout, s)
	}
}


