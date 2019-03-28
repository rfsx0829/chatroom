package graphic

import "fmt"

func CenterStr(str string, length int) string {
	l := len(str)
	if l&1 != length&1 {
		str += " "
		l++
	}
	return fmt.Sprintf("%*s%s%*s", (length-l)/2, "", str, (length-l)/2, "")
}

func StrsLine(div string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	str := div
	for i := 0; i < len(strs); i++ {
		str += strs[i] + div
	}
	return str
}

func StrRepeat(single string, times int) string {
	str := ""
	for times > 0 {
		times--
		str += single
	}
	return str
}
