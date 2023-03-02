package obu

import (
	"encoding/json"
	"fmt"

	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
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
	State State
	Size  int
}

func NewObu(sz int, state State, b *bitstream.BitStream) Obu {
	obu := Obu{}
	obu.State = state

	return obu
}

// open_bitstream_unit(sz)
func (o *Obu) build(sz int, b *bitstream.BitStream) {
	o.State.Header = header.NewHeader(b)

	if o.State.Header.HasSizeField {
		o.Size = b.Leb128()
	} else {
		extensionFlagInt := 0
		if o.State.Header.ExtensionFlag {
			extensionFlagInt = 1
		}
		o.Size = sz - 1 - extensionFlagInt
	}

	startPosition := b.Position

	if o.State.Header.Type != OBU_SEQUENCE_HEADER &&
		o.State.Header.Type != OBU_TEMPORAL_DELIMITER &&
		o.State.OperatingPointIdc != 0 &&
		o.State.Header.ExtensionFlag {
		inTemporalLayer := ((o.State.OperatingPointIdc >> o.State.Header.ExtensionHeader.TemporalID) & 1) != 0
		inSpatialLayer := ((o.State.OperatingPointIdc >> (o.State.Header.ExtensionHeader.SpatialID + 8)) & 1) != 0

		if !inTemporalLayer || !inSpatialLayer {
			//drop_obu()
			b.Position = b.Position + o.Size*8
			return
		}
	}

	x, _ := json.MarshalIndent(o, "", "	")
	fmt.Printf("%s\n", string(x))

	switch o.State.Header.Type {
	case OBU_SEQUENCE_HEADER:
		sequenceheader, result := sequenceheader.NewSequenceHeader(b)
		o.State.SequenceHeader = sequenceheader
		o.State.OperatingPointIdc = result.OperatingPointIdc

		x, _ := json.MarshalIndent(o.State.SequenceHeader, "", "	")
		fmt.Printf("%s\n", string(x))
	case OBU_TEMPORAL_DELIMITER:
		o.State.SeenFrameHeader = false
	case OBU_FRAME:
		o.newFrame(o.Size, b)
	case OBU_METADATA:

	default:
		fmt.Printf("not implemented type %d\n", o.State.Header.Type)
		panic("")
	}

	payloadBits := b.Position - startPosition

	fmt.Println("----------------------------------------")
	fmt.Printf("p.position: %d\n", b.Position)
	fmt.Printf("startPosition: %d\n", startPosition)
	fmt.Printf("payloadBits: %d\n", payloadBits)
	fmt.Printf("obu.Size*8 - payloadBits: %d\n", o.Size*8-payloadBits)
	fmt.Println("----------------------------------------")

	if o.Size > 0 &&
		o.State.Header.Type != OBU_TILE_GROUP &&
		o.State.Header.Type != OBU_TILE_LIST &&
		o.State.Header.Type != OBU_FRAME {
		b.TrailingBits(o.Size*8 - payloadBits)
	}
}

// frame_obu( sz )
func (o *Obu) newFrame(sz int, b *bitstream.BitStream) {
	startBitPos := b.Position

	o.ParseFrameHeader(b)
	b.ByteAlignment()

	endBitPos := b.Position

	headerBytes := (endBitPos - startBitPos) / 8
	sz -= headerBytes
	_ = NewTileGroup(b, sz)
}

// frame_header_obu()
func (o *Obu) ParseFrameHeader(b *bitstream.BitStream) {
	if o.State.SeenFrameHeader {
		FrameHeaderCopy()
	} else {
		o.State.SeenFrameHeader = true

		inputState := o.State.newUncompressedHeaderState()
		uncompressedHeader := uncompressedheader.NewUncompressedHeader(b, inputState)

		if uncompressedHeader.ShowExistingFrame {
			uncompressedHeader.DecodeFrameWrapup()
			o.State.SeenFrameHeader = false
		} else {
			o.State.TileNum = 0
			o.State.SeenFrameHeader = true
		}

		o.State.update(uncompressedHeader.State)
	}
}

// frame_header_copy()
func FrameHeaderCopy() {
	panic("not implemented")
}