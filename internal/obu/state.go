package obu

import (
	"github.com/m4tthewde/gov1/internal/header"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
)

type State struct {
	header            header.Header
	sequenceHeader    sequenceheader.SequenceHeader
	operatingPointIdc int
	seenFrameHeader   bool
}
