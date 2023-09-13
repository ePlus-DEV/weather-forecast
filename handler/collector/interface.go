package collector

import (
	"context"
	"github.com/ePlus-DEV/weather-forecast/model"
)

type WeatherService interface {
	Forecast(ctx context.Context, city string, days int) ([]model.Weather, error)
}
