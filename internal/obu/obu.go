package obu

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/tilegroup"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

type ObuType int

type Obu struct {
	State State
	Size  int
}

func NewObu(sz int, state State, b *bitstream.BitStream) Obu {
	obu := Obu{}
	obu.State = state
	obu.build(sz, b)

	return obu
}

// open_bitstream_unit(sz)
func (o *Obu) build(sz int, b *bitstream.BitStream) {
	o.State.Header = header.NewHeader(b)

	if o.State.Header.HasSizeField {
		o.Size = b.Leb128()
	} else {
		o.Size = sz - 1 - util.Int(o.State.Header.ExtensionFlag)
	}

	startPosition := b.Position

	if o.State.Header.Type != header.OBU_SEQUENCE_HEADER &&
		o.State.Header.Type != header.OBU_TEMPORAL_DELIMITER &&
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

	switch o.State.Header.Type {
	case header.OBU_SEQUENCE_HEADER:
		sequenceheader, result := sequenceheader.NewSequenceHeader(b)
		o.State.SequenceHeader = sequenceheader
		o.State.OperatingPointIdc = result.OperatingPointIdc
	case header.OBU_TEMPORAL_DELIMITER:
		o.State.SeenFrameHeader = false
	case header.OBU_FRAME_HEADER:
		o.ParseFrameHeader(b)
	case header.OBU_REDUNDANT_FRAME_HEADER:
		o.ParseFrameHeader(b)
	case header.OBU_FRAME:
		o.newFrame(o.Size, b)
	case header.OBU_PADDING:
		o.paddingObu(b)
	default:
		panic("not implemented")
	}

	payloadBits := b.Position - startPosition

	if o.Size > 0 &&
		o.State.Header.Type != header.OBU_TILE_GROUP &&
		o.State.Header.Type != header.OBU_TILE_LIST &&
		o.State.Header.Type != header.OBU_FRAME {
		b.TrailingBits(o.Size*8 - payloadBits)
	}
}

// TODO: remove size, should be included in struct
// frame_obu( sz )
func (o *Obu) newFrame(sz int, b *bitstream.BitStream) {
	startBitPos := b.Position

	o.ParseFrameHeader(b)
	b.ByteAlignment()

	endBitPos := b.Position

	headerBytes := (endBitPos - startBitPos) / 8
	sz -= headerBytes

	inputState := o.State.newTileGroupState()
	_ = tilegroup.NewTileGroup(sz, b, inputState)
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

// padding_obu( )
func (o *Obu) paddingObu(b *bitstream.BitStream) {
	obuPaddingByte := 1
	for obuPaddingByte != 0 {
		obuPaddingByte = b.F(8)
	}

	b.Position -= 8
}
