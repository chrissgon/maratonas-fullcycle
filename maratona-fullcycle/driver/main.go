package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func getDrivers() interface{} {
	file, _ := ioutil.ReadFile("drivers.json")

	var drivers interface{}

	json.Unmarshal([]byte(file), &drivers)

	return drivers
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/drivers", func(r render.Render) {
		r.JSON(200, getDrivers())
	})

	m.Run()
}
