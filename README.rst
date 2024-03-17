###################################
SCCC: Source Code Character Checker
###################################

I wrote this program because occasionally I have had real headaches with
garbage or invisible characters in files, particularly source files. A
zero-width space in the middle of an identifier will be parsed as two
separate identifiers. A Python file indented with a mix of spaces and tabs
can fail to have its indentation recognized properly, and thus parse into
execution blocks incorrectly. And so on.

There not being any standard utilities I could find to detect and report
such things, I decided to write one.

For each offending character, the numeric code point, (if available) the
descriptive name, and the line and character offsets is repored. For
example::

    $ sccc gunk
    gunk: U+00A0 (NO-BREAK SPACE) at char 4 line 2
    gunk: U+200B (ZERO WIDTH SPACE) at char 5 line 3
    gunk: U+0009 (HORIZONTAL TAB) at char 1 line 4

Building
========

This section not yet finished.

Installing
==========

This section not yet finished.

Running
=======

This section not yet finished.

For more information, see the man page.
