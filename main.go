package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "golang.org/x/text/encoding"
    "golang.org/x/text/encoding/ianaindex"
)

var MyName string
var Allow string
var Charset encoding.Encoding
var IsAscii bool
var FilesWithMatches bool
var Quiet bool
var Status int

func main() {
    Status = 0
    parseArgs()
    fileNames := flag.Args()
    if len(fileNames) == 0 {
        Process("standard input", wrapReader(os.Stdin))
    } else {
        for _, fileName := range fileNames {
            ProcessFile(fileName)
        }
    }
    os.Exit(Status)
}

func ProcessFile(fileName string) {
    if inputFile := mustOpen(fileName); inputFile != nil {
        Process(fileName, wrapReader(inputFile))
        inputFile.Close()
    }
}

func parseArgs() {
    var rawCharset string
    var help bool
    MyName = filepath.Base(os.Args[0])
    shortLongString(&Allow, "a", "allow", "Additional characters to allow")
    shortLongString(&rawCharset, "c", "charset", "Use this charset (coding), instead of UTF-8")
    shortLongBool(&help, "h", "help", "Print this help message")
    shortLongBool(&FilesWithMatches, "l", "files-with-matches", "Print file names only")
    shortLongBool(&Quiet, "q", "quiet", "Suppress output")
    flag.Parse()
    if help {
        flag.PrintDefaults()
        os.Exit(0)
    }
    if rawCharset != "" {
        GetCharset(rawCharset)
    }
}

func shortLongBool(bp *bool, short string, long string, usage string) {
    flag.BoolVar(bp, long, false, usage + ".")
    flag.BoolVar(bp, short, false, usage + " (shorthand).")
}

func shortLongString(sp *string, short string, long string, usage string) {
    flag.StringVar(sp, long, "", usage + ".")
    flag.StringVar(sp, short, "", usage + " (shorthand).")
}

func GetCharset(rawCharset string) {
    var err error
    Charset, err = ianaindex.IANA.Encoding(rawCharset)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: unknown charset %q\n", MyName, rawCharset)
        os.Exit(2)
    }
    if Charset == nil {
        fmt.Fprintf(os.Stderr, "%s: unsupported charset %q\n", MyName, rawCharset)
        os.Exit(2)
    }
    if name, err := ianaindex.IANA.Name(Charset); err == nil && name == "US-ASCII" {
        IsAscii = true
    } else {
        IsAscii = false
    }
}

func wrapReader(reader io.Reader) io.RuneReader {
    if Charset == nil {
       return bufio.NewReader(reader)
    } else {
       // Decoder.Reader() returns something that buffers its input but is not
       // a RuneReader. The following code will need to change if that changes.
       return bufio.NewReader(Charset.NewDecoder().Reader(reader))
    }
}

func mustOpen(fileName string) io.ReadCloser {
    ret, err := os.Open(fileName)
    if err != nil {
        // TODO: make sure this message shows the file name
        fmt.Fprintf(os.Stderr, "%s: %s\n", MyName, err)
        Status |= 2
        return nil
    }
    return ret
}
