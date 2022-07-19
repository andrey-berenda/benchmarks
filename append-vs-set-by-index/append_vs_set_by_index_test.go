package append_vs_set_by_index

import (
	"testing"

	"github.com/google/uuid"
)

type ID string

const count = 1 << 12

func NewUser(i int) *User {
	return &User{
		ID: ID(uuid.NewString()),
	}
}

type User struct {
	ID ID
}

func CopyUsingAppend(src []*User) []ID {
	dst := make([]ID, 0, len(src))
	for _, user := range src {
		dst = append(dst, user.ID)
	}
	return dst
}

func CopyByIndex(src []*User) []ID {
	dst := make([]ID, len(src))
	for i, user := range src {
		dst[i] = user.ID
	}
	return dst
}

func CopyUsingAppendWithoutAlloc(src []*User) []ID {
	var dst []ID
	for _, user := range src {
		dst = append(dst, user.ID)
	}
	return dst
}

func BenchmarkCopyUsingAppendWithoutAlloc(b *testing.B) {
	src := make([]*User, count)
	for i := 0; i < count; i++ {
		src[i] = NewUser(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	var result []ID
	if result == nil {
		for i := 0; i < b.N; i++ {
			CopyUsingAppendWithoutAlloc(src)
		}
	}
}

func BenchmarkCopyUsingAppend(b *testing.B) {
	src := make([]*User, count)
	for i := 0; i < count; i++ {
		src[i] = NewUser(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyUsingAppend(src)
	}
}

func BenchmarkCopyByIndex(b *testing.B) {
	src := make([]*User, count)
	for i := 0; i < count; i++ {
		src[i] = NewUser(i)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyByIndex(src)
	}
}
