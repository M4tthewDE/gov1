package parser

import (
	"os"
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
)

func TestParseEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	fileName := "testdata/argon_coveragetool_av1_base_and_extended_profiles_v2.1/profile0_core/streams/test1228.obu"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	b := bitstream.NewBitStream(data)
	p := NewParser(&b)
	p.bitstream()
}
