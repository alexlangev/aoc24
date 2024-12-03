package main

import (
  "fmt"
  //"slices"
  //"strconv"
  //"bufio"
  //"strings"
  //"os"
)

func solveDay1pt2() {
  list1 := []int{}
  list2 := []int{}
  list1, list2 = createLists()

  // Create map of l2
  l2map := make(map [int]int)
  for _, v := range list2 {
    l2map[v] += 1  
  }

  // calculate score
  score := 0
  for _, val := range list1 {
    score += val * l2map[val]
  }

  fmt.Println("The similarity score is :", score)
}

