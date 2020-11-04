package models

import (
	"github.com/mitchellh/mapstructure"
	"encoding/json"
	"time"
	"fmt"
	"reflect"
	"github.com/fatih/structtag"
)

type Model interface {
	GetType() string
}

type Base struct {
	Time              time.Time  `mapstructure:"time" pg:"time,type:timestamptz" json:"time,omitempty"`

	Type              string     `mapstructure:"type" pg:"-" json:"type,omitempty"`

	InternalMongoId   string    `pg:"internal_mongo_id" json:"_id"`

	ArchivedTime      time.Time  `mapstructure:"archivedTime" pg:"archived_time,type:timestamptz" json:"-"`
	CreatedTime       time.Time `mapstructure:"createdTime" pg:"created_time,type:timestamptz" json:"-"`
	ModifiedTime      time.Time `mapstructure:"modifiedTime" pg:"modified_time,type:timestamptz" json:"-"`
	DeviceTime        time.Time      `mapstructure:"deviceTime" pg:"device_time,type:timestamptz" json:"-"`

	DeviceId          string   `mapstructure:"deviceId,omitempty" pg:"device_id" json:"deviceId,omitempty"`
	Id                string   `mapstructure:"id,omitempty" pg:"id" json:"id,omitempty"`
	Guid              string     `mapstructure:"guid,omitempty" pg:"guid" json:"guid,omitempty"`


	Timezone          string   `mapstructure:"timezone,omitempty" pg:"timezone" json:"timezone,omitempty"`
	TimezoneOffset    int64    `mapstructure:"timezoneOffset,omitempty" pg:"timezone_offset" json:"timezoneOffset,omitempty"`
	ClockDriftOffset  int64    `mapstructure:"clockDriftOffset,omitempty" pg:"clock_drift_offset" json:"clockDriftOffset,omitempty"`
	ConversionOffset  int64    `mapstructure:"conversionOffset,omitempty" pg:"conversion_offset" json:"conversionOffset,omitempty"`

	UploadId          string   `mapstructure:"uploadId,omitempty" pg:"upload_id" json:"uploadId,omitempty"`
	UserId            string   `mapstructure:"_userId,omitempty" pg:"user_id" json:"-"`

	Payload        map[string]interface{}      `mapstructure:"payload" pg:"payload" json:"uploadId,omitempty"`
	Origin         map[string]interface{}      `mapstructure:"origin" pg:"origin" json:"-"`
	Annotations    []interface{}      `mapstructure:"annotations" pg:"annotations" json:"annotations,omitempty"`

	Active            bool       `mapstructure:"_active" pg:"active" json:"-"`

	Revision          int64      `mapstructure:"revision,omitempty" pg:"revision" json:"revision"`

	Remaining       map[string]interface{}   `mapstructure:",remain" pg:"-" json:"-"`
}

func GetMongoIdFromFilterField(filter_field interface{}) *string {
	filter_field_string := fmt.Sprintf("%v", filter_field)
	var rec map[string]interface{}
	if err := json.Unmarshal([]byte(filter_field_string), &rec); err != nil {
		fmt.Println("Get Mongo Id Error Unmarshalling", err)
		return nil
	}
	return GetMongoId(rec["_id"])
}

func GetMongoId(rawMongoId interface{}) *string {

	type Oid struct {
		Oid       string   `mapstructure:"$oid"`
	}

	mongoId, ok := rawMongoId.(string)
	if ok {
		return &mongoId
	} else {
		var oid Oid
		if err := mapstructure.Decode(rawMongoId, &oid); err != nil {
			return nil
		}
		return &oid.Oid
	}
}



func (b *Base) DecodeBase() error {
	mongoIdField, ok := b.Remaining["_id"]
	if ok {
		b.InternalMongoId = *GetMongoId(mongoIdField)
	}
	return nil
}

func (b *Base) GetType() string {
	return b.Type
}

func (b *Base) GetUserId() string {
	return b.UserId
}

var modelDefs []interface{}
func GetModels() []interface{} {

	if modelDefs == nil {

		modelDefs = append(modelDefs, (*Cbg)(nil))
		modelDefs = append(modelDefs, (*Basal)(nil))
		modelDefs = append(modelDefs, (*Smbg)(nil))
		modelDefs = append(modelDefs, (*Upload)(nil))

		modelDefs = append(modelDefs, (*BloodKetone)(nil))
		modelDefs = append(modelDefs, (*Bolus)(nil))
		modelDefs = append(modelDefs, (*CgmSettings)(nil))
		modelDefs = append(modelDefs, (*DeviceEvent)(nil))
		modelDefs = append(modelDefs, (*DeviceMeta)(nil))
		modelDefs = append(modelDefs, (*DosingDecision)(nil))
		modelDefs = append(modelDefs, (*Food)(nil))
		modelDefs = append(modelDefs, (*Insulin)(nil))
		modelDefs = append(modelDefs, (*PhysicalActivity)(nil))
		modelDefs = append(modelDefs, (*PumpSettings)(nil))
		modelDefs = append(modelDefs, (*ReportedState)(nil))
		modelDefs = append(modelDefs, (*Settings)(nil))
		modelDefs = append(modelDefs, (*Wizard)(nil))
	}

	return modelDefs
}

func GetTagMap(s interface{}) map[string]string {

	var tagMap map[string]string
	tagMap = make(map[string]string)
	e := reflect.Indirect(reflect.ValueOf(s))

	for i := 0; i < e.NumField(); i++ {
		tag := e.Type().Field(i).Tag
		tags, err := structtag.Parse(string(tag))
		if err != nil {
			continue
		}
		pgTag, _ := tags.Get("pg")
		msTag, _ := tags.Get("mapstructure")
		if pgTag != nil  && msTag != nil {
			name := pgTag.Name
			msName := msTag.Name
			if !(name == "-" || name == "" || msName == "") {
				tagMap[msName] = name
			}
		}
	}
	return tagMap
}

type ModelType struct {
	Type string
	TagMap map[string]string
}
var types []ModelType
func GetModelTypes() []ModelType {
	GetTagMap(&Base{})
	GetTagMap(&Cbg{})
	if types == nil {

		baseTagMap := GetTagMap(&Base{})
		types = append(types, ModelType{Type: "cbg", TagMap: combineMaps(baseTagMap, GetTagMap(&Cbg{}))})
		types = append(types, ModelType{Type: "basal", TagMap: combineMaps(baseTagMap, GetTagMap(&Basal{}))})
		types = append(types, ModelType{Type: "smbg", TagMap: combineMaps(baseTagMap, GetTagMap(&Smbg{}))})
		types = append(types, ModelType{Type: "upload", TagMap: combineMaps(baseTagMap, GetTagMap(&Upload{}))})

		types = append(types, ModelType{Type: "bloodKetone", TagMap: combineMaps(baseTagMap, GetTagMap(&BloodKetone{}))})
		types = append(types, ModelType{Type: "bolus", TagMap: combineMaps(baseTagMap, GetTagMap(&Bolus{}))})
		types = append(types, ModelType{Type: "cgmSettings", TagMap: combineMaps(baseTagMap, GetTagMap(&CgmSettings{}))})
		types = append(types, ModelType{Type: "deviceEvent", TagMap: combineMaps(baseTagMap, GetTagMap(&DeviceEvent{}))})
		types = append(types, ModelType{Type: "deviceMeta", TagMap: combineMaps(baseTagMap, GetTagMap(&DeviceMeta{}))})
		types = append(types, ModelType{Type: "dosingDecision", TagMap: combineMaps(baseTagMap, GetTagMap(&DosingDecision{}))})
		types = append(types, ModelType{Type: "food", TagMap: combineMaps(baseTagMap, GetTagMap(&Food{}))})
		types = append(types, ModelType{Type: "insulin", TagMap: combineMaps(baseTagMap, GetTagMap(&Insulin{}))})
		types = append(types, ModelType{Type: "physicalActivity", TagMap: combineMaps(baseTagMap, GetTagMap(&PhysicalActivity{}))})
		types = append(types, ModelType{Type: "pumpSettings", TagMap: combineMaps(baseTagMap, GetTagMap(&PumpSettings{}))})
		types = append(types, ModelType{Type: "reportedState", TagMap: combineMaps(baseTagMap, GetTagMap(&ReportedState{}))})
		types = append(types, ModelType{Type: "settings", TagMap: combineMaps(baseTagMap, GetTagMap(&Settings{}))})
		types = append(types, ModelType{Type: "wizard", TagMap: combineMaps(baseTagMap, GetTagMap(&Wizard{}))})

	}
	return types
}

func combineMaps(a map[string]string, b map[string]string) map[string]string {
	var c map[string]string
	c = make(map[string]string)
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		c[k] = v
	}
	return c
}

