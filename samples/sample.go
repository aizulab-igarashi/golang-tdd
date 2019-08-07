package samples

func Fib(n int) int {
  switch(n){
  case 0, 1:
    return n
  default:
    return Fib(n-1) + Fib(n-2)
  }
}
