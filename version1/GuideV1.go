package version1

import "time"

type GuideV1 struct {
	// Identification
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	App    string `json:"app"`
	MinVer int    `json:"min_ver"`
	MaxVer int    `json:"max_ver"`

	// Automatically managed fields
	CreatedTime time.Time `json:"created_time"`

	// Content
	Pages []GuidePageV1 `json:"pages"`

	// Search
	Tags    []string `json:"tags"`
	AllTags []string `json:"all_tags"`

	// Status
	Status string `json:"status"`

	// Custom fields
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
}
