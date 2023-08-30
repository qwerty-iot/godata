package main

import (
	"github.com/qwerty-iot/godata"
	"github.com/qwerty-iot/godata/edmx"
	"github.com/qwerty-iot/tox"
	"time"
)

type StationHandler struct {
}

func (h *StationHandler) List(s *godata.Server, scope string, options *godata.ListOptions) []any {
	var arr []any
	station := tox.NewObject(nil)
	station.Set("Name", "Station 1")
	station.Set("Temperature", 25.0)
	station.Set("Humidity", 50.0)
	arr = append(arr, station)
	return arr
}

func main() {
	var model *edmx.Model

	godata.SetLogLevel(godata.LogLevelDebug)

	model = edmx.New("http://localhost:8200", "Weather")
	e := model.AddEntityType("Station", nil)
	e.SetKey("Name")
	e.AddProperty("Name", edmx.TypeString, nil)
	e.AddProperty("Temperature", edmx.TypeDouble, nil)
	e.AddProperty("Humidity", edmx.TypeDouble, nil)
	ec := model.AddEntityContainer("Container")
	ec.AddEntitySet("Stations", "Station")

	server := godata.NewServer("http://localhost:8200", model)
	server.SetHandler("Stations", &StationHandler{})
	server.Start(":8200", nil)

	time.Sleep(time.Hour * 60)
}
