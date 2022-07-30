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
	"hash/crc32"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash64(t *testing.T) {
	f := Hash64(func(b []byte) uint64 {
		return uint64(crc32.ChecksumIEEE(b))
	})

	N := 10000
	group := sync.WaitGroup{}
	m := sync.Map{}
	for i := 0; i < N; i++ {
		i := i
		for j := 0; j < 100; j++ {
			group.Add(1)
			go func() {
				defer group.Done()
				code := f([]byte(strconv.Itoa(i)))
				m.Store(code, i)
			}()
		}
	}
	group.Wait()

	for i := 0; i < N; i++ {
		i := i
		group.Add(1)
		go func() {
			defer group.Done()
			code := f([]byte(strconv.Itoa(i)))
			v, ok := m.Load(code)
			assert.True(t, ok)
			assert.EqualValues(t, v, i)
		}()
	}
	group.Wait()
}
