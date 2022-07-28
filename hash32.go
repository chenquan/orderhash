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
)

func Hash32(hashFunc func(b []byte) uint32) func(b []byte) uint32 {
	n := uint32(0)
	m := make(map[uint32]uint32, 16)
	rw := &sync.RWMutex{}
	return func(b []byte) uint32 {
		hashCode := hashFunc(b)
		rw.RLock()
		index, ok := m[hashCode]
		rw.RUnlock()

		if !ok {
			rw.Lock()
			m[hashCode], index = n, n
			n++
			rw.Unlock()
			return index
		}

		return index
	}
}