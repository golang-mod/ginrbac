package datatypes

import (
	//"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type Date string

func (date *Date) Scan(value interface{}) (err error) {
	datetime, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to time.Time value:", value))
	}
	*date = Date(datetime.Format("2006-01-02"))
	return nil
}

func (date Date) Value() (driver.Value, error) {
	return date, nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string {
	return "date"
}
