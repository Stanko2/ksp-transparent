package freefood

import (
	"time"
)

type DailyMenu struct {
	menus []Food
	date time.Time
}

type Restaurant int16 

const (
	freefood Restaurant = 0
	faynfood Restaurant = 1
)

type Menu struct {
	menus map[time.Weekday]DailyMenu
	restaurant Restaurant
}

type FoodType int16

const (
	Soup FoodType = 1
	Main FoodType = 0
)

type Food struct {
	Type FoodType
	Name string
	Price float64
}