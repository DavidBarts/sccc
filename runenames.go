package main

// Go, like Python, has a broken utility function for determining the name
// of a rune. The control characters are missing! Work around that.

import (
    "golang.org/x/text/unicode/runenames"
)

var _CONTROL_NAMES = map[rune]string {
    '\x00': "NULL",
    '\x01': "START OF HEADER",
    '\x02': "START OF TEXT",
    '\x03': "END OF TEXT",
    '\x04': "END OF TRANSMISSION",
    '\x05': "ENQUIRY",
    '\x06': "ACKNOWLEDGEMENT",
    '\x07': "BELL",
    '\x08': "BACKSPACE",
    '\x09': "HORIZONTAL TAB",
    '\x0a': "LINE FEED",
    '\x0b': "VERTICAL TAB",
    '\x0c': "FORM FEED",
    '\x0d': "CARRIAGE RETURN",
    '\x0e': "SHIFT OUT",
    '\x0f': "SHIFT IN",
    '\x10': "DATA LINK ESCAPE",
    '\x11': "DEVICE CONTROL 1",
    '\x12': "DEVICE CONTROL 2",
    '\x13': "DEVICE CONTROL 3",
    '\x14': "DEVICE CONTROL 4",
    '\x15': "NEGATIVE ACKNOWLEDGEMENT",
    '\x16': "SYNCHRONOUS IDLE",
    '\x17': "END OF TRANSMISSION BLOCK",
    '\x18': "CANCEL",
    '\x19': "END OF MEDIUM",
    '\x1a': "SUBSTITUTE",
    '\x1b': "ESCAPE",
    '\x1c': "FILE SEPARATOR",
    '\x1d': "GROUP SEPARATOR",
    '\x1e': "RECORD SEPARATOR",
    '\x1f': "UNIT SEPARATOR",
    '\x7f': "DELETE",
    '\x80': "PADDING CHARACTER",
    '\x81': "HIGH OCTET PRESET",
    '\x82': "BREAK PERMITTED HERE",
    '\x83': "NO BREAK HERE",
    '\x84': "INDEX",
    '\x85': "NEXT LINE",
    '\x86': "START OF SELECTED AREA",
    '\x87': "END OF SELECTED AREA",
    '\x88': "CHARACTER TABULATION SET",
    '\x89': "CHARACTER TABULATION WITH JUSTIFICATION",
    '\x8a': "LINE TABULATION SET",
    '\x8b': "PARTIAL LINE DOWN",
    '\x8c': "PARTIAL LINE BACKWARD",
    '\x8d': "REVERSE INDEX",
    '\x8e': "SINGLE SHIFT TWO",
    '\x8f': "SINGLE SHIFT THREE",
    '\x90': "DEVICE CONTROL STRING",
    '\x91': "PRIVATE USE ONE",
    '\x92': "PRIVATE USE TWO",
    '\x93': "SET TRANSMIT STATE",
    '\x94': "CANCEL CHARACTER",
    '\x95': "MESSAGE WAITING",
    '\x96': "START OF GUARDED AREA",
    '\x97': "END OF GUARDED AREA",
    '\x98': "START OF STRING",
    '\x99': "SINGLE GRAPHIC CHARACTER INTRODUCER",
    '\x9a': "SINGLE CHARACTER INTRODUCER",
    '\x9b': "CONTROL SEQUENCE INTRODUCER",
    '\x9c': "STRING TERMINATOR",
    '\x9d': "OPERATING SYSTEM COMMAND",
    '\x9e': "PRIVACY MESSAGE",
    '\x9f': "APPLICATION PROGRAM COMMAND",
}

const _UNKNOWN_CONTROL_CHAR string = "<control>"

func RuneName(r rune) string {
    if name, found := _CONTROL_NAMES[r]; found {
        return name
    }
    ret := runenames.Name(r)
    if ret == _UNKNOWN_CONTROL_CHAR {
        return ""
    } else {
        return ret
    }
}
