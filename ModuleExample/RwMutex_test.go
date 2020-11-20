package ModuleExample

import (
	"testing"
)

//go test
func Test_RwMutexExample(t *testing.T) {
	RwMutexExample()
	t.Log("测试通过")
}

//go test -bench=”.”
func BenchmarkRwMutexExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RwMutexExample()
	}
}

//go test -bench=”BenchmarkAdd”
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
