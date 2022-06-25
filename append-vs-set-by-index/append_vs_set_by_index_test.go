package append_vs_set_by_index

import (
	"strconv"
	"testing"
)

type User struct {
	ID string
}

const count = 4096

func CopyUsingAppend(src []*User) []string {
	dst := make([]string, 0, count)
	for _, user := range src {
		dst = append(dst, user.ID)
	}
	return dst
}

func CopyByIndex(src []*User) []string {
	dst := make([]string, count)
	for i, user := range src {
		dst[i] = user.ID
	}
	return dst
}

func BenchmarkCopyByIndex(b *testing.B) {
	src := make([]*User, count)
	for i := 0; i < count; i++ {
		src[i] = &User{strconv.Itoa(i)}
	}
	for i := 0; i < b.N; i++ {
		CopyByIndex(src)
	}
}

func BenchmarkCopyUsingAppend(b *testing.B) {
	src := make([]*User, count)
	for i := 0; i < count; i++ {
		src[i] = &User{strconv.Itoa(i)}
	}
	for i := 0; i < b.N; i++ {
		CopyUsingAppend(src)
	}
}
