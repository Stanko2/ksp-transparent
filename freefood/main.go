package freefood

import (
	"time"
)

type DailyMenu struct {
	soup Food
	main []Food
	date time.Time
}

type Food struct {
	Name string
	Price float64
}