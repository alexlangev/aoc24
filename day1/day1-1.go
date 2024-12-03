package main

import (
  "fmt"
  "slices"
  "strconv"
  "bufio"
  "strings"
  "os"
)

func createLists() (l1, l2 []int){
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("os.Open error", err)
    return
  }

  defer file.Close()
  
  reader := bufio.NewReader(file)

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    } 

    sub := strings.Split(line, " ")
    // TODO why are index 1 and 2 whitespaces??
    x, _ := strconv.Atoi(strings.TrimSpace(sub[0]))
    l1 = append(l1, x)
    y, _ := strconv.Atoi(strings.TrimSpace(sub[3]))
    l2 = append(l2, y)
  }

  return l1, l2
}

func solveDay1pt1() {
  list1 := []int{}
  list2 := []int{}
  list1, list2 = createLists()

  slices.Sort(list1)
  slices.Sort(list2)

  if len(list1) != len(list2) {
    fmt.Println("The list are not the same length")
    return
  }

  sum := 0

  for i := 0; i < len(list1); i++ {
    if list1[i] >= list2[i] {
      sum += list1[i] - list2[i]
    } else {
      sum += list2[i] - list1[i]
    }
  }

  fmt.Println("The sum of the distances is :", sum)
}

