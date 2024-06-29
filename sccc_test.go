package main

import (
    "testing"
    "path/filepath"
)

func setUp() {
    MyName = "sccc_test"
    Allow = ""
    Charset = nil
    IsAscii = false
    FilesWithMatches = false
    Quiet = true
    Status = 0
}

func testFile(fileName string) {
    ProcessFile(filepath.Join("testdata", fileName))
}

func testStdFile(fileName string) {
    setUp()
    testFile(fileName)
}

func shouldNotCrash(t *testing.T) {
    if Status & 2 != 0 {
        t.Fatal("an unexpected error occurred")
    }
}

func shouldFail(t *testing.T) {
    shouldNotCrash(t)
    if Status & 1 != 1 {
        t.Fatal("file passed, but should have failed")
    }
}

func shouldPass(t *testing.T) {
    shouldNotCrash(t)
    if Status & 1 != 0 {
        t.Fatal("file failed, but should have passed")
    }
}

// an empty file should pass (use actual file, not /dev/null, for portability)
func TestEmpty(t *testing.T) {
    testStdFile("empty.txt")
    shouldPass(t)
}

// something that contains plain ASCII should pass
func TestAscii(t *testing.T) {
    testStdFile("ascii.txt")
    shouldPass(t)
}

// something that contains valid UTF-8 Unicode should pass
func TestUtf8(t *testing.T) {
    testStdFile("utf8.txt")
    shouldPass(t)
}

// something that contains ISO-8859-1 should fail in default (UTF-8) mode
func TestLatin1FailsInUtf8Mode(t *testing.T) {
    testStdFile("latin1.txt")
    shouldFail(t)
}

// ... but it should pass in ISO-8859-1 mode
func TestLatin1(t *testing.T) {
    setUp()
    GetCharset("iso-8859-1")
    testFile("latin1.txt")
    shouldPass(t)
}

// total garbage should fail
func TestRandomGarbage(t *testing.T) {
    testStdFile("garbage.dat")
    shouldFail(t)
}

// horizontal tabs should fail
func TestTabs(t *testing.T) {
    testStdFile("makefile.txt")
    shouldFail(t)
}

// ... but they should pass if allowed
func TestTabsAllowed(t *testing.T) {
    setUp()
    Allow = "\t"
    testFile("makefile.txt")
    shouldPass(t)
}

// NBSP should fail
func TestNbsp(t *testing.T) {
    testStdFile("nbsp.txt")
    shouldFail(t)
}

// ZWSP should fail
func TestZwsp(t *testing.T) {
    testStdFile("zwsp.txt")
    shouldFail(t)
}

// BEL should fail
func TestBell(t *testing.T) {
    testStdFile("bell.txt")
    shouldFail(t)
}

// ESC should fail
func TestEscape(t *testing.T) {
    testStdFile("cls.txt")
    shouldFail(t)
}

// CR should pass
func TestCr(t *testing.T) {
    testStdFile("crlf.txt")
    shouldPass(t)
}

// Cyrillic text should pass
func TestCyrillic(t *testing.T) {
    testStdFile("russian.txt")
    shouldPass(t)
}

// Arabic text should pass
func TestArabic(t *testing.T) {
    testStdFile("arabic.txt")
    shouldPass(t)
}

// Native American text should pass
func TestNativeAmerican(t *testing.T) {
    testStdFile("chinuk.txt")
    shouldPass(t)
}

// Chinese text should pass
func TestChinese(t *testing.T) {
    testStdFile("chinese.txt")
    shouldPass(t)
}

// Japanese (katakana and hiragana) should pass
func TestJapanese(t *testing.T) {
    testStdFile("japanese.txt")
    shouldPass(t)
}

// Korean should pass
func TestKorean(t *testing.T) {
    testStdFile("korean.txt")
    shouldPass(t)
}

// Canadian Aboriginal syllabics should pass
func TestCanadianAboriginal(t *testing.T) {
    testStdFile("inuktitut.txt")
    shouldPass(t)
}

// emojis should pass
func TestEmojis(t *testing.T) {
    testStdFile("duck_turkey.txt")
    shouldPass(t)
}

// unassibned char should fail
func TestUnassigned(t *testing.T) {
    testStdFile("unassigned.txt")
    shouldFail(t)
}
