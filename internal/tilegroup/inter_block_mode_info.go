package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/util"
	"github.com/m4tthewde/gov1/internal/wedgemask"
)

// inter_block_mode_info()
func (t *TileGroup) interBlockModeInfo(b *bitstream.BitStream) {
	t.PaletteSizeY = 0
	t.PaletteSizeUV = 0
	t.readRefFrames(b)

	isCompound := t.State.RefFrame[1] > INTRA_FRAME
	t.findMvStack(util.Int(isCompound))

	if util.Bool(t.SkipMode) {
		t.YMode = shared.NEAREST_NEARESTMV
	} else if t.segFeatureActive(shared.SEG_LVL_SKIP) || t.segFeatureActive(shared.SEG_LVL_GLOBALMV) {
		t.YMode = shared.GLOBALMV
	} else if isCompound {
		compoundMode := b.S()
		t.YMode = shared.NEAREST_NEARESTMV + compoundMode
	} else {
		newMv := b.S()
		if newMv == 0 {
			t.YMode = shared.NEWMV
		} else {
			zeroMv := b.S()
			if zeroMv == 0 {
				t.YMode = shared.GLOBALMV
			} else {
				refMv := b.S()
				if refMv == 0 {
					t.YMode = shared.NEARESTMV
				} else {
					t.YMode = shared.NEARMV
				}
			}
		}
	}

	t.RefMvIdx = 0
	if t.YMode == shared.NEWMV || t.YMode == shared.NEW_NEWMV {
		for idx := 0; idx < 2; idx++ {
			if t.NumMvFound > idx+1 {
				drlMode := b.S()
				if drlMode == 0 {
					t.RefMvIdx = idx
					break
				}
				t.RefMvIdx = idx + 1
			}
		}
	} else if t.hasNearmv() {
		t.RefMvIdx = 1
		for idx := 1; idx < 3; idx++ {
			if t.NumMvFound > idx+1 {
				drlMode := b.S()
				if drlMode == 0 {
					t.RefMvIdx = idx
					break
				}
				t.RefMvIdx = idx + 1
			}
		}
	}

	t.assignMv(util.Int(isCompound), b)
	t.readInterIntraMode(isCompound, b)
	t.readMotionMode(isCompound, b)
	t.readCompoundType(isCompound, b)

	if t.State.UncompressedHeader.InterpolationFilter == shared.SWITCHABLE {
		x := 1
		if t.State.SequenceHeader.EnableDualFilter {
			x = 2
		}
		for dir := 0; dir < x; dir++ {
			if t.needsInterpFilter() {
				t.InterpFilter[dir] = b.S()
			} else {
				t.InterpFilter[dir] = shared.EIGHTTAP
			}
		}

		if !t.State.SequenceHeader.EnableDualFilter {
			t.InterpFilter[1] = t.InterpFilter[0]
		}
	} else {
		for dir := 0; dir < 2; dir++ {
			t.InterpFilter[dir] = t.State.UncompressedHeader.InterpolationFilter
		}
	}
}

// read_ref_frames()
func (t *TileGroup) readRefFrames(b *bitstream.BitStream) {
	if util.Bool(t.SkipMode) {
		t.State.RefFrame[0] = t.State.UncompressedHeader.SkipModeFrame[0]
		t.State.RefFrame[1] = t.State.UncompressedHeader.SkipModeFrame[1]
	} else if t.segFeatureActive(shared.SEG_LVL_REF_FRAME) {
		t.State.RefFrame[0] = t.State.FeatureData[t.SegmentId][shared.SEG_LVL_REF_FRAME]
		t.State.RefFrame[1] = NONE
	} else {
		bw4 := t.State.Num4x4BlocksWide[t.State.MiSize]
		bh4 := t.State.Num4x4BlocksHigh[t.State.MiSize]

		var compMode int
		if t.State.UncompressedHeader.ReferenceSelect && util.Min(bw4, bh4) >= 2 {
			compMode = b.S()
		} else {
			compMode = SINGLE_REFERENCE
		}

		if compMode == COMPOUND_REFERENCE {
			compRefType := b.S()
			if compRefType == UNIDIR_COMP_REFERENCE {
				uniCompRef := b.S()
				if util.Bool(uniCompRef) {
					t.State.RefFrame[0] = shared.BWDREF_FRAME
					t.State.RefFrame[1] = shared.ALTREF_FRAME
				} else {
					uniCompRefP1 := b.S()
					if util.Bool(uniCompRefP1) {
						uniCompRefP2 := b.S()

						if util.Bool(uniCompRefP2) {
							t.State.RefFrame[0] = shared.LAST_FRAME
							t.State.RefFrame[1] = shared.GOLDEN_FRAME
						} else {
							t.State.RefFrame[0] = shared.LAST_FRAME
							t.State.RefFrame[1] = shared.LAST3_FRAME
						}
					} else {
						t.State.RefFrame[0] = shared.LAST_FRAME
						t.State.RefFrame[1] = shared.LAST2_FRAME

					}
				}
			} else {
				compRef := b.S()
				if compRef == 0 {
					compRefP1 := b.S()

					if util.Bool(compRefP1) {
						t.State.RefFrame[0] = shared.LAST2_FRAME
					} else {
						t.State.RefFrame[0] = shared.LAST_FRAME

					}
				} else {
					compRefP2 := b.S()

					if util.Bool(compRefP2) {
						t.State.RefFrame[0] = shared.GOLDEN_FRAME
					} else {
						t.State.RefFrame[0] = shared.LAST3_FRAME

					}

				}

				compBwdref := b.S()
				if compBwdref == 0 {
					compBwdrefP1 := b.S()

					if util.Bool(compBwdrefP1) {
						t.State.RefFrame[1] = shared.ALTREF2_FRAME
					} else {
						t.State.RefFrame[1] = shared.BWDREF_FRAME

					}
				} else {
					t.State.RefFrame[1] = shared.ALTREF_FRAME
				}
			}
		} else {
			singleRefP1 := b.S()
			if util.Bool(singleRefP1) {
				singleRefP2 := b.S()
				if singleRefP2 == 0 {
					singleRefP6 := b.S()
					if util.Bool(singleRefP6) {
						t.State.RefFrame[0] = shared.ALTREF2_FRAME
					} else {
						t.State.RefFrame[0] = shared.BWDREF_FRAME

					}
				} else {
					t.State.RefFrame[0] = shared.ALTREF_FRAME
				}
			} else {
				singleRefP3 := b.S()
				if util.Bool(singleRefP3) {
					singleRefP5 := b.S()
					if util.Bool(singleRefP5) {
						t.State.RefFrame[0] = shared.GOLDEN_FRAME
					} else {
						t.State.RefFrame[0] = shared.LAST3_FRAME
					}
				} else {
					singleRefP4 := b.S()
					if util.Bool(singleRefP4) {
						t.State.RefFrame[0] = shared.LAST2_FRAME
					} else {
						t.State.RefFrame[0] = shared.LAST_FRAME
					}
				}
			}
			t.State.RefFrame[1] = NONE
		}
	}
}

// read_motion_mode( isCompound )
func (t *TileGroup) readMotionMode(isCompound bool, b *bitstream.BitStream) {
	if util.Bool(t.SkipMode) {
		t.MotionMode = SIMPLE
		return
	}

	if !t.State.UncompressedHeader.IsMotionModeSwitchable {
		t.MotionMode = SIMPLE
		return
	}

	if util.Min(t.Block_Width[t.State.MiSize], t.Block_Height[t.State.MiSize]) < 8 {
		t.MotionMode = SIMPLE
		return
	}

	if !t.State.UncompressedHeader.ForceIntegerMv && (t.YMode == shared.GLOBALMV || t.YMode == shared.GLOBAL_GLOBALMV) {
		if t.State.GmType[t.State.RefFrame[0]] > shared.TRANSLATION {
			t.MotionMode = SIMPLE
			return
		}
	}

	t.findWarpSamples()
	if t.State.UncompressedHeader.ForceIntegerMv || t.NumSamples == 0 || !t.State.UncompressedHeader.AllowWarpedMotion || t.isScaled(t.State.RefFrame[0]) {
		useObmc := b.S()
		if util.Bool(useObmc) {
			t.MotionMode = OBMC

		} else {
			t.MotionMode = SIMPLE
		}
	} else {
		t.MotionMode = b.S()
	}
}

// TODO: consider separate file for this procedure
// find_warp_samples() 7.10.4.
func (t *TileGroup) findWarpSamples() {
	t.NumSamples = 0
	t.NumSamplesScanned = 0

	w4 := t.State.Num4x4BlocksWide[t.State.MiSize]
	h4 := t.State.Num4x4BlocksHigh[t.State.MiSize]

	doTopLeft := 1
	doTopRight := 1

	if t.State.AvailU {
		srcSize := t.State.MiSizes[t.State.MiRow-1][t.State.MiCol]
		srcW := t.State.Num4x4BlocksWide[srcSize]

		if w4 <= srcW {
			colOffset := -(t.State.MiCol & (srcW - 1))
			if colOffset < 0 {
				doTopLeft = 0
			}
			if colOffset+srcW > w4 {
				doTopRight = 0
			}
			t.addSample(-1, 0)
		} else {
			var miStep int
			for i := 0; i < util.Min(w4, t.State.MiCols-t.State.MiCol); i += miStep {
				srcSize = t.State.MiSizes[t.State.MiRow-1][t.State.MiCol+i]
				srcW = t.State.Num4x4BlocksWide[srcSize]
				miStep = util.Min(w4, srcW)
				t.addSample(-1, i)
			}
		}
	}
	if t.State.AvailL {
		srcSize := t.State.MiSizes[t.State.MiRow][t.State.MiCol-1]
		srcH := t.State.Num4x4BlocksHigh[srcSize]

		if h4 <= srcH {
			rowOffset := -(t.State.MiRow & (srcH - 1))
			if rowOffset < 0 {
				doTopLeft = 0
			}
			t.addSample(0, -1)
		} else {
			var miStep int
			for i := 0; i < util.Min(h4, t.State.MiRows-t.State.MiRow); i += miStep {
				srcSize = t.State.MiSizes[t.State.MiRow+i][t.State.MiCol-1]
				srcH = t.State.Num4x4BlocksHigh[srcSize]
				miStep = util.Min(h4, srcH)
				t.addSample(i, -1)
			}
		}
	}

	if util.Bool(doTopLeft) {
		t.addSample(-1, -1)
	}

	if util.Bool(doTopRight) {
		if util.Max(w4, h4) <= 16 {
			t.addSample(-1, w4)
		}
	}

	if t.NumSamples == 0 && t.NumSamplesScanned > 0 {
		t.NumSamples = 1
	}

}

// add_sample 7.10.4.2.
func (t *TileGroup) addSample(deltaRow int, deltaCol int) {
	if t.NumSamplesScanned >= LEAST_SQUARES_SAMPLES_MAX {
		return
	}

	mvRow := t.State.MiRow + deltaRow
	mvCol := t.State.MiCol + deltaCol

	if !t.isInside(mvRow, mvCol) {
		return
	}

	// TODO: how do we know if something has not been writte to?
	if t.State.RefFrames[mvRow][mvCol][0] == 0 {
		return
	}

	if t.State.RefFrames[mvRow][mvCol][0] != t.State.RefFrame[0] {
		return
	}

	if t.State.RefFrames[mvRow][mvCol][1] != NONE {
		return
	}

	candSz := t.State.MiSizes[mvRow][mvCol]
	candW4 := t.State.Num4x4BlocksWide[candSz]
	candH4 := t.State.Num4x4BlocksHigh[candSz]
	candRow := mvRow & ^(candH4 - 1)
	candCol := mvCol & ^(candW4 - 1)
	midY := candRow*4 + candH4*2 - 1
	midX := candCol*4 + candW4*2 - 1
	threshold := util.Clip3(16, 112, util.Max(t.Block_Width[t.State.MiSize], t.Block_Height[t.State.MiSize]))
	mvDiffRow := util.Abs(t.Mvs[candRow][candCol][0][0] - t.Mv[0][0])
	mvDiffCol := util.Abs(t.Mvs[candRow][candCol][0][1] - t.Mv[0][1])
	valid := (mvDiffRow + mvDiffCol) <= threshold

	var cand []int
	cand[0] = midY * 8
	cand[1] = midX * 8
	cand[2] = midY*8 + t.Mvs[candRow][candCol][0][0]
	cand[3] = midX*8 + t.Mvs[candRow][candCol][0][1]

	t.NumSamplesScanned++
	if valid && t.NumSamplesScanned > 1 {
		return
	}

	for j := 0; j < 4; j++ {
		t.CandList[t.NumSamples][j] = cand[j]
	}

	if valid {
		t.NumSamples++
	}
}

// needs_interp_filter()
func (t *TileGroup) needsInterpFilter() bool {
	large := util.Min(t.Block_Width[t.State.MiSize], t.Block_Height[t.State.MiSize]) >= 8

	if util.Bool(t.SkipMode) || t.MotionMode == LOCALWARP {
		return false
	} else if large && t.YMode == shared.GLOBALMV {
		return t.State.GmType[t.State.RefFrame[0]] == shared.TRANSLATION
	} else if large && t.YMode == shared.GLOBAL_GLOBALMV {
		return t.State.GmType[t.State.RefFrame[0]] == shared.TRANSLATION || t.State.GmType[1] == shared.TRANSLATION
	} else {
		return true
	}
}

// read_compound_type( isCompound )
func (t *TileGroup) readCompoundType(isCompound bool, b *bitstream.BitStream) {
	t.CompGroupIdx = 0
	t.CompoundIdx = 1
	if util.Bool(t.SkipMode) {
		t.CompoundType = COMPOUND_AVERAGE
		return
	}

	if isCompound {
		n := wedgemask.Wedge_Bits[t.State.MiSize]
		if t.State.SequenceHeader.EnableMaskedCompound {
			t.CompGroupIdx = b.S()
		}

		if t.CompGroupIdx == 0 {
			if t.State.SequenceHeader.EnableJntComp {
				t.CompoundIdx = b.S()
				if util.Bool(t.CompoundIdx) {
					t.CompoundType = COMPOUND_AVERAGE

				} else {
					t.CompoundType = COMPOUND_DISTANCE
				}
			} else {
				t.CompoundType = COMPOUND_AVERAGE
			}
		} else {
			if n == 0 {
				t.CompoundType = COMPOUND_DIFFWTD
			} else {
				t.CompoundType = b.S()
			}
		}

		if t.CompoundType == COMPOUND_WEDGE {
			t.WedgeIndex = b.S()
			t.WedgeIndex = b.L(1)
		} else if t.CompoundType == COMPOUND_DIFFWTD {
			t.MaskType = b.L(1)
		}
	} else {
		if util.Bool(t.InterIntra) {
			if util.Bool(t.WedgeInterIntra) {
				t.CompoundType = COMPOUND_WEDGE
			} else {
				t.CompoundType = COMPOUND_INTRA
			}
		} else {
			t.CompoundType = COMPOUND_AVERAGE
		}
	}

}

// read_interintra_mode( isCompound )
func (t *TileGroup) readInterIntraMode(isCompound bool, b *bitstream.BitStream) {
	if util.Bool(t.SkipMode) && t.State.SequenceHeader.EnableInterIntraCompound && !isCompound && t.State.MiSize > +shared.BLOCK_8X8 && t.State.MiSize <= shared.BLOCK_32X32 {
		t.InterIntra = b.S()

		if util.Bool(t.InterIntra) {
			t.InterIntraMode = b.S()
			t.State.RefFrame[1] = INTRA_FRAME
			t.AngleDeltaY = 0
			t.AngleDeltaUV = 0
			t.UseFilterIntra = 0
			t.WedgeInterIntra = b.S()
			if util.Bool(t.WedgeInterIntra) {
				t.WedgeIndex = b.S()
				t.WedgeSign = 0
			}
		}
	} else {
		t.InterIntra = 0
	}
}

// has_nearmv()
func (t *TileGroup) hasNearmv() bool {
	return t.YMode == shared.NEARMV || t.YMode == shared.NEAR_NEARMV || t.YMode == shared.NEAR_NEWMV || t.YMode == shared.NEW_NEARMV
}
