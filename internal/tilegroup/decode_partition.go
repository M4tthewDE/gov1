package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/symbol"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// decode_partition(r, c, bSize)
func (t *TileGroup) decodePartition(r int, c int, bSize int, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	if r >= state.MiRows || c >= state.MiCols {
		return
	}

	state.AvailU = t.isInside(r-1, c, state)
	state.AvailL = t.isInside(r, c-1, state)
	num4x4 := shared.NUM_4X4_BLOCKS_WIDE[bSize]
	halfBlock4x4 := num4x4 >> 1
	quarterBlock4x4 := halfBlock4x4 >> 1
	hasRows := (r + halfBlock4x4) < state.MiRows
	hasCols := (c + halfBlock4x4) < state.MiCols

	var partition int
	if bSize < shared.BLOCK_8X8 {
		partition = PARTITION_NONE
	} else if hasRows && hasCols {
		partition = partitionElement(bSize, r, c, state, uh, b)
	} else if hasCols {
		splitOrHorz := util.Bool(splitOrHorzElement(bSize, r, c, state, uh, b))
		if splitOrHorz {
			partition = PARTITION_SPLIT
		} else {
			partition = PARTITION_HORZ
		}
	} else if hasRows {
		splitOrVert := util.Bool(splitOrVertElement(bSize, r, c, state, uh, b))
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

func splitOrVertElement(bSize int, r int, c int, state *state.State, uh uncompressedheader.UncompressedHeader, b *bitstream.BitStream) int {
	bsl := shared.MI_WIDTH_LOG2[bSize]
	above := state.AvailU && (shared.MI_WIDTH_LOG2[state.MiSizes[r-1][c]] < bsl)
	left := state.AvailL && (shared.MI_HEIGHT_LOG2[state.MiSizes[r][c-1]] < bsl)
	ctx := util.Int(left)*2 + util.Int(above)

	var partitionCdf []int
	if bsl == 1 {
		partitionCdf = state.TilePartitionW8Cdf[ctx]
	} else if bsl == 2 {
		partitionCdf = state.TilePartitionW16Cdf[ctx]
	} else if bsl == 3 {
		partitionCdf = state.TilePartitionW32Cdf[ctx]
	} else if bsl == 4 {
		partitionCdf = state.TilePartitionW64Cdf[ctx]
	} else {
		partitionCdf = state.TilePartitionW128Cdf[ctx]
	}

	psum := (partitionCdf[PARTITION_HORZ] - partitionCdf[PARTITION_HORZ-1] +
		partitionCdf[PARTITION_SPLIT] - partitionCdf[PARTITION_SPLIT-1] +
		partitionCdf[PARTITION_HORZ_A] - partitionCdf[PARTITION_HORZ_A-1] +
		partitionCdf[PARTITION_HORZ_B] - partitionCdf[PARTITION_HORZ_B-1] +
		partitionCdf[PARTITION_VERT_A] - partitionCdf[PARTITION_VERT_A-1])

	if bSize != shared.BLOCK_128X128 {
		psum += partitionCdf[PARTITION_HORZ_4] - partitionCdf[PARTITION_HORZ_4-1]
	}

	cdf := make([]int, 3)
	cdf[0] = (1 << 15) - psum
	cdf[1] = 1 << 15
	cdf[3] = 0

	return symbol.ReadSymbol(cdf, state, b, uh)
}

func splitOrHorzElement(bSize int, r int, c int, state *state.State, uh uncompressedheader.UncompressedHeader, b *bitstream.BitStream) int {
	bsl := shared.MI_WIDTH_LOG2[bSize]
	above := state.AvailU && (shared.MI_WIDTH_LOG2[state.MiSizes[r-1][c]] < bsl)
	left := state.AvailL && (shared.MI_HEIGHT_LOG2[state.MiSizes[r][c-1]] < bsl)
	ctx := util.Int(left)*2 + util.Int(above)

	var partitionCdf []int
	if bsl == 1 {
		partitionCdf = state.TilePartitionW8Cdf[ctx]
	} else if bsl == 2 {
		partitionCdf = state.TilePartitionW16Cdf[ctx]
	} else if bsl == 3 {
		partitionCdf = state.TilePartitionW32Cdf[ctx]
	} else if bsl == 4 {
		partitionCdf = state.TilePartitionW64Cdf[ctx]
	} else {
		partitionCdf = state.TilePartitionW128Cdf[ctx]
	}

	psum := (partitionCdf[PARTITION_VERT] - partitionCdf[PARTITION_VERT-1] +
		partitionCdf[PARTITION_SPLIT] - partitionCdf[PARTITION_SPLIT-1] +
		partitionCdf[PARTITION_HORZ_A] - partitionCdf[PARTITION_HORZ_A-1] +
		partitionCdf[PARTITION_VERT_A] - partitionCdf[PARTITION_VERT_A-1] +
		partitionCdf[PARTITION_VERT_B] - partitionCdf[PARTITION_VERT_B-1])

	if bSize != shared.BLOCK_128X128 {
		psum += partitionCdf[PARTITION_VERT_4] - partitionCdf[PARTITION_VERT_4-1]
	}

	cdf := make([]int, 3)
	cdf[0] = (1 << 15) - psum
	cdf[1] = 1 << 15
	cdf[3] = 0

	return symbol.ReadSymbol(cdf, state, b, uh)
}

func partitionElement(bSize int, r int, c int, state *state.State, uh uncompressedheader.UncompressedHeader, b *bitstream.BitStream) int {
	bsl := shared.MI_WIDTH_LOG2[bSize]
	above := state.AvailU && (shared.MI_WIDTH_LOG2[state.MiSizes[r-1][c]] < bsl)
	left := state.AvailL && (shared.MI_HEIGHT_LOG2[state.MiSizes[r][c-1]] < bsl)
	ctx := util.Int(left)*2 + util.Int(above)

	var cdf []int
	if bsl == 1 {
		cdf = state.TilePartitionW8Cdf[ctx]
	} else if bsl == 2 {
		cdf = state.TilePartitionW16Cdf[ctx]
	} else if bsl == 3 {
		cdf = state.TilePartitionW32Cdf[ctx]
	} else if bsl == 4 {
		cdf = state.TilePartitionW64Cdf[ctx]
	} else {
		cdf = state.TilePartitionW128Cdf[ctx]
	}

	return symbol.ReadSymbol(cdf, state, b, uh)
}
