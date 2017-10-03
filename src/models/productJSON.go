package models

type (
	// ProductJSON represents the structure of our resource
	ProductJSON struct {
		ProductName   string `json:"productName"`
		Color         string `json:"color"`
		Size          string `json:"size"`
		EventType     string `json:"eventType"`
		TraditionType string `json:"traditionType"`
	}
)
