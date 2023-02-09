package main

import (
	"encoding/json"
	"fmt"
)

type ObuType int

const (
	SequenceHeader       ObuType = 1
	TemporalLimiter      ObuType = 2
	FrameHeader          ObuType = 3
	TileGroup            ObuType = 4
	Metadata             ObuType = 5
	Frame                ObuType = 6
	RedundantFrameHeader ObuType = 7
	TileList             ObuType = 8
	Padding              ObuType = 15
)

type Obu struct {
	Header         ObuHeader
	SequenceHeader ObuSequenceHeader
	Size           int
}

// temporal_unit( sz )
func (p *Parser) temporalUnit(sz int) {
	for sz > 0 {
		frameUnitSize := p.leb128()
		sz -= p.leb128Bytes
		p.frameUnit(frameUnitSize)
		sz -= frameUnitSize
	}
}

// frame_unit( sz )
func (p *Parser) frameUnit(sz int) {
	for sz > 0 {
		obuLength := p.leb128()
		sz -= p.leb128Bytes
		p.ParseObu(obuLength)
		sz -= obuLength

	}
}

// open_bitstream_unit(sz)
func (p *Parser) ParseObu(sz int) {
	obu := Obu{}

	obu.Header = p.ParseObuHeader()

	if obu.Header.HasSizeField {
		obu.Size = p.leb128()
	} else {
		extensionFlagInt := 0
		if obu.Header.ExtensionFlag {
			extensionFlagInt = 1
		}
		obu.Size = sz - 1 - extensionFlagInt
	}

	startPosition := p.position

	if obu.Header.Type != SequenceHeader &&
		obu.Header.Type != TemporalLimiter &&
		p.operatingPointIdc != 0 &&
		obu.Header.ExtensionFlag {
		inTemporalLayer := ((p.operatingPointIdc >> obu.Header.ObuExtensionHeader.TemporalID) & 1) != 0
		inSpatialLayer := ((p.operatingPointIdc >> (obu.Header.ObuExtensionHeader.SpatialID + 8)) & 1) != 0

		if !inTemporalLayer || !inSpatialLayer {
			//drop_obu()
			p.position = p.position + obu.Size*8
			return
		}
	}

	x, _ := json.MarshalIndent(obu, "", "	")
	fmt.Printf("%s\n", string(x))

	switch obu.Header.Type {
	case SequenceHeader:
		obu.SequenceHeader = p.ParseObuSequenceHeader()

		x, _ := json.MarshalIndent(obu.SequenceHeader, "", "	")
		fmt.Printf("%s\n", string(x))
	case Frame:
		p.ParseFrame(obu.Size, obu.SequenceHeader, obu.Header.ObuExtensionHeader)
	default:
		fmt.Printf("not implemented type %d\n", obu.Header.Type)
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
		obu.Header.Type != TileGroup &&
		obu.Header.Type != TileList &&
		obu.Header.Type != Frame {
		p.trailingBits(obu.Size*8 - payloadBits)
	}

}

// frame_obu( sz )
func (p *Parser) ParseFrame(sz int, sequenceHeader ObuSequenceHeader, extensionHeader ObuExtensionHeader) {
	startBitPos := p.position

	p.ParseFrameHeader(sequenceHeader, extensionHeader)
	p.byteAlignment()

	endBitPos := p.position

	headerBytes := (endBitPos - startBitPos) / 8
	sz -= headerBytes
	p.tileGroupObu(sz)
}

// frame_header_obu()
func (p *Parser) ParseFrameHeader(sequenceHeader ObuSequenceHeader, extensionHeader ObuExtensionHeader) {
	if p.seenFrameHeader {
		p.FrameHeaderCopy()
	} else {
		p.seenFrameHeader = true
		uncompressedHeader := p.UncompressedHeader(sequenceHeader, extensionHeader)

		if uncompressedHeader.ShowExistingFrame {
			p.DecodeFrameWrapup()
			p.seenFrameHeader = false
		} else {
			p.tileNum = 0
			p.seenFrameHeader = true
		}

	}
}

// tile_group_obu(sz)
func (p *Parser) tileGroupObu(sz int) {
	panic("not implemented!")
}
