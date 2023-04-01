package parser

import (
	"bufio"
	"os"
	"testing"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/obu"
	"github.com/stretchr/testify/assert"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getStreamNames() []string {
	streamList, err := os.Open("../../testdata/argon_coveragetool_av1_base_and_extended_profiles_v2.1/profile0_core/level5.x_list.txt")
	check(err)

	fileScanner := bufio.NewScanner(streamList)
	fileScanner.Split(bufio.ScanLines)

	streamNames := []string{}
	for fileScanner.Scan() {
		streamNames = append(streamNames, fileScanner.Text())
	}

	return streamNames
}

func saveObu(t *testing.T, obu obu.Obu) {
	bytes := []byte{}
	for _, yArr := range obu.TileGroup.OutY {
		for _, y := range yArr {
			bytes = append(bytes, byte(y))
		}
	}

	for _, uArr := range obu.TileGroup.OutU {
		for _, u := range uArr {
			bytes = append(bytes, byte(u))
		}
	}

	for _, vArr := range obu.TileGroup.OutV {
		for _, v := range vArr {
			bytes = append(bytes, byte(v))
		}
	}

	f, err := os.Create("../../out.yuv")
	check(err)
	defer f.Close()
	n, err := f.Write(bytes)
	check(err)

	assert.NotEqual(t, 0, n)
}

func cleanUp() {
	err := os.Remove("../../out.yuv")
	check(err)
}

func TestParseEndToEndProfile0Core(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	streamNames := getStreamNames()

	for _, name := range streamNames {
		fileName := "../../testdata/argon_coveragetool_av1_base_and_extended_profiles_v2.1/profile0_core/streams/" + name
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		b := bitstream.NewBitStream(data)
		p := NewParser(&b)
		p.bitstream()
	}

}
