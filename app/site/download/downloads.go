// Content managed by Project Forge, see [projectforge.md] for details.
package download

import (
	"fmt"

	"github.com/kyleu/loadtoad/app/util"
)

var (
	arms = []string{ArchARMV5, ArchARMV6, ArchARMV7}
	mips = []string{ArchMIPS64Hard, ArchMIPS64Soft, ArchMIPS64LEHard, ArchMIPS64LESoft, ArchMIPSHard, ArchMIPSSoft, ArchMIPSLEHard, ArchMIPSLESoft}
)

func GetLinks(version string) Links {
	if availableLinks == nil {
		availableLinks = calcDownloadLinks(version)
	}
	return availableLinks
}

func calcDownloadLinks(version string) Links {
	ret := Links{}
	add := func(url string, mode string, os string, arch string) {
		ret = append(ret, &Link{URL: url, Mode: mode, OS: os, Arch: arch})
	}
	addDefault := func(mode string, os string, arch string) {
		var u string
		switch mode {
		case ModeServer:
			msg := "%s_%s_%s_%s.zip"
			if os == OSMac {
				msg = "%s_%s_%s_%s_notarized.zip"
			}
			u = fmt.Sprintf(msg, util.AppKey, version, os, arch)
		case ModeMobile:
			u = fmt.Sprintf("%s_%s_%s_%s.zip", util.AppKey, version, os, arch)
		case ModeDesktop:
			u = fmt.Sprintf("%s_%s_%s_%s_desktop.zip", util.AppKey, version, os, arch)
		}
		add(u, mode, os, arch)
	}

	addDefault(ModeServer, OSMac, ArchAMD64)
	addDefault(ModeServer, OSMac, ArchARM64)
	addDefault(ModeServer, OSMac, ArchUniversal)
	addDefault(ModeServer, OSWindows, ArchAMD64)
	addDefault(ModeServer, OSWindows, Arch386)
	addDefault(ModeServer, OSLinux, ArchAMD64)
	addDefault(ModeServer, OSLinux, Arch386)

	return ret
}
