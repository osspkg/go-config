/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package env

import (
	"os"
)

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

func (e *Resolver) Name() string {
	return "env"
}

func (e *Resolver) Value(name string) (string, bool) {
	return os.LookupEnv(name)
}
