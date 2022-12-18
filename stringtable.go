package stringtable

// A StringTable keeps a table of strings such that every unique
// string is represented using an integer value. This implementation
// of string table uses reference counting to maintain an active set
// of strings. That means, if you use a string, you have to get the
// index of that string using `Put()`, and once you no longer use that
// string, you should remove it using `Remove()` or `RemoveIndex()`.
//
// All StringTable methods are O(1) operations.
//
// The zero-value of the StringTable is ready for use.
type StringTable struct {
	m         map[string]int
	t         []tableItem
	firstFree int
}

type tableItem struct {
	str string
	c   int
}

// Put places the string into the string table if it is not there
// already, and returns its index. If the string is already in the
// table, then its reference count is incremented by one.
//
// Once you no longer use the string, you should Remove() it
func (s *StringTable) Put(str string) int {
	if s.m == nil {
		s.init()
	}
	index, ok := s.m[str]
	if ok {
		s.t[index].c++
		return index
	}
	index = s.firstFree
	if index != -1 {
		s.firstFree = s.t[index].c
		s.t[index].c = 1
		s.t[index].str = str
		s.m[str] = index
		return index
	}
	s.t = append(s.t, tableItem{
		str: str,
		c:   1,
	})
	s.m[str] = len(s.t) - 1
	return len(s.t) - 1
}

// Get returns the string for the given index
func (s *StringTable) Get(index int) string {
	return s.t[index].str
}

// Remove a string from the string table. The string is actually
// removed when its reference count reaches zero.
func (s *StringTable) Remove(str string) {
	index, ok := s.m[str]
	if !ok {
		panic("String not in table:" + str)
	}
	s.t[index].c--
	if s.t[index].c > 0 {
		return
	}
	s.t[index].c = s.firstFree
	s.firstFree = index
	delete(s.m, str)
}

func (s *StringTable) init() {
	if s.m == nil {
		s.m = make(map[string]int)
		s.firstFree = -1
	}
}

// Len returns the number of unique strings in the table
func (s *StringTable) Len() int {
	return len(s.m)
}
