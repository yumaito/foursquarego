package foursquarego

type apiResponse struct {
	Meta          Meta               `json:"meta"`
	Notifications Omit               `json:"-"`
	Response      foursquareResponse `json:"response"`
}

type Meta struct {
	Code        int    `json:"code"`
	ErrorType   string `json:"errorType"`
	ErrorDetail string `json:"errorDetail"`
}

type foursquareResponse struct {
	Venue      Venue          `json:"venue,omitempty"`
	Categories []Category     `json:"categories,omitempty"`
	Photos     PhotosResponse `json:"photos,omitempty"`
	Events     EventsResponse `json:"events,omitempty"`
	HereNow    HereNow        `json:"hereNow,omitempty"`
	Hours      Hours          `json:"hours,omitempty"`
	Popular    Hours          `json:"popular,omitempty"`
}

type PhotosResponse struct {
	Count        int     `json:"count"`
	Items        []Photo `json:"items"`
	DupesRemoved int     `json:"dupesRemoved"`
}

type EventsResponse struct {
	Count   int     `json:"count"`
	Summary string  `json:"summary"`
	Items   []Event `json:"items"`
}

type HoursResponse struct {
	Hours   Hours `json:"hours,omitempty"`
	Popular Hours `json:"popular,omitempty"`
}
