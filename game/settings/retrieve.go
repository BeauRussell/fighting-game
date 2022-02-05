package settings

import (
	"embed"
	"log"

	"gopkg.in/yaml.v2"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Data struct {
	Window struct {
		Height float64 `yaml:"height"`
		Width  float64 `yaml:"width"`
	}
	Movement struct {
		Right   float64 `yaml:"right"`
		Left    float64 `yaml:"left"`
		Up      float64 `yaml:"up"`
		Gravity float64 `yaml:"gravity"`
	}
	Stage struct {
		Ground float64 `yaml:"ground"`
	}
	Player struct {
		Player1 struct {
			Health float64 `yaml:"health"`
			Speed  float64 `yaml:"speed"`
			Damage float64 `yaml:"damage"`
		}
		Player2 struct {
			Health float64 `yaml:"health"`
			Speed  float64 `yaml:"speed"`
			Damage float64 `yaml:"damage"`
		}
	}
}

func RetrieveSettings(content embed.FS) Data {
	d, err := content.ReadFile("settings.yml")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	data := Data{}

	err = yaml.Unmarshal([]byte(d), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return data
}
