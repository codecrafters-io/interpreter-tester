import argparse
import sys

from lox.lox import Lox


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('command')
    parser.add_argument('filename')
    args = parser.parse_args()

    lox = Lox(args.filename)
    # fmt: off
    match args.command:
        case 'tokenize' : lox.scan()
        case 'parse'    : lox.parse()
        case 'evaluate' : lox.evaluate()
        case _          : sys.exit(1)
    # fmt: on


if __name__ == '__main__':
    main()
