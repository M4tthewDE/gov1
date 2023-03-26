package obu

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/logger"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/tilegroup"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
	"go.uber.org/zap"
)

type ObuType int

type Obu struct {
	Size               int
	Header             header.Header
	SequenceHeader     sequenceheader.SequenceHeader
	UncompressedHeader uncompressedheader.UncompressedHeader
	TileGroup          tilegroup.TileGroup
}

func NewObu(sz int, state *state.State, b *bitstream.BitStream, sh sequenceheader.SequenceHeader) Obu {
	obu := Obu{}
	obu.SequenceHeader = sh
	obu.build(sz, b, state)

	return obu
}

// open_bitstream_unit(sz)
func (o *Obu) build(sz int, b *bitstream.BitStream, state *state.State) {
	o.Header = header.NewHeader(b)

	if o.Header.HasSizeField {
		o.Size = b.Leb128()
	} else {
		o.Size = sz - 1 - util.Int(o.Header.ExtensionFlag)
	}

	startPosition := b.Position

	if o.Header.Type != header.OBU_SEQUENCE_HEADER &&
		o.Header.Type != header.OBU_TEMPORAL_DELIMITER &&
		state.OperatingPointIdc != 0 &&
		o.Header.ExtensionFlag {
		inTemporalLayer := ((state.OperatingPointIdc >> o.Header.ExtensionHeader.TemporalID) & 1) != 0
		inSpatialLayer := ((state.OperatingPointIdc >> (o.Header.ExtensionHeader.SpatialID + 8)) & 1) != 0

		if !inTemporalLayer || !inSpatialLayer {
			//drop_obu()
			b.Position = b.Position + o.Size*8
			return
		}
	}

	logger.Logger.Info("Parsing obu...", zap.Int("type", o.Header.Type))

	switch o.Header.Type {
	case header.OBU_SEQUENCE_HEADER:
		o.SequenceHeader = sequenceheader.NewSequenceHeader(b, state)
	case header.OBU_TEMPORAL_DELIMITER:
		state.SeenFrameHeader = false
	case header.OBU_FRAME_HEADER:
		o.UncompressedHeader = o.ParseFrameHeader(b, state, o.Header, o.SequenceHeader)
	case header.OBU_REDUNDANT_FRAME_HEADER:
		o.ParseFrameHeader(b, state, o.Header, o.SequenceHeader)
	case header.OBU_FRAME:
		o.TileGroup = o.newFrame(o.Size, b, state, o.Header, o.SequenceHeader)
	case header.OBU_PADDING:
		o.paddingObu(b)
	default:
		panic("not implemented")
	}

	payloadBits := b.Position - startPosition

	if o.Size > 0 &&
		o.Header.Type != header.OBU_TILE_GROUP &&
		o.Header.Type != header.OBU_TILE_LIST &&
		o.Header.Type != header.OBU_FRAME {
		b.TrailingBits(o.Size*8 - payloadBits)
	}
}

// TODO: remove size, should be included in struct
// frame_obu( sz )
func (o *Obu) newFrame(sz int, b *bitstream.BitStream, state *state.State, h header.Header, sh sequenceheader.SequenceHeader) tilegroup.TileGroup {
	startBitPos := b.Position

	uh := o.ParseFrameHeader(b, state, h, sh)
	o.UncompressedHeader = uh
	b.ByteAlignment()

	endBitPos := b.Position

	headerBytes := (endBitPos - startBitPos) / 8
	sz -= headerBytes

	return tilegroup.NewTileGroup(sz, b, state, uh, sh)
}

// frame_header_obu()
func (o *Obu) ParseFrameHeader(b *bitstream.BitStream, state *state.State, h header.Header, sh sequenceheader.SequenceHeader) uncompressedheader.UncompressedHeader {
	if state.SeenFrameHeader {
		return FrameHeaderCopy()
	} else {
		state.SeenFrameHeader = true

		uncompressedHeader := uncompressedheader.NewUncompressedHeader(h, sh, b, state)

		if uncompressedHeader.ShowExistingFrame {
			uncompressedHeader.DecodeFrameWrapup()
			state.SeenFrameHeader = false
		} else {
			state.TileNum = 0
			state.SeenFrameHeader = true
		}

		return uncompressedHeader
	}
}

// frame_header_copy()
func FrameHeaderCopy() uncompressedheader.UncompressedHeader {
	panic("not implemented")
	return uncompressedheader.UncompressedHeader{}
}

// padding_obu( )
func (o *Obu) paddingObu(b *bitstream.BitStream) {
	obuPaddingByte := 1
	for obuPaddingByte != 0 {
		obuPaddingByte = b.F(8)
	}

	b.Position -= 8
}
