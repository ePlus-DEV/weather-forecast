package collector

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/ePlus-DEV/weather-forecast/model"
	"github.com/ePlus-DEV/weather-forecast/pkg/errs"
	"html/template"
	"log/slog"
	"os"
	"time"
)

type Collector struct {
	weatherService WeatherService
}

func NewCollector(weatherService WeatherService) *Collector {
	return &Collector{weatherService}
}

func (c *Collector) Collect(ctx context.Context, city string, days int, readmeTemplateFile string, outFilePath string) error {
	slog.Info(fmt.Sprintf("Collecting weather for %s for %d days - Template file: %s", city, days, readmeTemplateFile))
	weathers, err := c.weatherService.Forecast(ctx, city, days)
	if err != nil {
		return errs.Joinf(err, "[weatherService.Forecast]")
	}
	readmeTemplate, err := os.ReadFile(readmeTemplateFile)
	if err != nil {
		return errs.Joinf(err, "[os.ReadFile] "+readmeTemplateFile)
	}
	readme, err := generateReadme(weathers, string(readmeTemplate), templates...)
	if err != nil {
		return errs.Joinf(err, "[generateReadme]")
	}

	return os.WriteFile(outFilePath, []byte(*readme), 0644)
}

func generateReadme(weathers []model.Weather, readmeTemplate string, templates ...string) (*string, error) {
	if len(weathers) == 0 {
		return nil, errors.New("weathers must be not empty")
	}
	tmpl, err := template.
		New("readme").
		Funcs(template.FuncMap{
			"formatDate": formatDate,
			"formatHour": formatHour,
			"formatTime": formatTime,
			"formatHourOnly": formatHourOnly,
			"currentHour": currentHour,
		}).
		Parse(readmeTemplate)
	if err != nil {
		return nil, err
	}

	for _, t := range templates {
		tmpl, err = tmpl.Parse(t)
		if err != nil {
			return nil, err
		}
	}

	var result bytes.Buffer
	err = tmpl.ExecuteTemplate(&result, "readme", map[string]any{
		"Weathers":     weathers,
		"UpdatedAt":    time.Now(),
		"TodayWeather": weathers[0],
	})
	if err != nil {
		return nil, err
	}
	stringResult := result.String()
	return &stringResult, nil
}
