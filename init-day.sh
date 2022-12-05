#!/bin/bash

# Check if the user provided a number as an argument
if [ $# -eq 0 ]; then
  echo "Please provide a number as an argument"
  exit 1
fi

# Check if the src/day directory already exists
if [ -d "src/day$1" ]; then
  echo "The src/day$1 directory already exists"
  exit 1
fi

# Create the src/day directory if it doesn't already exist
mkdir -p src/day$1

# Create three empty files with names that are suffixed with the provided number
touch src/day$1/dataset.txt
touch src/day$1/day$1.go
touch src/day$1/day$1_test.go

# Add the empty ChallengeOfTheDay function to the day.go file
echo "package day$1

func ChallengeOfTheDay() {}" >> src/day$1/day$1.go

# Add the test function to the day_test.go file
echo "package day$1

func TestChallengeOfTheDay(t *testing.T) {}" >> src/day$1/day$1_test.go