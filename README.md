SCCC: Source Code Character Checker
===================================

I wrote this program because occasionally I have had real headaches with
garbage or invisible characters in files, particularly source files. A
zero-width space in the middle of an identifier will be parsed as two
separate identifiers. A Python file indented with a mix of spaces and
tabs can fail to have its indentation recognized properly, and thus
parse into execution blocks incorrectly. And so on.

There not being any standard utilities I could find to detect and report
such things, I decided to write one.

For each offending character, the numeric code point, the descriptive
name of the character (if available), and the line and character offsets
are repored. For example:

    $ sccc gunk
    gunk: U+00A0 (NO-BREAK SPACE) at char 4 line 2
    gunk: U+200B (ZERO WIDTH SPACE) at char 5 line 3
    gunk: U+0009 (HORIZONTAL TAB) at char 1 line 4

Files are by default assumed to contain UTF-8 encoded text, but this can
be changed with the `-c` option.

Building
--------

**Prerequisites:** A Go compiler is required. If for some reason you are
editing the README for this project, [Pandoc](https://pandoc.org/) is
required.

Edit `Makefile` and set the `PANDOC_DIR` and `GO_DIR` variables as
needed. If the directories containing `pandoc` and/or `go` are in your
standard search path, define the corresponding variable to be the empty
string.

Then just type:

    make sccc

Installing
----------

Copy the resulting `sccc` binary to the directory of your choice.

Running
-------

The general command syntax is:

> `sccc` \[*options*\] \[*file* ...\]

Typing `sccc -h` will print a help message listing the available
options. If no files are specified, standard input will be processed.
