package energymonitor

import (
	"encoding/json"
	"net/http"
)

type Config struct {
	URL string
}

type Client struct {
	URL    string
	Sensor PowerSensor
}

type PowerSensor struct {
	SiteName      string
	ModuleName    string
	Irms          float32
	Vrms          float32
	RealPower     float32
	ApparentPower float32
	PowerFactor   float32
	ExecTime      int
	Time          int
}

func NewClient(config Config) (*Client, error) {
	return &Client{
		URL:    config.URL,
		Sensor: PowerSensor{},
	}, nil
}

func (c *Client) Read() error {
	resp, err := http.Get(c.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&c.Sensor)
}

func (c *Client) PowerSensor() PowerSensor {
	return c.Sensor
}
