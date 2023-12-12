#!/usr/bin/env bash

usage() {
  echo "Usage: $0 <day number> <puzzle part> [input part]"
  echo "  day number:  a number from 1 to 24"
  echo "  puzzle part: 1 or 2"
  echo "  input part:  (optional) 1 for 'test input' or 2 for 'puzzle input'; defaults to puzzle part if omitted"
  exit 1
}

if [ "$#" -lt 2 ] || [ "$#" -gt 3 ]; then
  usage
fi

day_number=$(printf "%02d" "$1")
puzzle_part="$2"

if [[ ! "$day_number" =~ ^[0-9]{2}$ || "$day_number" -lt 1 || "$day_number" -gt 24 ]]; then
  echo "Error: Invalid day number. It should be a number from 1 to 24 (0 padded)."
  usage
fi

if [[ ! "$puzzle_part" =~ ^[1-2]$ ]]; then
  echo "Error: Invalid puzzle part. It should be either 1 or 2."
  usage
fi

input_part="$3"
if [ -z "$input_part" ]; then
  # If input_part is not provided, default to puzzle_part
  input_part="$puzzle_part"
fi

if [[ ! "$input_part" =~ ^[1-2]$ ]]; then
  echo "Error: Invalid input part. It should be either 1 or 2."
  usage
fi

input_type=""
case "$input_part" in
  1) input_type="test" ;;
  2) input_type="puzzle" ;;
  *) echo "Error: Invalid input part. It should be either 1 or 2."
     usage
     ;;
esac

echo "Running day$day_number #$puzzle_part using ${input_type} input:"
go run "src/day$day_number$puzzle_part/main.go" < "data/input$day_number$input_part.txt"
