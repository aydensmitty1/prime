package prime

import (
  "fmt"
  "time"
)


func Pnums(max int){
  start := time.Now()
  numbers := make([]int, max)
  for i:=0; i<max; i++ {
    numbers[i] = i+1
  }


  for i:=0; i<len(numbers); i++ {
    for j:=2; j<max;j++{
      if i >= len(numbers) {
        break
      }

      if numbers[i]%j==0 && numbers[i]!=j{

        numbers = removal(i, numbers)
      }

  }
  }

  duration := time.Since(start)
  fmt.Printf("numbers: %v\n", numbers)
  fmt.Printf("ran in: %f seconds\n", duration.Seconds())
}

// how to remove an element from a slice
func removal(i int, numbers []int) []int {

  s1 := numbers[:i]
  s2 := numbers[i+1:]
  s3 := numbers[i:]


  if i>=len(numbers){
    numbers = append(s1,s3...)
  }else{
    numbers = append(s1,s2...)
  }

  return numbers
}


//how to add an element to a slice
func add(numbers []int) {
  numbers = append(numbers, 1)
}
