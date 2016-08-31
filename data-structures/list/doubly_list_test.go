package list

import "testing"

func TestList(t *testing.T) {
	l := NewDLL()

	if !l.IsEmpty() {
		t.Error("List is expected to be empty but has length", l.Length)
	}

	for i := 11; i <= 21; i++ {
		l.Append(i)
	}

	for i := 10; i > 0; i-- {
		l.Prepend(i)
	}

	if err := l.Add(0, 0); err != nil {
		t.Error("Error while adding element", err)
	}

	if err := l.Remove(21); err != nil {
		t.Error("Error while removing element", err)
	}

	if l.Head.Value != 0 || l.Tail.Value != 20 || l.Length != 21 {
		t.Error("Expected value of head, tail and length are 0, 20, 21 but got ", l.Head.Value, l.Tail.Value, l.Length)
	}

	if l.Clear(); l.Length != 0 {
		t.Error("Expected length is zero")
	}

}

func BenchmarkList(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// fmt.Sprintf("hello")
	}
}
