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

package hash64

import (
	"hash/crc32"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash_Index(t *testing.T) {
	hash := New(func(b []byte) uint64 {
		return uint64(crc32.ChecksumIEEE(b))
	})

	index := hash.Sum([]byte("1"))
	assert.EqualValues(t, 0, index)
	index = hash.Sum([]byte("1"))
	assert.EqualValues(t, 0, index)

	index = hash.Sum([]byte("2"))
	assert.EqualValues(t, 1, index)

	index = hash.Sum([]byte("32"))
	assert.EqualValues(t, 2, index)

	index = hash.Sum([]byte("2"))
	assert.EqualValues(t, 1, index)
}
