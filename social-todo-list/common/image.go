package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"colunm:id;"`
	Url       string `json:"url" gorm:"colunm:url;"`
	Width     int    `json:"width" gorm:"colunm:width;"`
	Height    int    `json:"height" gorm:"colunm:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "image"
}

func (j *Image) Fulfill(domain string) {
	j.Url = fmt.Sprintf("%s/%s", domain, j.Url)
}

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal data from DB: %v", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
