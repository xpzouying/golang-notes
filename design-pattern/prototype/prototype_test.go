package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {

	t.Run("clone circle", func(t *testing.T) {
		srcCircle := NewCircle(13)

		clone := srcCircle.Clone()
		cloneCircle, ok := clone.(*Circle)

		assert.True(t, ok)
		assert.Equal(t, srcCircle.Area(), cloneCircle.Area())
	})

	t.Run("clone rectangle", func(t *testing.T) {
		srcRectangle := NewRectangle(13, 31)

		clone := srcRectangle.Clone()

		cloneRectangle, ok := clone.(*Rectangle)

		assert.True(t, ok)
		assert.Equal(t, srcRectangle.Area(), cloneRectangle.Area())
	})
}
