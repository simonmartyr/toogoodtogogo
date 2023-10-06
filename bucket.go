package toogoodtogo

type BucketFilter struct {
	FillerType string `json:"filler_type"`
}
type MobileBucket struct {
	MobileBucket Bucket `json:"mobile_bucket"`
}
type Bucket struct {
	FillerType  string `json:"filler_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Items       []Item `json:"items"`
	BucketType  string `json:"bucket_type"`
	DisplayType string `json:"display_type"`
}
