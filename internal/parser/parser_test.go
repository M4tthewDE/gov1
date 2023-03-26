package parser

import (
	"os"
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestParseEndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	fileName := "../../testdata/argon_coveragetool_av1_base_and_extended_profiles_v2.1/profile0_core/streams/test1228.obu"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	b := bitstream.NewBitStream(data)
	p := NewParser(&b)
	obus := p.bitstream()
	logger.Logger.Info("DONE ", zap.Int("obu count", len(obus)))

	obu := obus[2]

	assert.Equal(t, true, obu.SequenceHeader.ColorConfig.SubsamplingX)
	assert.Equal(t, true, obu.SequenceHeader.ColorConfig.SubsamplingY)
	assert.Equal(t, false, obu.SequenceHeader.ColorConfig.MonoChrome)
}
