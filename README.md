unpaste
=======

[![Build Status](https://travis-ci.org/kusabashira/unpaste.svg?branch=master)](https://travis-ci.org/kusabashira/unpaste)

Separate corresponding or subsequent lines of files.

```
$ echo "$PATH"
/bin:/usr/bin:/usr/local/bin:/bin:/usr/bin:/usr/local/bin

$ echo "$PATH" | unpaste -sd: | awk '!a[$0]++' | paste -sd:
/bin:/usr/bin:/usr/local/bin
```

```
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

Options:
  -d, --delimiters=LIST   reuse characters from LIST instead of TABs
  -s, --serial            unpaste one file at a time instead of in parallel
      --help              display this help and exit
      --version           display version information and exit
```

License
-------

MIT License

Author
------

kusabashira <kusabashira227@gmail.com>
