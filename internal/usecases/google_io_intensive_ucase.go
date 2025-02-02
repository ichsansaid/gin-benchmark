package ucases

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// GoogleIOIntensiveUcase implements GoogleIOIntensiveUcaseContract
type GoogleIOIntensiveUcase struct{}

// NewGoogleIOIntensiveUcase creates a new instance of GoogleIOIntensiveUcase
func NewGoogleIOIntensiveUcase() *GoogleIOIntensiveUcase {
	return &GoogleIOIntensiveUcase{}
}

// DoTask executes an IO-intensive task (Fetching data from an external API)
func (uc *GoogleIOIntensiveUcase) DoTask(ctx context.Context) {

	// Simulating an IO operation: Fetching a web page
	url := "https://www.google.com"
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, resp.Body) // Read response but discard output
}
