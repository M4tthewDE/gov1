package main

type ObuType int

const (
	ObuSequenceHeader       ObuType = 1
	ObuTemporalLimiter      ObuType = 2
	ObuFrameHeader          ObuType = 3
	ObuTileGroup            ObuType = 4
	ObuMetadata             ObuType = 5
	ObuFrame                ObuType = 6
	ObuRedundantFrameHeader ObuType = 7
	ObuTileList             ObuType = 8
	ObuPadding              ObuType = 15
)

type ObuHeader struct {
	ForbiddenBit  bool
	Type          ObuType
	ExtensionFlag bool
	HasSizeField  bool
	ReservedBit   bool
}

func ParseObuHeader(data byte) ObuHeader {
	return ObuHeader{
		ForbiddenBit:  data&128 != 0,
		Type:          ObuType(int(data & 127 >> 3)),
		ExtensionFlag: data&4 != 0,
		HasSizeField:  data&2 != 0,
		ReservedBit:   data&1 != 0,
	}
}

type ObuExtensionHeader struct {
	TemporalID    int
	SpatialID     int
	Reserved3Bits int
}

func ParseObuExtensionHeader(data byte) ObuExtensionHeader {
	return ObuExtensionHeader{
		TemporalID:    int(data >> 5),
		SpatialID:     int(data & 24 >> 3),
		Reserved3Bits: int(data & 7),
	}
}
