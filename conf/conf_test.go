package conf

import (
	"encoding/json"
	"testing"
)

func TestLoadFromEnv(t *testing.T) {
	LoadFromEnv()

	c := C()
	bs, _ := json.MarshalIndent(c, "", "  ")
	t.Logf("%s", bs)
}

func TestLoadFromToml(t *testing.T) {
	LoadFromToml("../etc/conf.toml")

	c := C()
	bs, _ := json.MarshalIndent(c, "", "  ")
	t.Logf("%s", bs)
}
