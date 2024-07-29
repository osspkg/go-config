/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package config_test

import (
	"os"
	"testing"

	"go.osspkg.com/casecheck"
	"go.osspkg.com/config"
	"go.osspkg.com/config/env"
)

func TestUnit_ConfigResolve(t *testing.T) {
	type (
		TestConfigItem struct {
			Home string `yaml:"home"`
			Path string `yaml:"path"`
		}
		TestConfig struct {
			Envs TestConfigItem `yaml:"envs"`
		}
	)

	filename := "/tmp/TestUnit_ConfigResolve.yaml"
	data := `
envs:
    home: "@env(HOME#fail)"
    path: "@env(PATH#fail)"
`
	err := os.WriteFile(filename, []byte(data), 0755)
	casecheck.NoError(t, err)

	res := config.New(env.New())

	err = res.OpenFile(filename)
	casecheck.NoError(t, err)
	err = res.Build()
	casecheck.NoError(t, err)

	var tc TestConfig

	casecheck.NoError(t, res.Decode(&tc))
	casecheck.NotEqual(t, "fail", tc.Envs.Home)
	casecheck.NotEqual(t, "fail", tc.Envs.Path)
}
