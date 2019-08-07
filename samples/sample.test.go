package samples

import "testing"

func TestFizzBuzz(t *testing.T){
  result := FizzBuzz(5)
  expected := "buzz"
  if result != expected {
    t.Fatal("failed")
  }
}
