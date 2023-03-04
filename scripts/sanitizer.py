#!/bin/python

import sys

# sanitizes copied arrays from the av1 spec

def main():
    data = sys.argv[1]
    data = data[data.find("{"):]
    data = data.replace("{{", "{")
    data = data.replace("}}", "}")
    data = data.replace("},},", "},")
    data = data.replace(",,", ",")
    numbers = data.split()

    for n in numbers:
        if not "{" in n and not "}" in n:
            if not n.endswith(','):
                print(n[int(len(n)/2):] + ",")
            else:
                print(n[int(len(n)/2):])
        else:
            print(n)

if __name__ == '__main__':
    print("-------------------------------- OUTPUT -------------------------")
    main()
