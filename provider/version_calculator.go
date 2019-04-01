package provider

import (
	"fmt"
	"strings"
	"time"
)

func SemverCalculator(r *Release) (string, error) {
	mmp, err := ReadSemver(r.CurrentVersion)
	if err != nil {
		return "", fmt.Errorf("cannot calculate next version. version: %s", r.CurrentVersion)
	}

	var mask byte = 0
	for i := range r.Changelog {
		mask = mask | r.Changelog[i].Level
	}

	switch {
	case mask&MAJOR == MAJOR:
		mmp[0] += 1
		mmp[1] = 0
		mmp[2] = 0
	case mask&MINOR == MINOR:
		mmp[1] += 1
		mmp[2] = 0
	case mask&PATCH == PATCH:
		mmp[2] += 1
	}
	version := fmt.Sprintf("%d.%d.%d", mmp[0], mmp[1], mmp[2])
	return strings.ReplaceAll(r.VersionPattern, "SEMVER", version), nil
}

func DateVersionCalculator(r *Release) (string, error) {
	t := time.Now()

	date := t.Format("2006-01-02-150405")

	return strings.ReplaceAll(r.VersionPattern, "DATE", date), nil
}
