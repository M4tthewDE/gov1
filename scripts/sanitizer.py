#!/bin/python

# sanitizes copied arrays from the av1 spec
# this only happens in firefox afaik. Just use chrome instead please :)


def main():
    data = '''    
    '''

    data = data[data.find("{"):]
    data = data.replace("{{", "{")
    data = data.replace("}}", "},")
    data = data.replace("},},", "},")
    data = data.replace(",,", ",")
    numbers = data.split()

    result = ""
    for n in numbers:
        if not "{" in n and not "}" in n:
            if not n.endswith(','):
                print(n[int(len(n)/2):] + ",", end ='')
                result = result + n[int(len(n)/2):] + ","
            else:
                print(n[int(len(n)/2):], end='')
                result = result + n[int(len(n)/2):]
        else:
            print(n)
            result = result + n + '\n'

    #print(result)

    i = 0
    balance = 0
    for c in result:
        if c == '{':
            balance += 1
        if c == '}':
            balance -= 1

        i += 1
        if balance < 0:
            print("ERROR: NEGATIVE BALANCE ACHIEVED AT i ", i)
            break;

    if balance != 0:
        print("ERROR: UNEVEN BALANCE:", balance)
    else:
        print("-------------------------------- END OUTPUT -------------------------")

if __name__ == '__main__':
    print("-------------------------------- OUTPUT -------------------------")
    main()
