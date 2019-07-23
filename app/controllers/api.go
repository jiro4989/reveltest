package controllers

import (
	"time"

	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

type Date struct {
	Today string `json:"today"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}

func (c Api) Today() revel.Result {
	now := time.Now()
	today := now.Format("2006/01/02")
	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	resp := Date{
		Today: today,
		Year:  year,
		Month: month,
		Day:   day,
	}
	return c.RenderJSON(resp)
}
