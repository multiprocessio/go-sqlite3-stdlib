package stdlib

import (
	"log"
	"os"
	"testing"
)

// SOURCE: https://stackoverflow.com/a/50123125/1507139
func TestMain(m *testing.M) {
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		// For some reason this value is lower than what -cover reports
		if c < 0.95 {
			log.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}
