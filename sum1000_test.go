package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSum1000(t *testing.T) {
	want := 0
	for i := 0; i <= 1000; i++ {
		want = want + i
	}
	require.Equal(t, want, sum1000(10))
}