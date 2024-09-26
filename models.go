package main

type SchedulerPlugin struct{}

type GenericRequest struct {
	SpaceGUID string `json:"spaceguid"`
}

type SourceListResponse struct {
	Sources []SourceResponse `json:"source_responses"`
}

type SourceResponse struct {
	Source string `json:"source"`
	Scope  string `json:"scope"`
	Org    string `json:"org"`
	Space  string `json:"space"`
}
