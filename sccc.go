package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "unicode"
    "golang.org/x/text/encoding/charmap"
)

// Bad runes read as this. XXX - so we can't tell if we got an actual
// replacement character, or are getting one due to a bad rune.
const _BAD rune = '\uFFFD'

// Characters we always allow, despite their being in a forbidden category.
// Space, line feed, carriage return.
const _ALWAYS_ALLOW string = " \n\r"

// Forbidden General_Categories. Characters in these categories are
// forbidden unless explicitly allowed. We forbid stuff which is likely to
// trip up a parser yet remain invisible when editing or printing. This
// includes the ^I horizontal tab character, which can confuse the Python
// parser.
var _FORBIDDEN = []*unicode.RangeTable{
    unicode.Cc, // Control
    unicode.Cf, // Format
    unicode.Cs, // Surrogate
    unicode.Co, // Private use
    unicode.Zs} // Space_separator

// Known characters, i.e. all defined Unicode code points, so that undefined
// ones can be reported
var _KNOWN []*unicode.RangeTable

// Where we are in the input file (1-based, origin at top left).
var row int
var col int

func init() {
    _KNOWN = make([]*unicode.RangeTable, len(unicode.Categories))
    i := 0
    for _, rt := range unicode.Categories {
        _KNOWN[i] = rt
        i++
    }
}

func Process(name string, reader io.RuneReader) {
    row = 1
    col = 1
    for {
        r, size, err := reader.ReadRune()
        if err != nil {
            break
        }
        processChar(name, r, size)
        if r == '\n' {
            row++
            col = 1
        } else {
            col++
        }
    }
}

func processChar(name string, r rune, size int) {
    // XXX - go provides no way for its x.text.Decoder class to distinguish
    // between reading a replacement character in its input, and one that was
    // inserted to indicate an invalid byte sequence. In-band signalling for
    // the loss.
    if r == _BAD {
        if Charset == nil && size == 1 {
            // Native UTF-8, no Decoder, we know it is bad.
            log("%s: invalid rune at char %d line %d\n", name, col, row)
            Status = 1
            return
        } else if Charset != nil {
            // Decoder used, might be in input.
            if _, isCharmap := Charset.(*charmap.Charmap); isCharmap || IsAscii {
                // 1-byte charset, not in input
                log("%s: invalid rune at char %d line %d\n", name, col, row)
                Status = 1
                return
            } else {
                // multibyte charset, no way to tell
                log("%s: possible invalid rune at char %d line %d\n", name, col, row)
                Status = 1
            }
        }
    }

    if strings.ContainsRune(_ALWAYS_ALLOW, r) || strings.ContainsRune(Allow, r) {
        return
    }
    if isForbidden(r) {
        expl := ""
        if runeName := RuneName(r); runeName != "" {
            expl = " (" + runeName + ")"
        }
        log("%s: %U%s at char %d line %d\n", name, r, expl, col, row)
        Status = 1
    }
}

func isForbidden(r rune) bool {
    return unicode.In(r, _FORBIDDEN...) || !unicode.In(r, _KNOWN...)
}

func log(format string, a ...any) {
    if !Quiet {
        fmt.Fprintf(os.Stderr, format, a...)
    }
}
