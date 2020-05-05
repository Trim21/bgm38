package vote

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type Filter struct {
	Permission []int `yaml:"permission" json:"permission"`
}
type Options struct {
	Vote    bool     `yaml:"vote" json:"vote"`
	Multi   bool     `yaml:"multi" json:"multi"`
	Filter  Filter   `yaml:"filter" json:"filter"`
	Options []string `yaml:"options" json:"options"`
}

func parseOption(s string) (options Options, err error) {
	err = yaml.UnmarshalStrict([]byte(strings.ReplaceAll(s, "\u00A0", " ")), &options)
	return
}
