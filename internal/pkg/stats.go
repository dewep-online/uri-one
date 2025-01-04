/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package pkg

import "sync"

type Stats struct {
	data map[uint64]uint64
	mux  sync.Mutex
}

func NewStats() *Stats {
	return &Stats{
		data: make(map[uint64]uint64, 1000),
	}
}

func (v *Stats) Add(id uint64) {
	v.mux.Lock()
	defer v.mux.Unlock()

	v.data[id]++
}

func (v *Stats) Get() map[uint64]uint64 {
	v.mux.Lock()
	defer v.mux.Unlock()

	result := make(map[uint64]uint64, len(v.data))
	for id, count := range v.data {
		result[id] = count
		delete(v.data, id)
	}
	return result
}
