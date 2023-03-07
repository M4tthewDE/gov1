package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
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
		t.findMvStack(0, state, uh)
		t.assignMv(0, b, state, sh, uh)
	} else {
		t.IsInter = 0
		intraFrameYMode := b.S()
		t.YMode = intraFrameYMode
		t.intraAngleInfoY(b, state)

		if t.HasChroma {
			uvMode := b.S()

			t.UVMode = uvMode

			if t.UVMode == UV_CFL_PRED {
				t.readCflAlphas(b)
			}

			t.intraAngleInfoUv(b, state)
		}

		t.PaletteSizeY = 0
		t.PaletteSizeUV = 0

		if state.MiSize >= shared.BLOCK_8X8 && t.Block_Width[state.MiSize] <= 64 && t.Block_Height[state.MiSize] <= 64 && util.Bool(uh.AllowScreenContentTools) {
			t.paletteModeInfo(b)
		}
		t.filterIntraModeInfo(b, sh, state)
	}
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
		t.SegmentId = b.S()
		t.SegmentId = util.NegDeinterleave(t.SegmentId, pred, uh.LastActiveSegId+1)
	}
}

// read_skip()
func (t *TileGroup) readSkip(b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State) {
	if (uh.SegIdPreSkip == 1) && t.segFeatureActive(shared.SEG_LVL_SKIP, uh, state) {
		t.Skip = 1
	} else {
		t.Skip = b.S()
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

	cdefSize4 := state.Num4x4BlocksWide[shared.BLOCK_64X64]
	cdefMask4 := ^(cdefSize4 - 1)
	r := state.MiRow & cdefMask4
	c := state.MiCol & cdefMask4

	if state.Cdef.CdefIdx[r][c] == -1 {
		state.Cdef.CdefIdx[r][c] = b.L(state.Cdef.CdefBits)
		w4 := state.Num4x4BlocksWide[state.MiSize]
		h4 := state.Num4x4BlocksHigh[state.MiSize]

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
		deltaQAbs := b.S()
		if deltaQAbs == DELTA_Q_SMALL {
			deltaQRemBits := b.L(3)
			deltaQRemBits++
			deltaQAbsBits := b.L(deltaQRemBits)
			deltaQAbs = deltaQAbsBits + (1 << deltaQRemBits) + 1
		}

		if util.Bool(deltaQAbs) {
			deltaQSignBit := b.L(1)
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
				frameLfCount = FRAME_LF_COUNT
			} else {
				frameLfCount = FRAME_LF_COUNT - 2
			}
		}

		for i := 0; i < frameLfCount; i++ {
			var deltaLfAbs int
			delta_lf_abs := b.S()

			if delta_lf_abs == DELTA_LF_SMALL {
				deltaLfRemBits := b.L(3)
				n := deltaLfRemBits + 1
				deltaLfAbsBits := b.L(n)
				deltaLfAbs = deltaLfAbsBits + (1 << n) + 1
			} else {
				deltaLfAbs = delta_lf_abs
			}

			var reducedDeltaLfLevel int
			if util.Bool(deltaLfAbs) {
				deltaLfSignBit := b.L(1)
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
func (t *TileGroup) intraAngleInfoY(b *bitstream.BitStream, state *state.State) {
	t.AngleDeltaY = 0

	if state.MiSize >= shared.BLOCK_8X8 {

		if t.isDirectionalMode(t.YMode) {
			angleDeltaY := b.S()
			t.AngleDeltaY = angleDeltaY - MAX_ANGLE_DELTA
		}
	}
}

// read_cfl_alphas()
func (t *TileGroup) readCflAlphas(b *bitstream.BitStream) {
	cflAlphaSigns := b.S()
	signU := (cflAlphaSigns + 1) / 3
	signV := (cflAlphaSigns + 1) % 3

	if signU != CFL_SIGN_ZERO {
		cflAlphaU := b.S()
		t.CflAlphaU = 1 + cflAlphaU
		if signU == CFL_SIGN_NEG {
			t.CflAlphaU = -t.CflAlphaU
		}
	} else {
		t.CflAlphaU = 0
	}

	if signV != CFL_SIGN_ZERO {
		cflAlphaV := b.S()
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
	if sh.EnableFilterIntra && t.YMode == DC_PRED && t.PaletteSizeY == 0 && util.Max(t.Block_Width[state.MiSize], t.Block_Height[state.MiSize]) <= 32 {
		useFilterIntra = util.Bool(b.S())

		if useFilterIntra {
			t.FilterIntraMode = b.S()
		}
	}
}
