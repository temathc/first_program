package main

import (
	"fmt"
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/require"
)

func TestNumberOne(t *testing.T) {
	const str, minor, want = "the string must be uppercase except", "", "The String Must Be Uppercase Except"
	got := titlecase.TitleCase(str, minor)
	require.Equal(t, want, got)
	fmt.Println(got)
}

func TestSum(t *testing.T) {
	require.Equal(t, 7, Ssum(4, 3))
}

func TestLen(t *testing.T) {
	require.Len(t, str, 8)
}
