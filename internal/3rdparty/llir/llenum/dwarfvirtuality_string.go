// Code generated by "stringer -linecomment -type DwarfVirtuality"; DO NOT EDIT.

package llenum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DwarfVirtualityNone-0]
	_ = x[DwarfVirtualityVirtual-1]
	_ = x[DwarfVirtualityPureVirtual-2]
}

const _DwarfVirtuality_name = "DW_VIRTUALITY_noneDW_VIRTUALITY_virtualDW_VIRTUALITY_pure_virtual"

var _DwarfVirtuality_index = [...]uint8{0, 18, 39, 65}

func (i DwarfVirtuality) String() string {
	if i < 0 || i >= DwarfVirtuality(len(_DwarfVirtuality_index)-1) {
		return "DwarfVirtuality(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DwarfVirtuality_name[_DwarfVirtuality_index[i]:_DwarfVirtuality_index[i+1]]
}
