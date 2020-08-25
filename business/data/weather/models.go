package weather

// Weather contains the weather data points captured from the API.
type Weather struct {
	ID       string  `json:"id,omitempty"`
	CityName string  `json:"city_name"`
	Desc     string  `json:"description"`
	Temp     float64 `json:"temp"`
	MinTemp  float64 `json:"temp_min"`
	MaxTemp  float64 `json:"temp_max"`
	Pressure int     `json:"pressure"`
}

type addResult struct {
	AddWeather struct {
		Weather []struct {
			ID string `json:"id"`
		} `json:"weather"`
	} `json:"addWeather"`
}

func (addResult) document() string {
	return `{
		weather {
			id
		}
	}`
}

type updateCityResult struct {
	UpdateCity struct {
		City []struct {
			ID string `json:"id"`
		} `json:"city"`
	} `json:"updateCity"`
}

func (updateCityResult) document() string {
	return `{
		city {
			id
		}
	}`
}

type deleteResult struct {
	DeleteWeather struct {
		Msg     string
		NumUids int
	} `json:"deleteWeather"`
}

func (deleteResult) document() string {
	return `{
		msg,
		numUids,
	}`
}
