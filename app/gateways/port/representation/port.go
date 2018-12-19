package representation

type PortEntityResponse struct {
}

type PortListResponse struct {
	Items []*PortEntityResponse `json:"items"`
	Found int                   `json:"found"`
}
