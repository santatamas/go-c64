// Code generated by "stringer -type=AddressingMode -linecomment=true"; DO NOT EDIT.

package MOS6510

import "strconv"

const _AddressingMode_name = "ImpliedIndexedIndirectXIndirectIndexedYIndirectAbsoluteAbsoluteXAbsoluteYImmidiateZeroPageZeroPageXZeroPageYAccumulatorRelative"

var _AddressingMode_index = [...]uint8{0, 7, 23, 39, 47, 55, 64, 73, 82, 90, 99, 108, 119, 127}

func (i AddressingMode) String() string {
	i -= 1
	if i < 0 || i >= AddressingMode(len(_AddressingMode_index)-1) {
		return "AddressingMode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AddressingMode_name[_AddressingMode_index[i]:_AddressingMode_index[i+1]]
}
