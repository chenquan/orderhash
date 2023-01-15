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

func Hash32(hashFunc func(b []byte) uint32) func(b []byte) uint32 {
	n := uint32(0)
	mu := &sync.Mutex{}
	value := &atomic.Value{}
	value.Store(make(map[uint32]uint32, 16))
	return func(b []byte) uint32 {
		hashCode := hashFunc(b)

		m := value.Load().(map[uint32]uint32)
		index, ok := m[hashCode]
		if ok {
			return index
		}

		mu.Lock()
		defer mu.Unlock()

		m = value.Load().(map[uint32]uint32)
		index, ok = m[hashCode]
		if ok {
			return index
		}

		mm := make(map[uint32]uint32, len(m))
		for k, v := range m {
			mm[k] = v
		}

		mm[hashCode], index = n, n
		n++
		value.Store(mm)

		return index
	}
}
