package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
	"github.com/m4tthewde/gov1/internal/wedgemask"
)

var COMPOUND_MODE_CTX_MAP = [][]int{
	{0, 1, 1, 1, 1},
	{1, 2, 3, 4, 4},
	{4, 4, 5, 6, 7},
}

// inter_block_mode_info()
func (t *TileGroup) interBlockModeInfo(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader, sh sequenceheader.SequenceHeader) {
	t.PaletteSizeY = 0
	t.PaletteSizeUV = 0
	t.readRefFrames(b, state, uh)

	isCompound := state.RefFrame[1] > shared.INTRA_FRAME
	t.findMvStack(util.Int(isCompound), state, uh)

	if util.Bool(t.SkipMode) {
		t.YMode = shared.NEAREST_NEARESTMV
	} else if t.segFeatureActive(shared.SEG_LVL_SKIP, uh, state) || t.segFeatureActive(shared.SEG_LVL_GLOBALMV, uh, state) {
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

	t.assignMv(util.Int(isCompound), b, state, sh, uh)
	t.readInterIntraMode(isCompound, b, state, sh)
	t.readMotionMode(isCompound, b, uh, state)
	t.readCompoundType(isCompound, b, state, sh)

	if uh.InterpolationFilter == shared.SWITCHABLE {
		x := 1
		if sh.EnableDualFilter {
			x = 2
		}
		for dir := 0; dir < x; dir++ {
			if t.needsInterpFilter(state) {
				t.InterpFilter[dir] = b.S()
			} else {
				t.InterpFilter[dir] = shared.EIGHTTAP
			}
		}

		if !sh.EnableDualFilter {
			t.InterpFilter[1] = t.InterpFilter[0]
		}
	} else {
		for dir := 0; dir < 2; dir++ {
			t.InterpFilter[dir] = uh.InterpolationFilter
		}
	}
}

// read_ref_frames()
func (t *TileGroup) readRefFrames(b *bitstream.BitStream, state *state.State, uh uncompressedheader.UncompressedHeader) {
	if util.Bool(t.SkipMode) {
		state.RefFrame[0] = uh.SkipModeFrame[0]
		state.RefFrame[1] = uh.SkipModeFrame[1]
	} else if t.segFeatureActive(shared.SEG_LVL_REF_FRAME, uh, state) {
		state.RefFrame[0] = state.FeatureData[t.SegmentId][shared.SEG_LVL_REF_FRAME]
		state.RefFrame[1] = shared.NONE
	} else {
		bw4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
		bh4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]

		var compMode int
		if uh.ReferenceSelect && util.Min(bw4, bh4) >= 2 {
			compMode = b.S()
		} else {
			compMode = SINGLE_REFERENCE
		}

		if compMode == COMPOUND_REFERENCE {
			compRefType := b.S()
			if compRefType == UNIDIR_COMP_REFERENCE {
				uniCompRef := b.S()
				if util.Bool(uniCompRef) {
					state.RefFrame[0] = shared.BWDREF_FRAME
					state.RefFrame[1] = shared.ALTREF_FRAME
				} else {
					uniCompRefP1 := b.S()
					if util.Bool(uniCompRefP1) {
						uniCompRefP2 := b.S()

						if util.Bool(uniCompRefP2) {
							state.RefFrame[0] = shared.LAST_FRAME
							state.RefFrame[1] = shared.GOLDEN_FRAME
						} else {
							state.RefFrame[0] = shared.LAST_FRAME
							state.RefFrame[1] = shared.LAST3_FRAME
						}
					} else {
						state.RefFrame[0] = shared.LAST_FRAME
						state.RefFrame[1] = shared.LAST2_FRAME

					}
				}
			} else {
				compRef := b.S()
				if compRef == 0 {
					compRefP1 := b.S()

					if util.Bool(compRefP1) {
						state.RefFrame[0] = shared.LAST2_FRAME
					} else {
						state.RefFrame[0] = shared.LAST_FRAME

					}
				} else {
					compRefP2 := b.S()

					if util.Bool(compRefP2) {
						state.RefFrame[0] = shared.GOLDEN_FRAME
					} else {
						state.RefFrame[0] = shared.LAST3_FRAME

					}

				}

				compBwdref := b.S()
				if compBwdref == 0 {
					compBwdrefP1 := b.S()

					if util.Bool(compBwdrefP1) {
						state.RefFrame[1] = shared.ALTREF2_FRAME
					} else {
						state.RefFrame[1] = shared.BWDREF_FRAME

					}
				} else {
					state.RefFrame[1] = shared.ALTREF_FRAME
				}
			}
		} else {
			singleRefP1 := t.singleRefP1Symbol(state, b, uh)
			if util.Bool(singleRefP1) {
				singleRefP2 := b.S()
				if singleRefP2 == 0 {
					singleRefP6 := b.S()
					if util.Bool(singleRefP6) {
						state.RefFrame[0] = shared.ALTREF2_FRAME
					} else {
						state.RefFrame[0] = shared.BWDREF_FRAME

					}
				} else {
					state.RefFrame[0] = shared.ALTREF_FRAME
				}
			} else {
				singleRefP3 := t.singleRefP3Symbol(state, b, uh)
				if util.Bool(singleRefP3) {
					singleRefP5 := b.S()
					if util.Bool(singleRefP5) {
						state.RefFrame[0] = shared.GOLDEN_FRAME
					} else {
						state.RefFrame[0] = shared.LAST3_FRAME
					}
				} else {
					singleRefP4 := t.singleRefP4Symbol(state, b, uh)
					if util.Bool(singleRefP4) {
						state.RefFrame[0] = shared.LAST2_FRAME
					} else {
						state.RefFrame[0] = shared.LAST_FRAME
					}
				}
			}
			state.RefFrame[1] = shared.NONE
		}
	}
}

// read_motion_mode( isCompound )
func (t *TileGroup) readMotionMode(isCompound bool, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader, state *state.State) {
	if util.Bool(t.SkipMode) {
		t.MotionMode = SIMPLE
		return
	}

	if !uh.IsMotionModeSwitchable {
		t.MotionMode = SIMPLE
		return
	}

	if util.Min(t.Block_Width[state.MiSize], t.Block_Height[state.MiSize]) < 8 {
		t.MotionMode = SIMPLE
		return
	}

	if !uh.ForceIntegerMv && (t.YMode == shared.GLOBALMV || t.YMode == shared.GLOBAL_GLOBALMV) {
		if state.GmType[state.RefFrame[0]] > shared.TRANSLATION {
			t.MotionMode = SIMPLE
			return
		}
	}

	t.findWarpSamples(state)
	if uh.ForceIntegerMv || t.NumSamples == 0 || !uh.AllowWarpedMotion || t.isScaled(state.RefFrame[0], uh) {
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
func (t *TileGroup) findWarpSamples(state *state.State) {
	t.NumSamples = 0
	t.NumSamplesScanned = 0

	w4 := shared.NUM_4X4_BLOCKS_WIDE[state.MiSize]
	h4 := shared.NUM_4X4_BLOCKS_HIGH[state.MiSize]

	doTopLeft := 1
	doTopRight := 1

	if state.AvailU {
		srcSize := state.MiSizes[state.MiRow-1][state.MiCol]
		srcW := shared.NUM_4X4_BLOCKS_WIDE[srcSize]

		if w4 <= srcW {
			colOffset := -(state.MiCol & (srcW - 1))
			if colOffset < 0 {
				doTopLeft = 0
			}
			if colOffset+srcW > w4 {
				doTopRight = 0
			}
			t.addSample(-1, 0, state)
		} else {
			var miStep int
			for i := 0; i < util.Min(w4, state.MiCols-state.MiCol); i += miStep {
				srcSize = state.MiSizes[state.MiRow-1][state.MiCol+i]
				srcW = shared.NUM_4X4_BLOCKS_WIDE[srcSize]
				miStep = util.Min(w4, srcW)
				t.addSample(-1, i, state)
			}
		}
	}
	if state.AvailL {
		srcSize := state.MiSizes[state.MiRow][state.MiCol-1]
		srcH := shared.NUM_4X4_BLOCKS_HIGH[srcSize]

		if h4 <= srcH {
			rowOffset := -(state.MiRow & (srcH - 1))
			if rowOffset < 0 {
				doTopLeft = 0
			}
			t.addSample(0, -1, state)
		} else {
			var miStep int
			for i := 0; i < util.Min(h4, state.MiRows-state.MiRow); i += miStep {
				srcSize = state.MiSizes[state.MiRow+i][state.MiCol-1]
				srcH = shared.NUM_4X4_BLOCKS_HIGH[srcSize]
				miStep = util.Min(h4, srcH)
				t.addSample(i, -1, state)
			}
		}
	}

	if util.Bool(doTopLeft) {
		t.addSample(-1, -1, state)
	}

	if util.Bool(doTopRight) {
		if util.Max(w4, h4) <= 16 {
			t.addSample(-1, w4, state)
		}
	}

	if t.NumSamples == 0 && t.NumSamplesScanned > 0 {
		t.NumSamples = 1
	}

}

// add_sample 7.10.4.2.
func (t *TileGroup) addSample(deltaRow int, deltaCol int, state *state.State) {
	if t.NumSamplesScanned >= LEAST_SQUARES_SAMPLES_MAX {
		return
	}

	mvRow := state.MiRow + deltaRow
	mvCol := state.MiCol + deltaCol

	if !t.isInside(mvRow, mvCol, state) {
		return
	}

	// TODO: how do we know if something has not been writte to?
	if state.RefFrames[mvRow][mvCol][0] == 0 {
		return
	}

	if state.RefFrames[mvRow][mvCol][0] != state.RefFrame[0] {
		return
	}

	if state.RefFrames[mvRow][mvCol][1] != shared.NONE {
		return
	}

	candSz := state.MiSizes[mvRow][mvCol]
	candW4 := shared.NUM_4X4_BLOCKS_WIDE[candSz]
	candH4 := shared.NUM_4X4_BLOCKS_HIGH[candSz]
	candRow := mvRow & ^(candH4 - 1)
	candCol := mvCol & ^(candW4 - 1)
	midY := candRow*4 + candH4*2 - 1
	midX := candCol*4 + candW4*2 - 1
	threshold := util.Clip3(16, 112, util.Max(t.Block_Width[state.MiSize], t.Block_Height[state.MiSize]))
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
func (t *TileGroup) needsInterpFilter(state *state.State) bool {
	large := util.Min(t.Block_Width[state.MiSize], t.Block_Height[state.MiSize]) >= 8

	if util.Bool(t.SkipMode) || t.MotionMode == LOCALWARP {
		return false
	} else if large && t.YMode == shared.GLOBALMV {
		return state.GmType[state.RefFrame[0]] == shared.TRANSLATION
	} else if large && t.YMode == shared.GLOBAL_GLOBALMV {
		return state.GmType[state.RefFrame[0]] == shared.TRANSLATION || state.GmType[1] == shared.TRANSLATION
	} else {
		return true
	}
}

// read_compound_type( isCompound )
func (t *TileGroup) readCompoundType(isCompound bool, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader) {
	t.CompGroupIdx = 0
	t.CompoundIdx = 1
	if util.Bool(t.SkipMode) {
		t.CompoundType = COMPOUND_AVERAGE
		return
	}

	if isCompound {
		n := wedgemask.Wedge_Bits[state.MiSize]
		if sh.EnableMaskedCompound {
			t.CompGroupIdx = b.S()
		}

		if t.CompGroupIdx == 0 {
			if sh.EnableJntComp {
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
func (t *TileGroup) readInterIntraMode(isCompound bool, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader) {
	if util.Bool(t.SkipMode) && sh.EnableInterIntraCompound && !isCompound && state.MiSize > +shared.BLOCK_8X8 && state.MiSize <= shared.BLOCK_32X32 {
		t.InterIntra = b.S()

		if util.Bool(t.InterIntra) {
			t.InterIntraMode = b.S()
			state.RefFrame[1] = shared.INTRA_FRAME
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
