package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "golang.org/x/text/encoding"
    "golang.org/x/text/encoding/ianaindex"
    "github.com/spf13/pflag"
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
    fileNames := pflag.Args()
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
    pflag.StringVarP(&Allow, "allow", "a", "", "Additional characters to allow")
    pflag.StringVarP(&rawCharset, "charset", "c", "UTF-8", "Use this charset (coding)")
    pflag.BoolVarP(&help, "help", "h", false, "Print this help message")
    pflag.BoolVarP(&FilesWithMatches, "files-with-matches", "l", false, "Print file names only")
    pflag.BoolVarP(&Quiet, "quiet", "q", false, "Suppress output")
    pflag.Parse()
    if help {
        fmt.Println("SCCC: Source Code Character Checker")
        fmt.Println("syntax: sccc [options] [files]")
        fmt.Println("Options:")
        pflag.PrintDefaults()
        os.Exit(0)
    }
    GetCharset(rawCharset)
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
    name, err := ianaindex.IANA.Name(Charset)
    if err != nil {
        return
    }
    IsAscii = name == "US-ASCII"
    if name == "UTF-8" {
        // use internal codec
        Charset = nil
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
