package controller

type spoty struct {
	Damso []struct {
		ID string `json:"id"`
		Images []string `json:"images"`
		Name string `jsnon:"name"`
		Numb int `json:"total_tracks"`
		Date int `json:"release_date"`
	} 
	Laylow [] struct {
		ID string `json:"id"`
		Name string `jsnon:"name"`
		Images []string `json:"images"`
		Anmae []string `json:"artists"`
	}
}
