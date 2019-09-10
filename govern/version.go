package govern

import (
	"fmt"
	"regexp"
	"strconv"
)

type Version struct {
	Major int
	Minor int
	Build int
	Tag   VersionTag
}

func (v Version) String() string {
	version := fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Build)
	if v.Tag != NoLabel {
		version += "-" + v.Tag.String()
	}
	return version
}

type VersionTag uint8

const (
	NoLabel VersionTag = iota
	Alpha
	Beta
	Gamma
	Demo
	Release
	Stable
)

func (v VersionTag) String() string {
	if s, ok := versionNames[v]; ok {
		return s
	}
	return ""
}

func convertTag(tag string) VersionTag {
	if t, ok := versionTypes[tag]; ok {
		return t
	}
	return NoLabel
}

var versionNames = map[VersionTag]string{
	NoLabel: "",
	Alpha:   "alpha",
	Beta:    "beta",
	Gamma:   "gamma",
	Demo:    "demo",
	Release: "release",
	Stable:  "stable",
}

var versionTypes = map[string]VersionTag{
	"":        NoLabel,
	"alpha":   Alpha,
	"beta":    Beta,
	"gamma":   Gamma,
	"demo":    Demo,
	"release": Release,
	"stable":  Stable,
}

var reg = regexp.MustCompile(`v(?P<v1>\d+)\.(?P<v2>\d+)\.(?P<v3>\d+)(-(?P<tag>\w+))?`)

func ConvertVersion(version string) Version {
	match := reg.FindStringSubmatch(version)
	vn := emptyVersion
	groups := reg.SubexpNames()
	for idx, name := range groups {
		switch name {
		case "tag":
			vn.Tag = convertTag(match[idx])
		default:
			if it, err := strconv.Atoi(match[idx]); err == nil {
				switch name {
				case "v1":
					vn.Major = it
				case "v2":
					vn.Minor = it
				case "v3":
					vn.Build = it
				}
			}
		}
	}
	return vn
}

var emptyVersion = Version{}
