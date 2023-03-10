#!/bin/python

# sanitizes copied arrays from the av1 spec


def main():
    data = '''    
    { 11981198,, 3276832768,, 00 },},
{{ 20702070,, 3276832768,, 00 },},
{{ 91669166,, 3276832768,, 00 },},
{{ 74997499,, 3276832768,, 00 },},
{{ 2247522475,, 3276832768,, 00 }}
    '''

    data = data[data.find("{"):]
    data = data.replace("{{", "{")
    data = data.replace("}}", "},")
    data = data.replace("},},", "},")
    data = data.replace(",,", ",")
    numbers = data.split()

    for n in numbers:
        if not "{" in n and not "}" in n:
            if not n.endswith(','):
                print(n[int(len(n)/2):] + ",", end ='')
            else:
                print(n[int(len(n)/2):], end='')
        else:
            print(n)

if __name__ == '__main__':
    print("-------------------------------- OUTPUT -------------------------")
    main()
