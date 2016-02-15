package results

// ProgramDetails is simplified structure for responses.Program, cotaining
// parsed & relevant fields only
type ProgramDetails struct {
	Title       string
	MediaID     string
	HLSURL      string
	Slug        string
	SubtitleURL string
}
