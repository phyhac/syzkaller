package thuffle

import (
	"fmt"
)

const (
	MaxSizeSchedule = 32
)

type SchedPoint struct {
	CallID   int    // in which call the point will be trigered
	IrqLine  int    // the interrupt handler trigged
	CodeAddr uint64 // trigged when prog PC in codeAddr
	MemAddr  uint64 // trigged when mem in memAddr being accessed
	Order    uint64 // hint order
}

type IrqSchedule struct {
	points []SchedPoint
	masks  []uint8 // only when masks[i] == 1, points[i] is treated
	size   int
}

func (s *IrqSchedule) Match(callID int) []*SchedPoint {
	res := []*SchedPoint{}
	for i := 0; i < s.size; i++ {
		if s.masks[i] == 1 && s.points[i].CallID == callID {
			res = append(res, &s.points[i])
		}
	}
	return res
}

func (s *IrqSchedule) Mask(i int) error {
	if i < 0 || i >= s.size {
		return fmt.Errorf("index out of bound: %v (supported: [0, %v)", i, s.size)
	}
	return nil
}

func (s *IrqSchedule) AfterMask() []*SchedPoint {
	res := []*SchedPoint{}
	for i, m := range s.masks {
		if m == 1 {
			res = append(res, &s.points[i])
		}
	}
	return res
}
