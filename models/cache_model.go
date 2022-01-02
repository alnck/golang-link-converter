package models

import "time"

type CacheModel struct {
	Date          time.Time //`json:"date"`
	ConvertedLink string    //`json:"convertedlink"`
	EventType     string    //`json:"eventtype"`
	ResourceType  string    //`json:"resourcetype"`
}
