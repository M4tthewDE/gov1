package parser

import (
	"encoding/json"
	"fmt"
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
func (p *Parser) parseObu(sz int) {
	obu := Obu{}

	p.header = NewHeader(p)

	if p.header.HasSizeField {
		obu.Size = p.leb128()
	} else {
		extensionFlagInt := 0
		if p.header.ExtensionFlag {
			extensionFlagInt = 1
		}
		obu.Size = sz - 1 - extensionFlagInt
	}

	startPosition := p.position

	if p.header.Type != OBU_SEQUENCE_HEADER &&
		p.header.Type != OBU_TEMPORAL_DELIMITER &&
		p.operatingPointIdc != 0 &&
		p.header.ExtensionFlag {
		inTemporalLayer := ((p.operatingPointIdc >> p.header.ExtensionHeader.TemporalID) & 1) != 0
		inSpatialLayer := ((p.operatingPointIdc >> (p.header.ExtensionHeader.SpatialID + 8)) & 1) != 0

		if !inTemporalLayer || !inSpatialLayer {
			//drop_obu()
			p.position = p.position + obu.Size*8
			return
		}
	}

	x, _ := json.MarshalIndent(obu, "", "	")
	fmt.Printf("%s\n", string(x))

	switch p.header.Type {
	case OBU_SEQUENCE_HEADER:
		p.sequenceHeader = NewSequenceHeader(p)

		x, _ := json.MarshalIndent(p.sequenceHeader, "", "	")
		fmt.Printf("%s\n", string(x))
	case OBU_TEMPORAL_DELIMITER:
		p.seenFrameHeader = false
	case OBU_FRAME:
		p.ParseFrame(obu.Size)
	case OBU_METADATA:

	default:
		fmt.Printf("not implemented type %d\n", p.header.Type)
		panic("")
	}

	payloadBits := p.position - startPosition

	fmt.Println("----------------------------------------")
	fmt.Printf("p.position: %d\n", p.position)
	fmt.Printf("startPosition: %d\n", startPosition)
	fmt.Printf("payloadBits: %d\n", payloadBits)
	fmt.Printf("obu.Size*8 - payloadBits: %d\n", obu.Size*8-payloadBits)
	fmt.Println("----------------------------------------")

	if obu.Size > 0 &&
		p.header.Type != OBU_TILE_GROUP &&
		p.header.Type != OBU_TILE_LIST &&
		p.header.Type != OBU_FRAME {
		p.trailingBits(obu.Size*8 - payloadBits)
	}

}

// frame_obu( sz )
func (p *Parser) ParseFrame(sz int) {
	startBitPos := p.position

	p.ParseFrameHeader()
	p.byteAlignment()

	endBitPos := p.position

	headerBytes := (endBitPos - startBitPos) / 8
	sz -= headerBytes
	_ = NewTileGroup(p, sz)
}

// frame_header_obu()
func (p *Parser) ParseFrameHeader() {
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
func (p *Parser) FrameHeaderCopy() {
	panic("not implemented")
}
