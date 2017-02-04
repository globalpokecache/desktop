package models

type ScanConfig struct {
	Latitude, Longitude, Radius float64
	Accounts                    []string
	Cooldown                    int
}
