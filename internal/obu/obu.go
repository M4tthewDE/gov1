package obu

import (
	"encoding/json"
	"fmt"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
)

type ObuType int

const OBU_SEQUENCE_HEADER = 1
const OBU_TEMPORAL_DELIMITER = 2
const OBU_FRAME_HEADER = 3
const OBU_TILE_GROUP = 4
const OBU_METADATA = 5
const OBU_FRAME = 6
const OBU_REDUNDANT_FRAME_HEADER = 7
const OBU_TILE_LIST = 8
const OBU_PADDING = 15

type Obu struct {
	Size int
}

// open_bitstream_unit(sz)
func NewObu(sz int, b *bitstream.BitStream) (Obu, State) {
	obu := Obu{}
	state := State{}

	state.header = header.NewHeader(b)

	if state.header.HasSizeField {
		obu.Size = b.Leb128()
	} else {
		extensionFlagInt := 0
		if state.header.ExtensionFlag {
			extensionFlagInt = 1
		}
		obu.Size = sz - 1 - extensionFlagInt
	}

	startPosition := b.Position

	if state.header.Type != OBU_SEQUENCE_HEADER &&
		state.header.Type != OBU_TEMPORAL_DELIMITER &&
		state.operatingPointIdc != 0 &&
		state.header.ExtensionFlag {
		inTemporalLayer := ((state.operatingPointIdc >> state.header.ExtensionHeader.TemporalID) & 1) != 0
		inSpatialLayer := ((state.operatingPointIdc >> (state.header.ExtensionHeader.SpatialID + 8)) & 1) != 0

		if !inTemporalLayer || !inSpatialLayer {
			//drop_obu()
			b.Position = b.Position + obu.Size*8
			return obu, state
		}
	}

	x, _ := json.MarshalIndent(obu, "", "	")
	fmt.Printf("%s\n", string(x))

	switch state.header.Type {
	case OBU_SEQUENCE_HEADER:
		sequenceheader, result := sequenceheader.NewSequenceHeader(b)
		state.sequenceHeader = sequenceheader
		state.operatingPointIdc = result.OperatingPointIdc

		x, _ := json.MarshalIndent(state.sequenceHeader, "", "	")
		fmt.Printf("%s\n", string(x))
	case OBU_TEMPORAL_DELIMITER:
		state.seenFrameHeader = false
	case OBU_FRAME:
		newFrame(obu.Size, b)
	case OBU_METADATA:

	default:
		fmt.Printf("not implemented type %d\n", state.header.Type)
		panic("")
	}

	payloadBits := b.Position - startPosition

	fmt.Println("----------------------------------------")
	fmt.Printf("p.position: %d\n", b.Position)
	fmt.Printf("startPosition: %d\n", startPosition)
	fmt.Printf("payloadBits: %d\n", payloadBits)
	fmt.Printf("obu.Size*8 - payloadBits: %d\n", obu.Size*8-payloadBits)
	fmt.Println("----------------------------------------")

	if obu.Size > 0 &&
		state.header.Type != OBU_TILE_GROUP &&
		state.header.Type != OBU_TILE_LIST &&
		state.header.Type != OBU_FRAME {
		b.TrailingBits(obu.Size*8 - payloadBits)
	}

	return obu, state
}

// frame_obu( sz )
func newFrame(sz int, b *bitstream.BitStream) {
	startBitPos := b.Position

	ParseFrameHeader()
	b.ByteAlignment()

	endBitPos := b.Position

	headerBytes := (endBitPos - startBitPos) / 8
	sz -= headerBytes
	_ = NewTileGroup(b, sz)
}

// frame_header_obu()
func ParseFrameHeader() {
	if p.seenFrameHeader {
		p.FrameHeaderCopy()
	} else {
		p.seenFrameHeader = true
		uncompressedHeader := NewUncompressedHeader(p)

		if uncompressedHeader.ShowExistingFrame {
			p.DecodeFrameWrapup()
			p.seenFrameHeader = false
		} else {
			p.TileNum = 0
			p.seenFrameHeader = true
		}

	}
}

// frame_header_copy()
func FrameHeaderCopy() {
	panic("not implemented")
}
