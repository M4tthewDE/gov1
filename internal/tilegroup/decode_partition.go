package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
)

// decode_partition(r, c, bSize)
func (t *TileGroup) decodePartition(r int, c int, bSize int, b *bitstream.BitStream) {
	if r >= t.State.MiRows || c >= t.State.MiCols {
		return
	}

	t.State.AvailU = t.isInside(r-1, c)
	t.State.AvailL = t.isInside(r, c-1)
	num4x4 := t.State.Num4x4BlocksWide[bSize]
	halfBlock4x4 := num4x4 >> 1
	quarterBlock4x4 := halfBlock4x4 >> 1
	hasRows := (r + halfBlock4x4) < t.State.MiRows
	hasCols := (c + halfBlock4x4) < t.State.MiCols

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
		t.decodeBlock(r, c, subSize, p)
	} else if partition == PARTITION_HORZ {
		t.decodeBlock(r+halfBlock4x4, c, subSize, p)
		if hasRows {
			t.decodeBlock(r+halfBlock4x4, c, subSize, p)
		}
	} else if partition == PARTITION_VERT {
		t.decodeBlock(r, c, subSize, p)
		if hasCols {
			t.decodeBlock(r, c+halfBlock4x4, subSize, p)
		}
	} else if partition == PARTITION_SPLIT {
		t.decodePartition(r, c, subSize, p)
		t.decodePartition(r, c+halfBlock4x4, subSize, p)
		t.decodePartition(r+halfBlock4x4, c, subSize, p)
		t.decodePartition(r+halfBlock4x4, c+halfBlock4x4, subSize, p)
	} else if partition == PARTITION_HORZ_A {
		t.decodeBlock(r, c, splitSize, p)
		t.decodeBlock(r, c+halfBlock4x4, splitSize, p)
		t.decodeBlock(r+halfBlock4x4, c, splitSize, p)
	} else if partition == PARTITION_HORZ_B {
		t.decodeBlock(r, c, subSize, p)
		t.decodeBlock(r+halfBlock4x4, c, splitSize, p)
		t.decodeBlock(r+halfBlock4x4, c+halfBlock4x4, splitSize, p)
	} else if partition == PARTITION_VERT_A {
		t.decodeBlock(r, c, splitSize, p)
		t.decodeBlock(r+halfBlock4x4, c, splitSize, p)
		t.decodeBlock(r, c+halfBlock4x4, subSize, p)
	} else if partition == PARTITION_VERT_B {
		t.decodeBlock(r, c, subSize, p)
		t.decodeBlock(r, c+halfBlock4x4, splitSize, p)
		t.decodeBlock(r+halfBlock4x4, c+halfBlock4x4, subSize, p)
	} else if partition == PARTITION_HORZ_4 {
		t.decodeBlock(r+quarterBlock4x4*0, c, subSize, p)
		t.decodeBlock(r+quarterBlock4x4*1, c, subSize, p)
		t.decodeBlock(r+quarterBlock4x4*2, c, subSize, p)
		if r+quarterBlock4x4*3 < t.State.MiRows {
			t.decodeBlock(r+quarterBlock4x4*3, c, subSize, p)
		}
	} else {
		t.decodeBlock(r, c+quarterBlock4x4*0, subSize, p)
		t.decodeBlock(r, c+quarterBlock4x4*1, subSize, p)
		t.decodeBlock(r, c+quarterBlock4x4*2, subSize, p)
		if c+quarterBlock4x4*3 < t.State.MiRows {
			t.decodeBlock(r, c+quarterBlock4x4*3, subSize, p)
		}
	}
}
