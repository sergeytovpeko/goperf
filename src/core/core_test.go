package core_test

import (
    "testing"
    "fmt"
)

func BenchmarkAssign(b *testing.B) {
    var test int
    var test2 int

    for i := 0; i < b.N; i++ {
        test2 = test
    }
    fmt.Printf("%v\n", test2)
}

func BenchmarkTypeAssertion(b *testing.B) {
    type TestStruct struct {
        i int
    }
    var test TestStruct
    var test2 TestStruct
    var empty interface{}
    empty = TestStruct(test)

    for i := 0; i < b.N; i++ {
        test2 = empty.(TestStruct)
    }
    fmt.Printf("%v\n", test2)
}
