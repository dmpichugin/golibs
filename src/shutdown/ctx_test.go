package shutdown

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCtx(t *testing.T) {
	ctx := Ctx()
	Force()
	require.NotNil(t, ctx, "nil context")
	require.NotNil(t, ctx.Err(), "nil err after force")
}
