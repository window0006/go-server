package dal

type Family struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Patriarch int    `json:"patriarch"`
	Members   string `json:"members"`
}
