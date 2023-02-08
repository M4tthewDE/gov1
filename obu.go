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
	Header ObuHeader
	Size   int
}

type ObuFrame struct{}

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
		inTemporalLayer := (p.operatingPointIdc >> obu.Header.ObuExtensionHeader.TemporalID) & 1
		inSpatialLayer := (p.operatingPointIdc >> (obu.Header.ObuExtensionHeader.SpatialID + 8)) & 1

		if !(inTemporalLayer != 0) || !(inSpatialLayer != 0) {
			//drop_obu()
			p.position = p.position + obu.Size*8
			return
		}
	}

	x, _ := json.MarshalIndent(obu, "", "	")
	fmt.Printf("%s\n", string(x))

	switch obu.Header.Type {
	case SequenceHeader:
		sequenceHeader := p.ParseObuSequenceHeader()
		x, _ := json.MarshalIndent(sequenceHeader, "", "	")
		fmt.Printf("%s\n", string(x))
	case Frame:
		frame := p.ParseFrame(obu.Size)
		x, _ := json.MarshalIndent(frame, "", "	")
		fmt.Printf("%s\n", string(x))
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
func (p *Parser) ParseFrame(sz int) ObuFrame {
	startBitPos := p.position
	frameHeader := p.ParseFrameHeader()

	return ObuFrame{}
}

// frame_header_obu()
func (p *Parser) ParseFrameHeader(sequenceHeader ObuSequenceHeader) {
	if p.seenFrameHeader {
		p.FrameHeaderCopy()
	} else {
		p.seenFrameHeader = true
		uncompressedHeader := p.UncompressedHeader(sequenceHeader)

		if uncompressedHeader.ShowExistingFrame {
			p.DecodeFrameWrapup()
			p.seenFrameHeader = false
		} else {
			p.tileNum = 0
			p.seenFrameHeader = true
		}

	}
}
