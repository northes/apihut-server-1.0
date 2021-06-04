package server

import (
	"apihut-server/model"
	"apihut-server/util"
	"fmt"
)

func GetWeather(p *model.Weather) (interface{}, error) {
	var r model.GaoDeWeather
	err := util.HttpGetRequest(fmt.Sprintf(
		"https://restapi.amap.com/v3/weather/weatherInfo?parameters&key=b9bd34580b7133934c40a831703cc3fb&city=%s&extensions=%s",
		p.City,
		p.Extensions,
	), &r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if r.Status == "0" {
		return nil, util.ErrRequest
	}
	if len(r.Lives) == 0 {
		return r.Forecasts, nil
	}
	return r.Lives, nil
}
