package samples

import "testing"

func TestFib(t *testing.T){
  result := Fib(8)
  expected := 21
  if result != expected {
    t.Fatal("failed")
  }
}
