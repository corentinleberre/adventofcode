#!/bin/bash

if [ $# -eq 0 ]; then
  echo "🔴 Please provide a number as an argument"
  exit 1
fi

if (("$1" < 1 || "$1" > 25))
then
  echo "🔴 Invalid argument: $1"
  exit 1
fi

if [ -d "src/day$1" ]; then
  echo "🔴 The src/day$1 directory already exists"
  exit 1
fi

mkdir -p src/day$1

touch src/day$1/dataset.txt
touch src/day$1/day$1.go
touch src/day$1/day$1_test.go

echo "package day$1

func ChallengeOfTheDay() {}" >> src/day$1/day$1.go

echo "package day$1

func TestChallengeOfTheDay(t *testing.T) {}" >> src/day$1/day$1_test.go

echo "👍 Great success 👍"