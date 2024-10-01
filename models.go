package main

type NpsbPlugin struct{}

type GenericRequest struct {
	SpaceGUID string `json:"spaceguid"`
}

type SourceListResponse struct {
	Sources []SourceResponse `json:"source_responses"`
}

type SourceResponse struct {
	Source      string `json:"source"`
	Scope       string `json:"scope"`
	Org         string `json:"org"`
	Space       string `json:"space"`
	Description string `json:"description"`
}
