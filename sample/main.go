package main

import (
	"github.com/qwerty-iot/godata"
	"github.com/qwerty-iot/godata/edmx"
	"github.com/qwerty-iot/tox"
	"time"
)

type StationHandler struct {
}

func (h *StationHandler) List(s *godata.Server, scope string, options *godata.ListOptions) ([]any, string, error) {
	var arr []any
	station := tox.NewObject(nil)
	station.Set("Name", "Station 1")
	station.Set("Temperature", 25.0)
	station.Set("Humidity", 50.0)
	station.Set("Region", "North")
	arr = append(arr, station)

	station = tox.NewObject(nil)
	station.Set("Name", "Station 2")
	station.Set("Temperature", 28.0)
	station.Set("Humidity", 48.0)
	station.Set("Region", "South")
	arr = append(arr, station)

	station = tox.NewObject(nil)
	station.Set("Name", "Station 3")
	station.Set("Temperature", 27.0)
	station.Set("Humidity", 42.0)
	station.Set("Region", "North")
	arr = append(arr, station)
	return arr, "", nil
}

type RegionHandler struct {
}

func (h *RegionHandler) List(s *godata.Server, scope string, options *godata.ListOptions) ([]any, string, error) {
	var arr []any
	region := tox.NewObject(nil)
	region.Set("Name", "North")
	arr = append(arr, region)

	region = tox.NewObject(nil)
	region.Set("Name", "South")
	arr = append(arr, region)

	region = tox.NewObject(nil)
	region.Set("Name", "East")
	arr = append(arr, region)

	region = tox.NewObject(nil)
	region.Set("Name", "West")
	arr = append(arr, region)

	return arr, "", nil
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
	e.AddProperty("Region", edmx.TypeString, nil)
	e.AddNavigationProperty("Region", "Region", false, "Stations")

	e = model.AddEntityType("Region", nil)
	e.SetKey("Name")
	e.AddProperty("Name", edmx.TypeString, nil)

	ec := model.AddEntityContainer("Container")
	ec.AddEntitySet("Stations", "Station")
	ec.AddEntitySet("Regions", "Region")

	server := godata.NewServer("http://localhost:8200", model)
	server.SetHandler("Stations", &StationHandler{})
	server.SetHandler("Regions", &RegionHandler{})
	server.Start(":8200", nil)

	time.Sleep(time.Hour * 60)
}
