package main

import(
  "fmt"
  "runtime"
)

func main() {

  n := runtime.NumCPU()
  plural := "s"
  if n < 2 { plural = "" }
 
  fmt.Printf("This system is %s %s and it has %d CPU%s.\n",
    runtime.GOOS, runtime.GOARCH, n, plural)

}
