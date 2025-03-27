package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [3]string{"Doing", "Done", "Deleted"}

func (item ItemStatus) String() string {
	return allItemStatuses[item]
}

func parseItemStatus(status string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if status == allItemStatuses[i] {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("Invalid item status string")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintln("Failed to scan data from sql:", value))
	}

	v, err := parseItemStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to scan data from sql:", value))
	}
	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue
	return nil
}
