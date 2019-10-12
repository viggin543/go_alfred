package alfred

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterItems(t *testing.T) {
	items := []*Item{
		{"da", "da", "da", "da"},
		{"op", "da", "da", "da"},
	}

	filterItems := FilterItems(items, "da")
	assert.Equal(t, 1, len(filterItems))
	assert.Equal(t,
		*filterItems[0],
		Item{"da", "da", "da", "da"})
}
