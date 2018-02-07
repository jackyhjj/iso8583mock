// Package helpers utilities code
// @author Valentino <daud.darianus@kudo.co.id>
package helpers

import (
	"os"
	"regexp"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// AlphaNumClean the alpha numeric
func AlphaNumClean(p string) string {
	return regexp.MustCompile("[^0-9a-zA-Z]+").ReplaceAllString(p, "")
}

// NumericClean is only number allow
func NumericClean(p string) string {
	return regexp.MustCompile("[^0-9]+").ReplaceAllString(p, "")
}

// ParseInteger convert the string to  integer
func ParseInteger(p string) (int64, error) {

	p = regexp.MustCompile("[^0-9]+").ReplaceAllString(p, "")

	r, err := strconv.ParseInt(p, 10, 64)

	if err != nil {
		log.Error(err)
	}
	return r, err
}

// PathExist check the path directory if exist
func PathExist(p string) bool {
	if stat, err := os.Stat(p); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// TimeTrack - track processing time
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
