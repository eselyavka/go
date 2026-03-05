package solutions

import (
	"strings"
)

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	if strs == nil {
		return ""
	}

	return strings.Join(strs[:], "\x00")
}

func (codec *Codec) Decode(strs string) []string {
	if strs == "" {
		return []string{""}
	}

	return strings.Split(strs, "\x00")
}
