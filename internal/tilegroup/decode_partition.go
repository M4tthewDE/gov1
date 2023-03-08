package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
)

// decode_partition(r, c, bSize)
func (t *TileGroup) decodePartition(r int, c int, bSize int, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	if r >= state.MiRows || c >= state.MiCols {
		return
	}

	state.AvailU = t.isInside(r-1, c, state)
	state.AvailL = t.isInside(r, c-1, state)
	num4x4 := state.Num4x4BlocksWide[bSize]
	halfBlock4x4 := num4x4 >> 1
	quarterBlock4x4 := halfBlock4x4 >> 1
	hasRows := (r + halfBlock4x4) < state.MiRows
	hasCols := (c + halfBlock4x4) < state.MiCols

	var partition int
	if bSize < shared.BLOCK_8X8 {
		partition = PARTITION_NONE
	} else if hasRows && hasCols {
		partition = b.S()
	} else if hasCols {
		splitOrHorz := b.S() != 0
		if splitOrHorz {
			partition = PARTITION_SPLIT
		} else {
			partition = PARTITION_HORZ
		}
	} else if hasRows {
		splitOrVert := b.S() != 0
		if splitOrVert {
			partition = PARTITION_SPLIT
		} else {
			partition = PARTITION_VERT
		}

	} else {
		partition = PARTITION_SPLIT
	}
	subSize := shared.Partition_Subsize[partition][bSize]
	splitSize := shared.Partition_Subsize[PARTITION_SPLIT][bSize]
	if partition == PARTITION_NONE {
		t.decodeBlock(r, c, subSize, b, state, sh, uh)
	} else if partition == PARTITION_HORZ {
		t.decodeBlock(r+halfBlock4x4, c, subSize, b, state, sh, uh)
		if hasRows {
			t.decodeBlock(r+halfBlock4x4, c, subSize, b, state, sh, uh)
		}
	} else if partition == PARTITION_VERT {
		t.decodeBlock(r, c, subSize, b, state, sh, uh)
		if hasCols {
			t.decodeBlock(r, c+halfBlock4x4, subSize, b, state, sh, uh)
		}
	} else if partition == PARTITION_SPLIT {
		t.decodePartition(r, c, subSize, b, state, sh, uh)
		t.decodePartition(r, c+halfBlock4x4, subSize, b, state, sh, uh)
		t.decodePartition(r+halfBlock4x4, c, subSize, b, state, sh, uh)
		t.decodePartition(r+halfBlock4x4, c+halfBlock4x4, subSize, b, state, sh, uh)
	} else if partition == PARTITION_HORZ_A {
		t.decodeBlock(r, c, splitSize, b, state, sh, uh)
		t.decodeBlock(r, c+halfBlock4x4, splitSize, b, state, sh, uh)
		t.decodeBlock(r+halfBlock4x4, c, splitSize, b, state, sh, uh)
	} else if partition == PARTITION_HORZ_B {
		t.decodeBlock(r, c, subSize, b, state, sh, uh)
		t.decodeBlock(r+halfBlock4x4, c, splitSize, b, state, sh, uh)
		t.decodeBlock(r+halfBlock4x4, c+halfBlock4x4, splitSize, b, state, sh, uh)
	} else if partition == PARTITION_VERT_A {
		t.decodeBlock(r, c, splitSize, b, state, sh, uh)
		t.decodeBlock(r+halfBlock4x4, c, splitSize, b, state, sh, uh)
		t.decodeBlock(r, c+halfBlock4x4, subSize, b, state, sh, uh)
	} else if partition == PARTITION_VERT_B {
		t.decodeBlock(r, c, subSize, b, state, sh, uh)
		t.decodeBlock(r, c+halfBlock4x4, splitSize, b, state, sh, uh)
		t.decodeBlock(r+halfBlock4x4, c+halfBlock4x4, subSize, b, state, sh, uh)
	} else if partition == PARTITION_HORZ_4 {
		t.decodeBlock(r+quarterBlock4x4*0, c, subSize, b, state, sh, uh)
		t.decodeBlock(r+quarterBlock4x4*1, c, subSize, b, state, sh, uh)
		t.decodeBlock(r+quarterBlock4x4*2, c, subSize, b, state, sh, uh)
		if r+quarterBlock4x4*3 < state.MiRows {
			t.decodeBlock(r+quarterBlock4x4*3, c, subSize, b, state, sh, uh)
		}
	} else {
		t.decodeBlock(r, c+quarterBlock4x4*0, subSize, b, state, sh, uh)
		t.decodeBlock(r, c+quarterBlock4x4*1, subSize, b, state, sh, uh)
		t.decodeBlock(r, c+quarterBlock4x4*2, subSize, b, state, sh, uh)
		if c+quarterBlock4x4*3 < state.MiRows {
			t.decodeBlock(r, c+quarterBlock4x4*3, subSize, b, state, sh, uh)
		}
	}
}
