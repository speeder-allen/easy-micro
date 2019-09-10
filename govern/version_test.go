package govern_test

import (
	"github.com/speeder-allen/easy-micro/govern"
	"gotest.tools/assert"
	"testing"
)

func TestConvertVersion(t *testing.T) {
	v1 := "v1.3.45-demo"
	ver1 := govern.ConvertVersion(v1)
	assert.Equal(t, ver1.Major, 1)
	assert.Equal(t, ver1.Minor, 3)
	assert.Equal(t, ver1.Build, 45)
	assert.Equal(t, ver1.Tag, govern.Demo)
	v2 := "v2.0.13-err"
	ver2 := govern.ConvertVersion(v2)
	assert.Equal(t, ver2.Major, 2)
	assert.Equal(t, ver2.Minor, 0)
	assert.Equal(t, ver2.Build, 13)
	assert.Equal(t, ver2.Tag, govern.NoLabel)
}

func TestVersionTag_String(t *testing.T) {
	assert.Equal(t, govern.Demo.String(), "demo")
	assert.Equal(t, govern.VersionTag(11).String(), govern.NoLabel.String())
}

func TestVersion_String(t *testing.T) {
	v := govern.Version{5, 35, 1451, govern.Release}
	assert.Equal(t, v.String(), "v5.35.1451-release")
}

func TestVersionCompare(t *testing.T) {
	v1 := govern.Version{1, 3, 41, govern.NoLabel}
	v2 := govern.Version{1, 3, 40, govern.NoLabel}
	v3 := govern.Version{2, 1, 34, govern.NoLabel}
	v4 := govern.Version{2, 5, 1, govern.NoLabel}
	v5 := govern.Version{2, 5, 1, govern.Release}
	v6 := govern.Version{2, 5, 1, govern.Gamma}
	v7 := govern.Version{2, 5, 1, govern.Gamma}
	assert.Equal(t, v1.GreaterThan(v2), true)
	assert.Equal(t, v3.LessThan(v4), true)
	assert.Equal(t, v6.LessOrEqualThan(v5), true)
	assert.Equal(t, v7.GreaterOrEqualThan(v6), true)
}
