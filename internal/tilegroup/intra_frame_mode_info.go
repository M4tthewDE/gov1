package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
)

// intra_frame_mode_info()
func (t *TileGroup) intraFrameModeInfo(b *bitstream.BitStream) {
	t.Skip = 0
	if t.State.UncompressedHeader.SegIdPreSkip == 1 {
		t.intraSegmentId(b)
	}

	t.SkipMode = 0
	t.readSkip(b)

	if !util.Bool(t.State.UncompressedHeader.SegIdPreSkip) {
		t.intraSegmentId(b)
	}
	t.readCdef(b)
	t.readDeltaQIndex(b)
	t.readDeltaLf(b)

	t.State.ReadDeltas = false
	t.State.RefFrame[0] = shared.INTRA_FRAME
	t.State.RefFrame[0] = shared.NONE

	if t.State.UncompressedHeader.AllowIntraBc {
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
		t.findMvStack(0)
		t.assignMv(0, b)
	} else {
		t.IsInter = 0
		intraFrameYMode := b.S()
		t.YMode = intraFrameYMode
		t.intraAngleInfoY(b)

		if t.HasChroma {
			uvMode := b.S()

			t.UVMode = uvMode

			if t.UVMode == UV_CFL_PRED {
				t.readCflAlphas(b)
			}

			t.intraAngleInfoUv(b)
		}

		t.PaletteSizeY = 0
		t.PaletteSizeUV = 0

		if t.State.MiSize >= shared.BLOCK_8X8 && t.Block_Width[t.State.MiSize] <= 64 && t.Block_Height[t.State.MiSize] <= 64 && util.Bool(t.State.UncompressedHeader.AllowScreenContentTools) {
			t.paletteModeInfo(b)
		}
		t.filterIntraModeInfo(b)
	}
}

// read_segment_id()
func (t *TileGroup) readSegmentId(b *bitstream.BitStream) {
	var prevU int
	var prevL int
	var prevUL int
	var pred int
	if t.State.AvailU && t.State.AvailL {
		prevUL = t.SegmentIds[t.State.MiRow-1][t.State.MiCol-1]
	} else {
		prevUL = -1
	}

	if t.State.AvailU {
		prevU = t.SegmentIds[t.State.MiRow-1][t.State.MiCol]
	} else {
		prevU = -1
	}

	if t.State.AvailL {
		prevL = t.SegmentIds[t.State.MiRow][t.State.MiCol-1]
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
		t.SegmentId = util.NegDeinterleave(t.SegmentId, pred, t.State.UncompressedHeader.LastActiveSegId+1)
	}
}

// read_skip()
func (t *TileGroup) readSkip(b *bitstream.BitStream) {
	if (t.State.UncompressedHeader.SegIdPreSkip == 1) && t.segFeatureActive(shared.SEG_LVL_SKIP) {
		t.Skip = 1
	} else {
		t.Skip = b.S()
	}
}

// seg_feature_active( feature )
func (t *TileGroup) segFeatureActive(feature int) bool {
	return t.segFeatureActiveIdx(t.SegmentId, feature)
}

// seg_feature_active_idx( idx, feature )
func (t *TileGroup) segFeatureActiveIdx(idx int, feature int) bool {
	return t.State.UncompressedHeader.SegmentationEnabled && (t.State.FeatureEnabled[idx][feature] == 1)
}

// intra_segment_id()
func (t *TileGroup) intraSegmentId(b *bitstream.BitStream) {
	if t.State.UncompressedHeader.SegmentationEnabled {
		t.readSegmentId(b)
	} else {
		t.SegmentId = 0
	}

	t.Lossless = t.State.UncompressedHeader.LosslessArray[t.SegmentId]
}

// read_cdef()
func (t *TileGroup) readCdef(b *bitstream.BitStream) {
	if util.Bool(t.Skip) || t.State.UncompressedHeader.CodedLossless || !t.State.SequenceHeader.EnableCdef || t.State.UncompressedHeader.AllowIntraBc {
		return
	}

	cdefSize4 := t.State.Num4x4BlocksWide[shared.BLOCK_64X64]
	cdefMask4 := ^(cdefSize4 - 1)
	r := t.State.MiRow & cdefMask4
	c := t.State.MiCol & cdefMask4

	if t.State.Cdef.CdefIdx[r][c] == -1 {
		t.State.Cdef.CdefIdx[r][c] = b.L(t.State.Cdef.CdefBits)
		w4 := t.State.Num4x4BlocksWide[t.State.MiSize]
		h4 := t.State.Num4x4BlocksHigh[t.State.MiSize]

		for i := r; i < r+h4; i += cdefSize4 {
			for j := c; i < c+w4; i += cdefSize4 {
				t.State.Cdef.CdefIdx[i][j] = t.State.Cdef.CdefIdx[r][c]
			}

		}
	}
}

// read_delta_qindex()
func (t *TileGroup) readDeltaQIndex(b *bitstream.BitStream) {
	var sbSize int
	if t.State.SequenceHeader.Use128x128SuperBlock {
		sbSize = shared.BLOCK_128X128
	} else {
		sbSize = shared.BLOCK_64X64
	}

	if t.State.MiSize == sbSize && util.Bool(t.Skip) {
		return
	}

	if t.State.ReadDeltas {
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

			t.State.CurrentQIndex = util.Clip3(1, 255, t.State.CurrentQIndex+(reducedDeltaQIndex<<t.State.UncompressedHeader.DeltaQRes))

		}
	}
}

// read_delta_lf()
func (t *TileGroup) readDeltaLf(b *bitstream.BitStream) {
	var sbSize int
	if t.State.SequenceHeader.Use128x128SuperBlock {
		sbSize = shared.BLOCK_128X128
	} else {
		sbSize = shared.BLOCK_64X64
	}

	if t.State.MiSize == sbSize && util.Bool(t.Skip) {
		return
	}

	if t.State.ReadDeltas && t.State.UncompressedHeader.DeltaLfPresent {
		frameLfCount := 1

		if util.Bool(t.State.UncompressedHeader.DeltaLfMulti) {
			if t.State.SequenceHeader.ColorConfig.NumPlanes > 1 {
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

				t.State.DeltaLF[i] = util.Clip3(-shared.MAX_LOOP_FILTER, shared.MAX_LOOP_FILTER, t.State.DeltaLF[i]+(reducedDeltaLfLevel<<t.State.UncompressedHeader.DeltaLfRes))
			}
		}
	}
}

// intra_angle_info_y()
func (t *TileGroup) intraAngleInfoY(b *bitstream.BitStream) {
	t.AngleDeltaY = 0

	if t.State.MiSize >= shared.BLOCK_8X8 {

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
func (t *TileGroup) intraAngleInfoUv(b *bitstream.BitStream) {
	t.AngleDeltaUV = 0
	if t.State.MiSize >= shared.BLOCK_8X8 {
		if t.isDirectionalMode(t.UVMode) {
			angleDeltaUv := b.S()
			t.AngleDeltaUV = angleDeltaUv - MAX_ANGLE_DELTA
		}
	}
}

// filter_intra_mode_info()
func (t *TileGroup) filterIntraModeInfo(b *bitstream.BitStream) {
	useFilterIntra := false
	if t.State.SequenceHeader.EnableFilterIntra && t.YMode == DC_PRED && t.PaletteSizeY == 0 && util.Max(t.Block_Width[t.State.MiSize], t.Block_Height[t.State.MiSize]) <= 32 {
		useFilterIntra = util.Bool(b.S())

		if useFilterIntra {
			t.FilterIntraMode = b.S()
		}
	}
}
