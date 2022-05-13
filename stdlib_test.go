package stdlib

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		// For some reason this value is lower than what -cover reports
		if c < 0.91 {
			log.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}
