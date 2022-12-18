package stringtable

import (
	"fmt"
	"testing"
)

func TestSanity(t *testing.T) {
	st := StringTable{}
	m := make(map[string]int)
	m2 := make(map[int]string)
	sliceSize := 0
	for i := 0; i < 1000; i++ {
		str := fmt.Sprintf("str%d", i)
		index := st.Put(str)
		m[str] = index
		m2[index] = str
		sliceSize++
	}
	if st.Len() != len(m) {
		t.Errorf("Wrong len: %d", st.Len())
	}
	// Can read whats written
	for k, v := range m {
		if st.Get(v) != k {
			t.Errorf("Got %s for %d but expected %s", st.Get(v), v, k)
		}
	}
	// Add another gives same index
	for i := 0; i < 1000; i++ {
		str := fmt.Sprintf("str%d", i)
		index := st.Put(str)
		if index != m[str] {
			t.Errorf("Second put failure for %s: %d", str, index)
		}
	}
	// Remove odd ones
	for i := 0; i < 1000; i++ {
		if i%2 != 0 {
			str := fmt.Sprintf("str%d", i)
			st.Remove(str)
		}
	}
	if len(st.t) != sliceSize {
		t.Errorf("Wrong slice size, expected %d got %d", sliceSize, len(st.t))
	}
	// Must be same amount
	if st.Len() != len(m) {
		t.Errorf("Wrong len after remove: %d", st.Len())
	}
	// Can read whats written
	for k, v := range m {
		if st.Get(v) != k {
			t.Errorf("Got %s for %d but expected %s", st.Get(v), v, k)
		}
	}
	// Remove odd ones
	for i := 0; i < 1000; i++ {
		if i%2 != 0 {
			str := fmt.Sprintf("str%d", i)
			st.Remove(str)
			delete(m2, m[str])
			delete(m, str)
		}
	}
	// Must be half amount
	if st.Len() != len(m) {
		t.Errorf("Wrong len after remove: %d", st.Len())
	}
	if len(st.t) != sliceSize {
		t.Errorf("Wrong slice size, expected %d got %d", sliceSize, len(st.t))
	}
	// Add new ones
	for i := 0; i < 1000; i++ {
		str := fmt.Sprintf("str2 %d", i)
		index := st.Put(str)
		m[str] = index
		m2[index] = str
	}
	sliceSize += 500
	// size of slice must be correct
	if len(st.t) != sliceSize {
		t.Errorf("Wrong slice size, expected %d got %d", sliceSize, len(st.t))
	}
}

func benchmarkAdd(b *testing.B, k int) {
	strs := make([]string, 0, k)
	for i := 0; i < k; i++ {
		strs = append(strs, fmt.Sprintf("String %d", i))
	}
	for i := 0; i < b.N; i++ {
		st := StringTable{}
		for _, str := range strs {
			st.Put(str)
		}
		for _, str := range strs {
			st.Put(str)
		}
	}
}

func benchmarkAddMap(b *testing.B, k int) {
	strs := make([]string, 0, k)
	for i := 0; i < k; i++ {
		strs = append(strs, fmt.Sprintf("String %d", i))
	}
	for i := 0; i < b.N; i++ {
		st := make(map[string]int)
		for _, str := range strs {
			st[str] = 1
		}
		for _, str := range strs {
			st[str] = 1
		}
	}
}

func BenchmarkAdd10000(b *testing.B)    { benchmarkAdd(b, 10000) }
func BenchmarkAdd1000(b *testing.B)     { benchmarkAdd(b, 1000) }
func BenchmarkAddMap10000(b *testing.B) { benchmarkAddMap(b, 10000) }
func BenchmarkAddMap1000(b *testing.B)  { benchmarkAddMap(b, 1000) }
