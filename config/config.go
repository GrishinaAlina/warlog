// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

// Package config provides app configuration reading and validating.
package config

// Config described configuration file with reading.
type Config interface {
	FromFile(path string)
}

// FromFile reads and validates config to `s` struct from `path` file.
func FromFile(path string, s interface{}) {

}

// V1_0 is the first implementation of config file.
type V1_0 struct {
	Config
}

// FromFile reads and validates config to struct instance from `path` file.
func (config *V1_0) FromFile(path string) {
	FromFile(path, config)
}
