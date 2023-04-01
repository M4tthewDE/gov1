package parser

import (
	"bytes"
	"encoding/binary"
	"os"
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

	obu := obus[len(obus)-1]

	buf := new(bytes.Buffer)
	for _, yArr := range obu.TileGroup.OutY {
		for _, y := range yArr {
			binary.Write(buf, binary.BigEndian, y)
		}
	}

	for _, uArr := range obu.TileGroup.OutU {
		for _, u := range uArr {
			binary.Write(buf, binary.BigEndian, u)
		}
	}

	for _, vArr := range obu.TileGroup.OutV {
		for _, v := range vArr {
			binary.Write(buf, binary.BigEndian, v)
		}
	}

	f, err := os.Create("../../out.yuv")
	check(err)

	defer f.Close()

	n, err := f.Write(buf.Bytes())
	check(err)

	assert.NotEqual(t, 0, n)
}
