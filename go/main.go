package main

import (
	"strconv"
	"syscall/js"

	"github.com/golang/geo/s2"
)

var document js.Value

func init() {
	document = js.Global().Get("document")
}

func convertToS2CellIDUINT(this js.Value, inputs []js.Value) interface{} {
	latStr := inputs[0].String()
	lngStr := inputs[1].String()
	levelStr := inputs[2].String()

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return "Invalid latitude"
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return "Invalid longitude"
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		return "Invalid level"
	}

	ll := s2.LatLngFromDegrees(lat, lng)
	return strconv.FormatUint(uint64(s2.CellIDFromLatLng(ll).Parent(level)), 10)
}

func convertToS2CellIDINT(this js.Value, inputs []js.Value) interface{} {
	latStr := inputs[0].String()
	lngStr := inputs[1].String()
	levelStr := inputs[2].String()

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return "Invalid latitude"
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return "Invalid longitude"
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		return "Invalid level"
	}

	ll := s2.LatLngFromDegrees(lat, lng)
	return strconv.Itoa(int(s2.CellIDFromLatLng(ll).Parent(level)))
}

func convertToS2CellToken(this js.Value, inputs []js.Value) interface{} {
	latStr := js.ValueOf(inputs[0]).String()
	latStr = inputs[0].String()
	lngStr := inputs[1].String()
	levelStr := inputs[2].String()

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return "Invalid latitude"
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return "Invalid longitude"
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		return "Invalid level"
	}

	ll := s2.LatLngFromDegrees(lat, lng)
	return s2.CellIDFromLatLng(ll).Parent(level).ToToken()
}

func registerCallbacks() {
	js.Global().Set("convert_to_s2_cell_id_uint", js.FuncOf(convertToS2CellIDUINT))
	js.Global().Set("convert_to_s2_cell_id_int", js.FuncOf(convertToS2CellIDINT))
	js.Global().Set("convert_to_s2_cell_token", js.FuncOf(convertToS2CellToken))
}

func main() {
	c := make(chan struct{}, 0)
	println("registering go callbacks")
	registerCallbacks()
	<-c
}
