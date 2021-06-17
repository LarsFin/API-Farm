package apifarm

type VideoGame struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Developers   []string   `json:"developers"`
	Publishers   []string   `json:"publishers"`
	Directors    []string   `json:"directors"`
	Producers    []string   `json:"producers"`
	Designers    []string   `json:"designers"`
	Programmers  []string   `json:"programmers"`
	Artists      []string   `json:"artists"`
	Composers    []string   `json:"composers"`
	Platforms    []string   `json:"platforms"`
	DateReleased CustomTime `json:"date_released"`
}
