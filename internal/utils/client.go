package utils

import (
	"fmt"

	"strings"
	"time"
)

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
