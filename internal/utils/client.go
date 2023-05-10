package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nandiheath/spacetraders/internal/api"
)

// NewAPIClient returns the APIClient with default configuration
func NewAPIClient() *api.APIClient {
	cfg := api.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN")))
	return api.NewAPIClient(cfg)
}

func FormatExpiration(t time.Time) string {
	diff := t.Sub(time.Now())
	if diff.Hours() > 24 {
		return fmt.Sprintf("%dd%dh", (int)(diff.Hours())/24, (int)(diff.Hours())%24)
	} else if diff.Minutes() > 60 {
		return fmt.Sprintf("%dh%dm", (int)(diff.Hours()), (int)(diff.Minutes())%60)
	}
	return fmt.Sprintf("%d", (int)(diff.Minutes()))
}

func PrintArray(array []string) {
	fmt.Printf("%s\n", strings.Join(array, "\t"))
}
