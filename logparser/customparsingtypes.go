package logparser

import (
	"strconv"
	"regexp"
)

type Bool struct {
	Value bool
}

type Regexp struct {
	Value *regexp.Regexp
}

func (b *Bool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var boolStr string
	err := unmarshal(&boolStr)
	if err != nil {
		return err
	}
	parse, err := strconv.ParseBool(boolStr)
	if err != nil {
		return err
	}

	b.Value = parse
	return nil
}

func (reg *Regexp) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var regexpStr string
	err := unmarshal(&regexpStr)
	if err != nil {
		return err
	}
	parse, err := regexp.Compile(regexpStr)
	if err != nil {
		return err
	}

	reg.Value = parse
	return nil
}