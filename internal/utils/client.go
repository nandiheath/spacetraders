package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nandiheath/spacetraders/internal/api"
)

// NewAPIClient returns the APIClient with default configuration
func NewAPIClient() *api.ClientWithResponses {
	c, _ := api.NewClientWithResponses("https://api.spacetraders.io/v2", func(client *api.Client) error {
		client.RequestEditors = append(client.RequestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN")))
			return nil
		})
		return nil
	})
	return c
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
