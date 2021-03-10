package main

import (
	"testing"

	_ "net/http/pprof"

	"github.com/stretchr/testify/assert"
)

func TestFungsiSederhana(t *testing.T) {
	result := FungsiSederhana(4.0, 3.0)
	assert.Equal(t, 9.0, result, "Hasil tidak sama. Error!")
}

func BenchmarkFungsiSederhana(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FungsiSederhana(4.0, 3.0)
	}
}
