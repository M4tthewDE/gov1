package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/literal"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/symbol"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// intra_frame_mode_info()
func (t *TileGroup) intraFrameModeInfo(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State, sh sequenceheader.SequenceHeader) {
	t.Skip = 0
	if uh.SegIdPreSkip == 1 {
		t.intraSegmentId(b, uh, state)
	}

	t.SkipMode = 0
	t.readSkip(b, uh, state)

	if !util.Bool(uh.SegIdPreSkip) {
		t.intraSegmentId(b, uh, state)
	}
	t.readCdef(b, uh, sh, state)
	t.readDeltaQIndex(b, sh, uh, state)
	t.readDeltaLf(b, sh, uh, state)

	state.ReadDeltas = false
	state.RefFrame[0] = shared.INTRA_FRAME
	state.RefFrame[0] = shared.NONE

	if uh.AllowIntraBc {
		t.useIntrabc = b.S()
	} else {
		t.useIntrabc = 0
	}

	if util.Bool(t.useIntrabc) {
		t.IsInter = -1
		t.YMode = DC_PRED
		t.UVMode = DC_PRED
		t.MotionMode = SIMPLE
		t.CompoundType = COMPUND_AVERAGE
		t.PaletteSizeY = 0
		t.PaletteSizeUV = 0
		t.InterpFilter[0] = shared.BILINEAR
		t.InterpFilter[1] = shared.BILINEAR
		t.findMvStack(false, state, uh)
		t.assignMv(0, b, state, sh, uh)
	} else {
		t.IsInter = 0
		intraFrameYMode := t.intraFrameYModeSymbol(state, b, uh)
		t.YMode = intraFrameYMode
		t.intraAngleInfoY(b, state, uh)

		if t.HasChroma {
			var uvMode int
			if t.Lossless && t.getPlaneResidualSize(state.MiSize, 1, sh) == shared.BLOCK_4X4 {
				uvMode = symbol.ReadSymbol(state.TileUVModeCflAllowedCdf[t.YMode], state, b, uh)
			} else if !t.Lossless && util.Max(shared.BLOCK_WIDTH[state.MiSize], shared.BLOCK_HEIGHT[state.MiSize]) < 32 {
				uvMode = symbol.ReadSymbol(state.TileUVModeCflAllowedCdf[t.YMode], state, b, uh)
			} else {
				uvMode = symbol.ReadSymbol(state.TileUVModeCflNotAllowedCdf[t.YMode], state, b, uh)
			}

			t.UVMode = uvMode

			if t.UVMode == UV_CFL_PRED {
				t.readCflAlphas(b, state, uh)
			}

			t.intraAngleInfoUv(b, state)
		}

		t.PaletteSizeY = 0
		t.PaletteSizeUV = 0

		if state.MiSize >= shared.BLOCK_8X8 && shared.BLOCK_WIDTH[state.MiSize] <= 64 && shared.BLOCK_HEIGHT[state.MiSize] <= 64 && util.Bool(uh.AllowScreenContentTools) {
			t.paletteModeInfo(b, state, sh, uh)
		}
		t.filterIntraModeInfo(b, sh, state)
	}
}

var INTRA_MODE_CONTEXT = []int{
	0, 1, 2, 3, 4, 4, 4, 4, 3, 0, 1, 2, 0,
}

func (t *TileGroup) intraFrameYModeSymbol(state *state.State, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	var aboveMode int
	var leftMode int

	if state.AvailU {
		aboveMode = INTRA_MODE_CONTEXT[t.YModes[state.MiRow-1][state.MiCol]]
	} else {
		aboveMode = DC_PRED
	}

	if state.AvailL {
		leftMode = INTRA_MODE_CONTEXT[t.YModes[state.MiRow][state.MiCol-1]]
	} else {
		leftMode = DC_PRED
	}

	return symbol.ReadSymbol(state.TileIntraFrameYModeCdf[aboveMode][leftMode], state, b, uh)
}

// read_segment_id()
func (t *TileGroup) readSegmentId(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State) {
	var prevU int
	var prevL int
	var prevUL int
	var pred int
	if state.AvailU && state.AvailL {
		prevUL = t.SegmentIds[state.MiRow-1][state.MiCol-1]
	} else {
		prevUL = -1
	}

	if state.AvailU {
		prevU = t.SegmentIds[state.MiRow-1][state.MiCol]
	} else {
		prevU = -1
	}

	if state.AvailL {
		prevL = t.SegmentIds[state.MiRow][state.MiCol-1]
	} else {
		prevL = -1
	}

	if prevU == -1 {
		if prevL == -1 {
			pred = 0
		} else {
			pred = prevL
		}
	} else if prevL == -1 {
		pred = prevU
	} else {
		if prevUL == prevU {
			pred = prevU
		} else {
			pred = prevL
		}
	}

	if t.Skip == 1 {
		t.SegmentId = pred
	} else {
		var ctx int
		if prevUL < 0 {
			ctx = 0
		} else if (prevUL == prevU) && prevUL == prevL {
			ctx = 2
		} else if prevUL == prevU || prevUL == prevL || prevU == prevL {
			ctx = 1
		} else {
			ctx = 0
		}

		t.SegmentId = symbol.ReadSymbol(state.TileSegmentIdCdf[ctx], state, b, uh)
		t.SegmentId = util.NegDeinterleave(t.SegmentId, pred, uh.LastActiveSegId+1)
	}
}

// read_skip()
func (t *TileGroup) readSkip(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State) {
	if (uh.SegIdPreSkip == 1) && t.segFeatureActive(shared.SEG_LVL_SKIP, uh, state) {
		t.Skip = 1
	} else {
		ctx := 0
		if state.AvailU {
			ctx += t.SkipModes[state.MiRow-1][state.MiCol]
		}
		if state.AvailL {
			ctx += t.SkipModes[state.MiRow][state.MiCol-1]
		}

		t.Skip = symbol.ReadSymbol(state.TileSkipCdf[ctx], state, b, uh)
	}
}

// seg_feature_active( feature )
func (t *TileGroup) segFeatureActive(feature int, uh uncompressedheader.UncompressedHeader, state *state.State) bool {
	return t.segFeatureActiveIdx(t.SegmentId, feature, uh, state)
}

// seg_feature_active_idx( idx, feature )
func (t *TileGroup) segFeatureActiveIdx(idx int, feature int, uh uncompressedheader.UncompressedHeader, state *state.State) bool {
	return uh.SegmentationEnabled && (state.FeatureEnabled[idx][feature] == 1)
}

// intra_segment_id()
func (t *TileGroup) intraSegmentId(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State) {
	if uh.SegmentationEnabled {
		t.readSegmentId(b, uh, state)
	} else {
		t.SegmentId = 0
	}

	t.Lossless = uh.LosslessArray[t.SegmentId]
}

// read_cdef()
func (t *TileGroup) readCdef(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader, state *state.State) {
	if util.Bool(t.Skip) || uh.CodedLossless || !sh.EnableCdef || uh.AllowIntraBc {
		return
	}

	cdefSize4 := shared.NUM_4X4_BLOCKS_WIDE[shared.BLOCK_64X64]
	cdefMask4 := ^(cdefSize4 - 1)
	r := state.MiRow & cdefMask4
	c := state.MiCol & cdefMask4

	if state.Cdef.CdefIdx[r][c] == -1 {
		state.Cdef.CdefIdx[r][c] = literal.L(state.Cdef.CdefBits, state, b, uh)
		w4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
		h4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]

		for i := r; i < r+h4; i += cdefSize4 {
			for j := c; i < c+w4; i += cdefSize4 {
				state.Cdef.CdefIdx[i][j] = state.Cdef.CdefIdx[r][c]
			}

		}
	}
}

// read_delta_qindex()
func (t *TileGroup) readDeltaQIndex(b *bitstream.BitStream, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader, state *state.State) {
	var sbSize int
	if sh.Use128x128SuperBlock {
		sbSize = shared.BLOCK_128X128
	} else {
		sbSize = shared.BLOCK_64X64
	}

	if state.MiSize == sbSize && util.Bool(t.Skip) {
		return
	}

	if state.ReadDeltas {
		deltaQAbs := symbol.ReadSymbol(state.TileDeltaQCdf, state, b, uh)
		if deltaQAbs == DELTA_Q_SMALL {
			deltaQRemBits := literal.L(3, state, b, uh)
			deltaQRemBits++
			deltaQAbsBits := literal.L(deltaQRemBits, state, b, uh)
			deltaQAbs = deltaQAbsBits + (1 << deltaQRemBits) + 1
		}

		if util.Bool(deltaQAbs) {
			deltaQSignBit := literal.L(1, state, b, uh)
			var reducedDeltaQIndex int
			if util.Bool(deltaQSignBit) {
				reducedDeltaQIndex = -deltaQAbs
			} else {
				reducedDeltaQIndex = deltaQAbs
			}

			state.CurrentQIndex = util.Clip3(1, 255, state.CurrentQIndex+(reducedDeltaQIndex<<uh.DeltaQRes))

		}
	}
}

// read_delta_lf()
func (t *TileGroup) readDeltaLf(b *bitstream.BitStream, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader, state *state.State) {
	var sbSize int
	if sh.Use128x128SuperBlock {
		sbSize = shared.BLOCK_128X128
	} else {
		sbSize = shared.BLOCK_64X64
	}

	if state.MiSize == sbSize && util.Bool(t.Skip) {
		return
	}

	if state.ReadDeltas && uh.DeltaLfPresent {
		frameLfCount := 1

		if util.Bool(uh.DeltaLfMulti) {
			if sh.ColorConfig.NumPlanes > 1 {
				frameLfCount = shared.FRAME_LF_COUNT
			} else {
				frameLfCount = shared.FRAME_LF_COUNT - 2
			}
		}

		for i := 0; i < frameLfCount; i++ {
			var deltaLfAbs int
			var delta_lf_abs int
			if uh.DeltaLfMulti == 0 {
				delta_lf_abs = symbol.ReadSymbol(state.TileDeltaLFCdf, state, b, uh)
			} else {
				delta_lf_abs = symbol.ReadSymbol(state.TileDeltaLFMultiCdf[i], state, b, uh)
			}

			if delta_lf_abs == DELTA_LF_SMALL {
				deltaLfRemBits := literal.L(3, state, b, uh)
				n := deltaLfRemBits + 1
				deltaLfAbsBits := literal.L(n, state, b, uh)
				deltaLfAbs = deltaLfAbsBits + (1 << n) + 1
			} else {
				deltaLfAbs = delta_lf_abs
			}

			var reducedDeltaLfLevel int
			if util.Bool(deltaLfAbs) {
				deltaLfSignBit := literal.L(1, state, b, uh)
				if util.Bool(deltaLfSignBit) {
					reducedDeltaLfLevel = -deltaLfAbs
				} else {
					reducedDeltaLfLevel = deltaLfAbs

				}

				state.DeltaLF[i] = util.Clip3(-shared.MAX_LOOP_FILTER, shared.MAX_LOOP_FILTER, state.DeltaLF[i]+(reducedDeltaLfLevel<<uh.DeltaLfRes))
			}
		}
	}
}

// intra_angle_info_y()
func (t *TileGroup) intraAngleInfoY(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	t.AngleDeltaY = 0

	if state.MiSize >= shared.BLOCK_8X8 {

		if t.isDirectionalMode(t.YMode) {
			angleDeltaY := symbol.ReadSymbol(state.TileAngleDeltaCdf[t.YMode-V_PRED], state, b, uh)
			t.AngleDeltaY = angleDeltaY - MAX_ANGLE_DELTA
		}
	}
}

// read_cfl_alphas()
func (t *TileGroup) readCflAlphas(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	cflAlphaSigns := symbol.ReadSymbol(state.TileCflSignCdf, state, b, uh)
	signU := (cflAlphaSigns + 1) / 3
	signV := (cflAlphaSigns + 1) % 3

	if signU != CFL_SIGN_ZERO {
		ctx := (signU-1)*3 + signV
		cflAlphaU := symbol.ReadSymbol(state.TileCflAlphaCdf[ctx], state, b, uh)

		t.CflAlphaU = 1 + cflAlphaU
		if signU == CFL_SIGN_NEG {
			t.CflAlphaU = -t.CflAlphaU
		}
	} else {
		t.CflAlphaU = 0
	}

	if signV != CFL_SIGN_ZERO {
		ctx := (signV-1)*3 + signU
		cflAlphaV := symbol.ReadSymbol(state.TileCflAlphaCdf[ctx], state, b, uh)

		t.CflAlphaV = 1 + cflAlphaV
		if signV == CFL_SIGN_NEG {
			t.CflAlphaV = -t.CflAlphaV
		}
	} else {
		t.CflAlphaV = 0
	}

}

// intra_angle_info_uv()
func (t *TileGroup) intraAngleInfoUv(b *bitstream.BitStream, state *state.State) {
	t.AngleDeltaUV = 0
	if state.MiSize >= shared.BLOCK_8X8 {
		if t.isDirectionalMode(t.UVMode) {
			angleDeltaUv := b.S()
			t.AngleDeltaUV = angleDeltaUv - MAX_ANGLE_DELTA
		}
	}
}

// filter_intra_mode_info()
func (t *TileGroup) filterIntraModeInfo(b *bitstream.BitStream, sh sequenceheader.SequenceHeader, state *state.State) {
	useFilterIntra := false
	if sh.EnableFilterIntra && t.YMode == DC_PRED && t.PaletteSizeY == 0 && util.Max(shared.BLOCK_WIDTH[state.MiSize], shared.BLOCK_HEIGHT[state.MiSize]) <= 32 {
		useFilterIntra = util.Bool(b.S())

		if useFilterIntra {
			t.FilterIntraMode = b.S()
		}
	}
}
