package shutdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtx(t *testing.T) {
	ctx := Ctx()
	Force()
	assert.NotNil(t, ctx, "nil context")
	assert.NotNil(t, ctx.Err(), "nil err after force")
}
