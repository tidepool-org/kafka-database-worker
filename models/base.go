package models

import (
	"encoding/json"
	"time"
	"fmt"
)

type Model interface {
	GetType() string
}

type Base struct {
	Time              time.Time  `mapstructure:"time" pg:"time type:timestamptz"`

	Type              string     `mapstructure:"type" pg:"-"`

	ArchivedTime      time.Time  `mapstructure:"archivedTime" pg:"archived_time type:timestamptz"`
	CreatedTime       time.Time  `mapstructure:"createdTime" pg:"created_time type:timestamptz"`
	ModifiedTime      time.Time  `mapstructure:"modifiedTime" pg:"modified_time type:timestamptz"`
	DeviceTime        time.Time  `mapstructure:"deviceTime" pg:"device_time type:timestamptz"`

	DeviceId          string     `mapstructure:"deviceId,omitempty" pg:"device_id"`
	Id                string     `mapstructure:"id,omitempty" pg:"id"`
	Guid              string     `mapstructure:"guid,omitempty" pg:"guid"`

	Timezone          string     `mapstructure:"timezone,omitempty" pg:"timezone"`
	TimezoneOffset    int64      `mapstructure:"timezoneOffset,omitempty" pg:"timezone_offset"`
	ClockDriftOffset  int64      `mapstructure:"clockDriftOffset,omitempty" pg:"clock_drift_offset"`
	ConversionOffset  int64      `mapstructure:"conversionOffset,omitempty" pg:"conversion_offset"`

	UploadId          string     `mapstructure:"uploadId,omitempty" pg:"upload_id"`
	UserId            string     `mapstructure:"_userId,omitempty" pg:"user_id"`

	BgTargetMap             []interface{}      `mapstructure:"bgTarget" pg:"-"`
	BgTargetJson            string                      `pg:"bg_target"`

	PayloadMap        []interface{}      `mapstructure:"payload" pg:"-"`
	PayloadJson       string     `pg:"payload"`
	OriginMap         []interface{}      `mapstructure:"origin" pg:"-"`
	OriginJson        string     `pg:"origin"`

	Active            bool       `mapstructure:"_active" pg:"active"`

	Revision          int64      `mapstructure:"revision,omitempty" pg:"revision"`
}

func (b *Base) DecodeBase() error {
	payloadByteArray, err := json.Marshal(b.PayloadMap)
	b.PayloadJson = string(payloadByteArray)
	if err != nil {
		return err
	}

	originByteArray, err := json.Marshal(b.OriginMap)
	b.OriginJson = string(originByteArray)
	if err != nil {
		return err
	}
	return nil
}

func (b *Base) GetType() string {
	return b.Type
}

func (b *Base) GetUserId() string {
	return b.UserId
}


