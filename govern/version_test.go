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
