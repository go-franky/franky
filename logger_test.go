package franky_test

import (
	"log"
	"os"
	"testing"

	"github.com/go-franky/franky"
	"github.com/sirupsen/logrus"
)

func TestLoggerCompatibility(t *testing.T) {
	var _ franky.Logger = franky.DefaultLogger
	var _ franky.Logger = franky.NoopLogger

	// I use logrus enough, so make sure it works
	var _ franky.Logger = logrus.New()
	logger := log.New(os.Stdout, "[Prefix]",
		log.Ldate|
			log.Ldate| // the date in the local time zone: 2009/01/23
			log.Ltime| // the time in the local time zone: 01:23:23
			log.Llongfile|
			log.Lshortfile| // final file name element and line number: d.go:23. overrides Llongfile
			log.LUTC, // if Ldate or Ltime is set, use UTC rather than the local time zone
	)
	logger.Printf("it works!")
}
