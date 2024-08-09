// Code generated by "stringer -type=OpCode"; DO NOT EDIT.

package code

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PUSH-0]
	_ = x[ADD-1]
	_ = x[SUB-2]
	_ = x[DIV-3]
	_ = x[MUL-4]
	_ = x[TRUE-5]
	_ = x[FALSE-6]
	_ = x[AND-7]
	_ = x[OR-8]
	_ = x[GT-9]
	_ = x[LT-10]
	_ = x[EQ-11]
	_ = x[NOT-12]
	_ = x[NEQ-13]
	_ = x[JCMP-14]
	_ = x[JMP-15]
	_ = x[JT-16]
	_ = x[SET-17]
	_ = x[GET-18]
}

const _OpCode_name = "PUSHADDSUBDIVMULTRUEFALSEANDORGTLTEQNOTNEQJCMPJMPJTSETGET"

var _OpCode_index = [...]uint8{0, 4, 7, 10, 13, 16, 20, 25, 28, 30, 32, 34, 36, 39, 42, 46, 49, 51, 54, 57}

func (i OpCode) String() string {
	if i >= OpCode(len(_OpCode_index)-1) {
		return "OpCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpCode_name[_OpCode_index[i]:_OpCode_index[i+1]]
}
