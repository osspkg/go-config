/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"go.osspkg.com/ioutils/codec"
)

type (
	Resolver interface {
		Name() string
		Value(name string) (string, bool)
	}

	Config struct {
		data *codec.BlobEncoder
		list []Resolver
	}
)

func New(list ...Resolver) *Config {
	return &Config{
		data: nil,
		list: list,
	}
}

func (v *Config) OpenFile(filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	v.data = &codec.BlobEncoder{
		Blob: b,
		Ext:  filepath.Ext(filename),
	}
	return nil
}

func (v *Config) Decode(configs ...interface{}) error {
	return v.data.Decode(configs...)
}

var rexName = regexp.MustCompile(`(?m)^[a-z][a-z0-9]+$`)

func (v *Config) Build() error {
	for _, r := range v.list {
		if !rexName.MatchString(r.Name()) {
			return fmt.Errorf("resolver '%s' has invalid name, must like regexp [a-z][a-z0-9]+", r.Name())
		}
		rex := regexp.MustCompile(fmt.Sprintf(`(?mUsi)@%s\((.+)#(.*)\)`, r.Name()))
		submatchs := rex.FindAllSubmatch(v.data.Blob, -1)

		for _, submatch := range submatchs {
			pattern, key, defval := submatch[0], submatch[1], submatch[2]

			if val, ok := r.Value(string(key)); ok && len(val) > 0 {
				defval = []byte(val)
			}

			v.data.Blob = bytes.ReplaceAll(v.data.Blob, pattern, defval)
		}
	}
	return nil
}
