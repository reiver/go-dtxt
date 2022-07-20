# go-dtxt

Package **dtxt** implements encoding and decoding of **ASCII delimited text**, for the Go programming language.

**ASCII delimited text** is similar to CSV, TSV, and other table & spreadsheet data formats.
Except that **ASCII delimited text** uses some of the deliminator **control code** characters that Unicode inherited from ASCII.

**ASCII delimited text** could also probably be validly called **Unicode delimited text**.
Especially when Unicode is encoded as UTF-8.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-dtxt

[![GoDoc](https://godoc.org/github.com/reiver/go-dtxt?status.svg)](https://godoc.org/github.com/reiver/go-dtxt)

## Encoding Example

This is a basic example of how to encode **table** or **spreadsheet** data into **ASCII delimited text** using this package:
```go
import "github.com/reiver/go-dtxt"

// ...

var writer io.Writer //@TODO: set to wherever you want the encoded ASCII Delimited Text data to go.

// ...

var encoder dtxt.Encoder = dtxt.EncoderWrap(writer)

err := encoder.Begin()

// ...

defer encoder.End()

// ...

// row 1
err := encode.EncodeRow("ONCE", '۱', "1", "Ⅰ", "یکی")

// ...

// row 2
err := encode.EncodeRow("TWICE", '۲', "2", "Ⅱ". "دو")

// ...

// row 3
err := encode.EncodeRow("THRICE", '۳', "3", "Ⅲ", "سه")

// ...

// row 3
err := encode.EncodeRow("FOURCE", '۴', "3", "Ⅳ", "چهار")

// ...


```

## Deliminators

Unicode inherited 5 deliminator **control code** characters from ASCII:

| Symbol | Name                      | Alternative Name | Abbreviation | Hexadecimal | Decimal | Caret     | UTF-8        |
|--------|---------------------------|------------------|--------------|-------------|---------|-----------|--------------|
| ␜      | File Separator            |                  | FS           |        0x1c |      28 | `` ^\ ``  | `0b00011100` |
| ␝      | Group Separator           | Table Separator  | GS           |        0x1d |      29 | `` ^] ``  | `0b00011101` |
| ␞      | Row Separator             |                  | RS           |        0x1e |      30 | `` ^^ ``  | `0b00011110` |
| ␟      | Unit Separator            |                  | US           |        0x1f |      31 | `` ^_ ``  | `0b00011111` |
| ␠      | Space                     | Word Separator   | SP           |        0x20 |      32 | `` ^` ``  | `0b00100000` |


## Tables

**Unit Separator** (**US**) and **Row Separator** (**RS**) can be used to construct a table.

## Escaping

One issue that can arise is — what if the data inside of a **unit** contains a **Unit Separator** (**US**), a **Row Separator** (**RS**), a **Group Separator** (**GS**), or a **File Separator** (**FS**)‽

How is that situation handled‽

The answer is that — Unicode inherited a **control code** character for **escaping**.
The aptly named **Escape** (**ESC**) **control code** character:


| Name    | Abbreviation | Hexadecimal | Decimal | Caret     | UTF-8        |
|---------|--------------|-------------|---------|-----------|--------------|
| Escape  | ESC          |        0x1b |      27 | `` ^[ ``  | `0b00011011` |


An **ESC** chararacter is stuffed before any **Unit Separator** (**US**), **Row Separator** (**RS**), **Group Separator** (**GS**), or **File Separator** (**FS**) that appears inside of a **unit**.
