package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
  csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
    flag.Parse()

    file, err := os.Open(*csvFilename)
    if err != nil {
      exit("Failed to open the CSV file.")
    }
    _ = file
}

func exit(msg string) {
  fmt.Printf(msg + "\n")
  os.Exit(1)
}
