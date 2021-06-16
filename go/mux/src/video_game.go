package apifarm

import "time"

type VideoGame struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Developers   []string  `json:"developers"`
	Publishers   []string  `json:"publishers"`
	Directors    []string  `json:"directors"`
	Producers    []string  `json:"producers"`
	Designers    []string  `json:"designers"`
	Programmers  []string  `json:"programmers"`
	Artists      []string  `json:"artists"`
	Composers    []string  `json:"composers"`
	Platforms    []string  `json:"platforms"`
	DateReleased time.Time `json:"date_released"`
}
