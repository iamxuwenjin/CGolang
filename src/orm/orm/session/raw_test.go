package session

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrBuilder(t *testing.T) {
	ss := []string{
		"sh",
		"hn",
		"test",
	}

	var b strings.Builder
	for _, s := range ss {
		_, _ = fmt.Fprint(&b, s)
	}

	print(b.String())
}
