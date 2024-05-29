package cmd

import (
	"os"
	"strings"
)

func Parse() {
	switch strings.ToLower(os.Getenv("TEST_MODE")) {
	case "destiny":
		ParseDestiny()
	case "load":
		ParseLoad()
	}
}
