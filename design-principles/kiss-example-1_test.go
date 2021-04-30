package design_principles

import "testing"

func BenchmarkIsValidIPAddressV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidIPAddressV1("1.1.1.1")
	}
}

func BenchmarkIsValidIPAddressV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidIPAddressV2("1.1.1.1")
	}
}
