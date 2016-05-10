unpaste
=======

[![Build Status](https://travis-ci.org/kusabashira/unpaste.svg?branch=master)](https://travis-ci.org/kusabashira/unpaste)

Separate corresponding or subsequent lines of files.

```
$ echo "$PATH"
/bin:/usr/bin:/usr/local/bin:/bin:/usr/bin:/usr/local/bin

$ echo "$PATH" | unpaste -sd: | awk '!a[$0]++' | paste -sd:
/bin:/usr/bin:/usr/local/bin


$ seq 3 | awk '{OFS="\t"; print $0, $0^2, $0^3}' | unpaste a b -
1
8
27

$ cat a
1
2
3

$ cat b
1
4
9
```

Usage
-----

```
$ unpaste [OPTION]... [FILE]...
Separate corresponding or subsequent lines of files.

With no FILE, or when FILE is -, output to standard output.

Options:
  -d, --delimiters=LIST   reuse characters from LIST instead of TABs
  -s, --serial            unpaste one file at a time instead of in parallel
      --help              display this help and exit
      --version           display version information and exit
```

Options
-------

### --help

Display a help message.

### --version

Display the version of unpaste.

### -d, --delimiters=LIST

Separate lines by specified the list of unicode characters.

```
$ cat src
1:2:3
4:5:6
7:8:9

$ cat src | unpaste -d: - a b
1
4
7

$ cat a
2
5
8

$ cat b
3
6
9


$ cat src2
hit|hit>hit|...
run|ran>run|...
say|said>said|...

$ cat src2 | unpaste -d'|>' - a b c
hit
run
say

$ cat a
hit
ran
said

$ cat b
hit
run
said

$ cat c
...
...
...
```

### -s, --serial

Swap delimiters of files and delimiters of lines.

```
$ cat src
1	2	3
4	5	6
7	8	9

$ cat src | unpaste -s - a b
1
2
3

$ cat a
4
5
6

$ cat b
7
8
9


$ cat src2
a:b:c:d:e
f:g:h:i:j

$ cat src2 | unpaste -sd:
a
b
c
d
e
```

Other specification
-------------------

#### `-` in `[FILE]...`

`-` in `[FILE]...` is interpreted as `/dev/stdout`.

```
$ printf "1\t2\t3\n4\t5\t6\n" | unpaste -
1
4

$ printf "1\t2\t3\n4\t5\t6\n" | unpaste - /dev/null -
1
3
4
6
```

License
-------

MIT License

Author
------

kusabashira <kusabashira227@gmail.com>
