#!/bin/bash

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 parse <filename>"
  exit 1
fi

command=$1
filename=$2
script_dir=$(dirname "$0")

if [ ! -f "$filename" ]; then
  echo "File $filename does not exist."
  exit 1
fi

case "$command" in
  parse)
    ${script_dir}/jlox "$filename"
    ;;
  *)
    echo "Unknown command: $command"
    echo "Usage: $0 parse <filename>"
    exit 1
    ;;
esac