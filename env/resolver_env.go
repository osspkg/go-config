/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package env

import "os"

type Resolver struct{}

func New() *Resolver {
	return &Resolver{}
}

func (e *Resolver) Name() string {
	return "env"
}

func (e *Resolver) Resolve(keys ...string) (map[string][]byte, error) {
	out := make(map[string][]byte, len(keys))

	for _, key := range keys {
		if val, ok := os.LookupEnv(key); ok {
			out[key] = []byte(val)
		}
	}

	return out, nil
}
