package model

type Weather struct {
	City string `json:"city" form:"city" binding:"required"`
	Type string `json:"type" form:"type"`
}

type GaoDeWeather struct {
	Status    string     `json:"status"`
	Count     string     `json:"count"`
	Info      string     `json:"info"`
	Infocode  string     `json:"infocode"`
	Lives     []Life     `json:"lives,omitempty"`
	Forecasts []Forecast `json:"forecasts,omitempty"`
}

type Life struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	Reporttime    string `json:"reporttime"`
}

type Forecast struct {
	City       string `json:"city"`
	Adcode     string `json:"adcode"`
	Province   string `json:"province"`
	Reporttime string `json:"reporttime"`
	Casts      []Cast `json:"casts"`
}

type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	Dayweather   string `json:"dayweather"`
	Nightweather string `json:"nightweather"`
	Daytemp      string `json:"daytemp"`
	Nighttemp    string `json:"nighttemp"`
	Daywind      string `json:"daywind"`
	Nightwind    string `json:"nightwind"`
	Daypower     string `json:"daypower"`
	Nightpower   string `json:"nightpower"`
}
