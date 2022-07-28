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

package hash32

import (
	"sync"
	"sync/atomic"
)

type Hash struct {
	hashFunc func(b []byte) uint32
	m        *sync.Map
	n        uint32
}

func New(hashFunc func(b []byte) uint32) *Hash {
	return &Hash{
		hashFunc: hashFunc,
		m:        &sync.Map{},
	}
}

func (h *Hash) Sum(b []byte) (index uint32) {
	hashCode := h.hashFunc(b)
	indexVal, ok := h.m.Load(hashCode)
	if !ok {
		index = atomic.AddUint32(&h.n, 1) - 1
		h.m.Store(hashCode, index)
		return index
	}

	return indexVal.(uint32)
}
