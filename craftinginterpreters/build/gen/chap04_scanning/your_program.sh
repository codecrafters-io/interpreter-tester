#!/bin/bash

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 tokenize <filename>"
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
  tokenize)
    ${script_dir}/jlox_04 "$filename"
    ;;
  *)
    echo "Unknown command: $command"
    echo "Usage: $0 tokenize <filename>"
    exit 1
    ;;
esac