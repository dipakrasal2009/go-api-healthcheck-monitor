package colorprint

import (
	"fmt"

	"github.com/dipakrasal2009/go-api-healthcheck-monitor/models"
	"github.com/fatih/color"
)

// PrintStatus prints the service health status with color.
func PrintStatus(s models.Service) {
	status := "Healthy"

	if !s.Healthy {
		fmt.Printf("%s: %s\n", s.Name, color.RedString(status))
	} else {
		fmt.Printf("%s: %s\n", s.Name, color.GreenString(status))
	}
}
