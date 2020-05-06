/**
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package realis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentBatchCalculator(t *testing.T) {
	t.Run("singleBatchOverflow", func(t *testing.T) {
		curBatch := calculateCurrentBatch(10, []int32{2})
		assert.Equal(t, 4, curBatch)
	})

	t.Run("noInstancesUpdating", func(t *testing.T) {
		curBatch := calculateCurrentBatch(0, []int32{2})
		assert.Equal(t, 0, curBatch)
	})

	t.Run("evenMatchSingleBatch", func(t *testing.T) {
		curBatch := calculateCurrentBatch(2, []int32{2})
		assert.Equal(t, 0, curBatch)
	})

	t.Run("moreInstancesThanBatches", func(t *testing.T) {
		curBatch := calculateCurrentBatch(5, []int32{1, 2})
		assert.Equal(t, 2, curBatch)
	})

	t.Run("moreInstancesThanBatchesDecreasing", func(t *testing.T) {
		curBatch := calculateCurrentBatch(5, []int32{2, 1})
		assert.Equal(t, 3, curBatch)
	})

	t.Run("unevenFit", func(t *testing.T) {
		curBatch := calculateCurrentBatch(2, []int32{1, 2})
		assert.Equal(t, 1, curBatch)
	})

	t.Run("halfWay", func(t *testing.T) {
		curBatch := calculateCurrentBatch(1, []int32{1, 2})
		assert.Equal(t, 0, curBatch)
	})
}
