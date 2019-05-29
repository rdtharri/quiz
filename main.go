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
      exit(fmt.Sprintf("Failed to open the csv file: %s", *csvFilename))
    }
    r := csv.NewReader(file)
}

func exit(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}
