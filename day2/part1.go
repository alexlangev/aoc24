package main
// 87 86 84 83 81 80
// 52 51 48 46 44 41
// 97 94 93 92 91 88

import (
  "fmt"
  "io"
  "os"
  "bufio"
  "strings"
)

func solvePart1() {
  file, err := os.Open("input.txt")
  defer file.Close()
  if err != nil {
    fmt.Println("Error happened in os.Open", err)
    return
  }

  // Read line by line
  reader := bufio.NewReader(file)
  for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println("Error happened at reader.ReadString", err)
      return
    }

    trimmedLine := strings.TrimSpace(line)

    fmt.Println(trimmedLine)
  }
  

  // Eval the line cell by cell
  //    is strictly increasing?
  //    is strictly decreasing?
  //    each step is a min of 1 and max of 3

  // if line is safe inc the counter

  // print counter
  fmt.Println("test")
}

func evalReport(r string) bool {

}
