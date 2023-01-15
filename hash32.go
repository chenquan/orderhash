//   Copyright 2022 chenquan
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package orderhash

import (
	"sync"
	"sync/atomic"
)

type Hash32 struct {
	n        uint32
	mu       sync.Mutex
	value    atomic.Value
	HashFunc func(b []byte) uint32
}

func (h *Hash32) Hash(b []byte) uint32 {
	hashCode := h.HashFunc(b)

	m := h.atomicLoad()
	index, ok := m[hashCode]
	if ok {
		return index
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	m = h.atomicLoad()
	index, ok = m[hashCode]
	if ok {
		return index
	}

	mm := make(map[uint32]uint32, len(m))
	for k, v := range m {
		mm[k] = v
	}

	mm[hashCode], index = h.n, h.n
	h.n++
	h.value.Store(mm)

	return index

}

func (h *Hash32) atomicLoad() map[uint32]uint32 {
	m := h.value.Load()
	if m != nil {
		return m.(map[uint32]uint32)
	}

	return map[uint32]uint32{}
}
