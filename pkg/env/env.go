package env

import (
	"log"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver/v3"
)

func PkgManagerVersion(pkgManager string) *semver.Version {
	cmd := exec.Command(pkgManager, "-v")
	out, err := cmd.Output()
	strOut := strings.Trim(string(out), "\n")
	v := semver.MustParse(strOut)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
