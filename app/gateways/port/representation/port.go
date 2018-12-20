package representation

type PortEntityResponse struct {
	PortID      string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Alias       []string  `json:"alias"`
	Unlocs      []string  `json:"unlocs"`
	Country     string    `json:"country"`
	Regions     []string  `json:"regions"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	Coordinates []float32 `json:"coordinates"`
	Timezone    string    `json:"timezone"`
}

type PortListResponse struct {
	Items []*PortEntityResponse `json:"items"`
	Found int                   `json:"found"`
}
