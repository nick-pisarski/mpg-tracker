package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type FillUpEntity struct {
	gorm.Model
	Price      float32
	Gallons    float32
	Miles      float32
	DateFilled time.Time
}

func (f FillUpEntity) toString() string {
	items := []any{f.ID, f.Price, f.Gallons, f.Miles, f.DateFilled}
	return fmt.Sprintf("ID: %d\nPrice: %f\nGallons: %f\nMiles: %f\nDateFilled: %s", items...)
}
