package rsc

import (
	_ "embed"
	"strings"
)

var (
	//go:embed about.txt
	aboutTxt string
	//go:embed support.txt
	SupportTxt string

	ChangelogMd string
)

func AboutTxt(version, revision string) string {
	txt := strings.Replace(aboutTxt, "%VERSION%", version, 1)
	txt = strings.Replace(txt, "%REVISION%", revision, 1)
	return txt
}
