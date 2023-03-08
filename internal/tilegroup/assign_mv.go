package tilegroup

import (
	"github.com/m4tthewde/gov1/internal/bitstream"
	"github.com/m4tthewde/gov1/internal/sequenceheader"
	"github.com/m4tthewde/gov1/internal/shared"
	"github.com/m4tthewde/gov1/internal/state"
	"github.com/m4tthewde/gov1/internal/uncompressedheader"
	"github.com/m4tthewde/gov1/internal/util"
)

// assign_mv( isCompound )
func (t *TileGroup) assignMv(isCompound int, b *bitstream.BitStream, state *state.State, sh sequenceheader.SequenceHeader, uh uncompressedheader.UncompressedHeader) {
	for i := 0; i < 1+isCompound; i++ {
		var compMode int
		if util.Bool(t.useIntrabc) {
			compMode = shared.NEWMV
		} else {
			compMode = t.getMode(i)
		}

		if util.Bool(t.useIntrabc) {
			t.PredMv[0] = t.RefStackMv[0][0]
			if t.PredMv[0][0] == 0 && t.PredMv[0][1] == 0 {
				t.PredMv[0] = t.RefStackMv[1][0]
			}
			if t.PredMv[0][0] == 0 && t.PredMv[0][1] == 0 {
				var sbSize int
				if sh.Use128x128SuperBlock {
					sbSize = shared.BLOCK_128X128
				} else {
					sbSize = shared.BLOCK_64X64
				}
				sbSize4 := state.Num4x4BlocksHigh[sbSize]

				if state.MiRow-sbSize4 < state.MiRowStart {
					t.PredMv[0][0] = 0
					t.PredMv[0][1] = -(sbSize4*MI_SIZE + INTRABC_DELAY_PIXELS) * 8
				} else {
					t.PredMv[0][0] = -(sbSize4 * MI_SIZE * 8)
					t.PredMv[0][0] = 1
				}
			}

		} else if compMode == shared.GLOBALMV {
			t.PredMv[i] = t.GlobalMvs[i]
		} else {
			var pos int
			if compMode == shared.NEARESTMV {
				pos = 0
			} else {
				pos = t.RefMvIdx
			}

			if compMode == shared.NEWMV && t.NumMvFound <= 1 {
				pos = 0
			}

			t.PredMv[i] = t.RefStackMv[pos][i]
		}

		if compMode == shared.NEWMV {
			t.readMv(i, b, uh)
		} else {
			t.Mv[i] = t.PredMv[i]
		}
	}
}

// read_mv( ref )
func (t *TileGroup) readMv(ref int, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) {
	var diffMv []int
	diffMv[0] = 0
	diffMv[1] = 0

	if util.Bool(t.useIntrabc) {
		t.MvCtx = MV_INTRABC_CONTEXT
	} else {
		t.MvCtx = 0
	}

	mvJoint := b.S()

	if mvJoint == MV_JOINT_HZVNZ || mvJoint == MV_JOINT_HNZVNZ {
		diffMv[0] = t.readMvComponent(0, b, uh)
	}

	if mvJoint == MV_JOINT_HNZVZ || mvJoint == MV_JOINT_HNZVNZ {
		diffMv[1] = t.readMvComponent(1, b, uh)
	}

	t.Mv[ref][0] = t.PredMv[ref][0] + diffMv[0]
	t.Mv[ref][1] = t.PredMv[ref][1] + diffMv[1]
}

// read_mv_component( comp )
func (t *TileGroup) readMvComponent(comp int, b *bitstream.BitStream, uh uncompressedheader.UncompressedHeader) int {
	mvSign := b.S()
	mvClass := b.S()

	var mag int
	if mvClass == MV_CLASS_0 {
		mvClass0Bit := b.S()

		var mvClass0Fr int
		if uh.ForceIntegerMv {
			mvClass0Fr = 3
		} else {
			mvClass0Fr = b.S()
		}

		var mvClass0Hp int
		if uh.AllowHighPrecisionMv {
			mvClass0Hp = b.S()
		} else {
			mvClass0Hp = 1
		}

		mag = ((mvClass0Bit << 3) | (mvClass0Fr << 1) | mvClass0Hp) + 1
	} else {
		d := 0
		for i := 0; i < mvClass; i++ {
			mvBit := b.S()
			d |= mvBit << 1
		}

		mag = CLASS0_SIZE << (mvClass + 2)

		var mvFr int
		var mvHp int
		if uh.ForceIntegerMv {
			mvFr = 3
		} else {
			mvFr = b.S()
		}

		if uh.AllowHighPrecisionMv {
			mvHp = b.S()
		} else {
			mvHp = 1
		}

		mag += ((d << 3) | (mvFr << 1) | mvHp) + 1
	}

	if util.Bool(mvSign) {
		return -mag
	} else {
		return mag
	}
}

// get_mode( refList )
func (t *TileGroup) getMode(refList int) int {
	var compMode int
	if refList == 0 {
		if t.YMode < shared.NEAREST_NEARESTMV {
			compMode = t.YMode
		} else if t.YMode == shared.NEW_NEWMV || t.YMode == shared.NEW_NEARESTMV || t.YMode == shared.NEW_NEARMV {
			compMode = shared.NEWMV
		} else if t.YMode == shared.NEAREST_NEARESTMV || t.YMode == shared.NEAREST_NEWMV {
			compMode = shared.NEARESTMV
		} else if t.YMode == shared.NEAR_NEARMV || t.YMode == shared.NEAR_NEWMV {
			compMode = shared.NEARMV
		} else {
			compMode = shared.GLOBALMV
		}
	} else {
		if t.YMode == shared.NEW_NEWMV || t.YMode == shared.NEAREST_NEWMV || t.YMode == shared.NEAR_NEWMV {
			compMode = shared.NEWMV
		} else if t.YMode == shared.NEAREST_NEARESTMV || t.YMode == shared.NEW_NEARESTMV {
			compMode = shared.NEARMV
		} else if t.YMode == shared.NEAR_NEARMV || t.YMode == shared.NEW_NEARMV {
			compMode = shared.NEARMV
		} else {
			compMode = shared.GLOBALMV
		}
	}

	return compMode
}
