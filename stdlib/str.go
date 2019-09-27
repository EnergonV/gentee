// Copyright 2018 Alexey Krivonogov. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package stdlib

import (
	"fmt"
	"strings"

	"github.com/gentee/gentee/core"
)

// InitStr appends stdlib string functions to the virtual machine
func InitStr(ws *core.Workspace) {
	for _, item := range []interface{}{
		//		core.Link{AddºStrStr,
		//			/*core.Bcode(core.TYPESTR<<16) |*/ core.ADDSTR}, // binary +
		//core.Link{EqualºStrStr, core.EQSTR},                         // binary ==
		//core.Link{GreaterºStrStr, core.GTSTR},                       // binary >
		//core.Link{LenºStr, core.Bcode(core.TYPESTR<<16) | core.LEN}, // the length of str
		//core.Link{LessºStrStr, core.LTSTR},                          // binary <
		//core.Link{intºStr, 2<<16 | core.EMBED},    // int( str )
		//core.Link{floatºStr, 20<<16 | core.EMBED}, // float( str )
		//core.Link{boolºStr, 112<<16 | core.EMBED}, // bool( str )
		//core.Link{ExpStrºStr, core.ADDSTR},                          // expression in string
		//core.Link{AssignºStrStr, core.ASSIGN},              // str = str
		core.Link{AssignAddºStrStr, core.ASSIGN + 1},       // str += str
		core.Link{AssignºStrBool, core.ASSIGN + 2},         // str = bool
		core.Link{AssignºStrInt, core.ASSIGN + 3},          // str = int
		core.Link{FindºStrStr, 113<<16 | core.EMBED},       // Find( str, str ) int
		core.Link{FormatºStr, 39<<16 | core.EMBED},         // Format( str, ... ) str
		core.Link{HasPrefixºStrStr, 86<<16 | core.EMBED},   // HasPrefix( str, str ) bool
		core.Link{HasSuffixºStrStr, 87<<16 | core.EMBED},   // HasSuffix( str, str ) bool
		core.Link{LeftºStrInt, 114<<16 | core.EMBED},       // Left( str, int ) str
		core.Link{LowerºStr, 115<<16 | core.EMBED},         // Lower( str ) str
		core.Link{RepeatºStrInt, 116<<16 | core.EMBED},     // Repeat( str, int )
		core.Link{ReplaceºStrStrStr, 117<<16 | core.EMBED}, // Replace( str, str, str )
		core.Link{ShiftºStr, 118<<16 | core.EMBED},         // unary bitwise OR
		core.Link{SubstrºStrIntInt, 62<<16 | core.EMBED},   // Substr( str, int, int ) str
		core.Link{TrimSpaceºStr, 35<<16 | core.EMBED},      // TrimSpace( str ) str
		core.Link{TrimRightºStr, 107<<16 | core.EMBED},     // TrimRight( str, str ) str
		core.Link{UpperºStr, 108<<16 | core.EMBED},         // Upper( str ) str
	} {
		ws.StdLib().NewEmbed(item)
	}

	for _, item := range []embedInfo{
		{core.Link{LinesºStr, 36<<16 | core.EMBED}, `str`, `arr.str`},        // Lines( str ) arr
		{core.Link{SplitºStrStr, 43<<16 | core.EMBED}, `str,str`, `arr.str`}, // Split( str, str ) arr
	} {
		ws.StdLib().NewEmbedExt(item.Func, item.InTypes, item.OutType)
	}
}

// AssignAddºStrStr appends one string to another
func AssignAddºStrStr(ptr *interface{}, value string) string {
	*ptr = (*ptr).(string) + value
	return (*ptr).(string)
}

// AssignºStrBool assigns boolean to string
func AssignºStrBool(ptr *interface{}, value bool) string {
	*ptr = fmt.Sprint(value)
	return (*ptr).(string)
}

// AssignºStrInt assigns integer to string
func AssignºStrInt(ptr *interface{}, value int64) string {
	*ptr = fmt.Sprint(value)
	return (*ptr).(string)
}

// FindºStrStr returns the index of the first instance of substr
func FindºStrStr(s, substr string) (off int64) {
	off = int64(strings.Index(s, substr))
	if off > 0 {
		off = int64(len([]rune(s[:off])))
	}
	return
}

// FormatºStr formats according to a format specifier and returns the resulting string
func FormatºStr(pattern string, pars ...interface{}) string {
	return fmt.Sprintf(pattern, pars...)
}

// HasPrefixºStrStr returns true if the string s begins with prefix
func HasPrefixºStrStr(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffixºStrStr returns true if the string s ends with suffix
func HasSuffixºStrStr(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// LinesºStr splits a string to a array of strings
func LinesºStr(in string) *core.Array {
	out := core.NewArray()
	items := strings.Split(in, "\n")
	for _, item := range items {
		out.Data = append(out.Data, strings.Trim(item, "\r"))
	}
	return out
}

// LeftºStrInt cuts the string.
func LeftºStrInt(s string, count int64) string {
	r := []rune(s)
	if int(count) > len(r) {
		count = int64(len(r))
	}
	return string(r[:count])
}

// LowerºStr converts a copy of the string to their lower case and returns it.
func LowerºStr(s string) string {
	return strings.ToLower(s)
}

// RepeatºStrInt returns a new string consisting of count copies of the specified string.
func RepeatºStrInt(input string, count int64) string {
	return strings.Repeat(input, int(count))
}

// ReplaceºStrStrStr replaces strings in a string
func ReplaceºStrStrStr(in, old, new string) string {
	return strings.Replace(in, old, new, -1)
}

func replaceArr(in string, old, new []string) string {
	input := []rune(in)
	out := make([]rune, 0, len(input))
	lin := len(input)
	for i := 0; i < lin; i++ {
		eq := -1
		maxLen := lin - i
		for k, item := range old {
			litem := len([]rune(item))
			if maxLen >= litem && string(input[i:i+litem]) == item {
				eq = k
				break
			}
		}
		if eq >= 0 {
			out = append(out, []rune(new[eq])...)
			i += len([]rune(old[eq])) - 1
		} else {
			out = append(out, input[i])
		}
	}
	return string(out)
}

// SplitºStrStr splits a string to a array of strings
func SplitºStrStr(in, sep string) *core.Array {
	out := core.NewArray()
	items := strings.Split(in, sep)
	for _, item := range items {
		out.Data = append(out.Data, item)
	}
	return out
}

// ShiftºStr trims white spaces characters in the each line of the string.
func ShiftºStr(par string) string {
	lines := strings.Split(par, "\n")
	for i, v := range lines {
		lines[i] = strings.TrimSpace(v)
	}
	return strings.Join(lines, "\n")
}

// SubstrºStrIntInt returns a substring with the specified offset and length
func SubstrºStrIntInt(in string, off, length int64) (string, error) {
	var rin []rune
	rin = []rune(in)
	rlen := int64(len(rin))
	if length < 0 {
		length = -length
		off -= length
	}
	if off < 0 || off >= rlen || off+length > rlen {
		return ``, fmt.Errorf(core.ErrorText(core.ErrInvalidParam))
	}
	if length == 0 {
		length = rlen - off
	}
	return string(rin[off : off+length]), nil
}

// TrimSpaceºStr trims white space in a string
func TrimSpaceºStr(in string) string {
	return strings.TrimSpace(in)
}

// TrimRightºStr trims white space in a string
func TrimRightºStr(in string, set string) string {
	return strings.TrimRight(in, set)
}

// UpperºStr converts a copy of the string to their upper case and returns it.
func UpperºStr(s string) string {
	return strings.ToUpper(s)
}
