package types

type Event struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	EventTime   string `json:"event_time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageURL    string `json:"imageURL"`
}
