package util

import "strconv"

func FormatUserIDStr(userID uint64) string {
	return strconv.FormatUint(userID, 10)
}
