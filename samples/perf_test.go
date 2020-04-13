package main 

import "testing"

func BenchmarkPrintfifty(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PrintFifty();
  }
}


