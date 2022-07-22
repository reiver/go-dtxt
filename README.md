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

This is a basic example of how to encode **tabular** data into **ASCII delimited text** using this package:
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

## Decoding Example

This is a basic example of how to dencode **tabular** data from **ASCII delimited text** using this package.

In this example it is known ahead of time how many columns there are in the data.

```go

import "github.com/reiver/go-dtxt"

// ...

var reader io.Reader //@TODO: set to wherever you want the encoded ASCII Delimited Text data to come from.

// ...

var decoder dtxt.Decoder = dtxt.WrapDecoder(reader)

// ...

for {
	var key string
	var value string
	
	err := decoder.DecodeRow(&key, &value)
	if dtxt.GS == err {
		break
	}
	if nil != err {
		return err
	}
}

```


## Deliminators

Unicode inherited 5 deliminator **control code** characters from ASCII:

| Symbol | Name                      | Alternative Name | Abbreviation | Hexadecimal | Decimal | Caret     | UTF-8        |
|--------|---------------------------|------------------|--------------|-------------|---------|-----------|--------------|
| ␜      | File Separator            |                  | FS           |        0x1c |      28 | `` ^\ ``  | `0b00011100` |
| ␝      | Group Separator           | Table Terminator | GS           |        0x1d |      29 | `` ^] ``  | `0b00011101` |
| ␞      | Row Separator             | Row Terminator   | RS           |        0x1e |      30 | `` ^^ ``  | `0b00011110` |
| ␟      | Unit Separator            | Field Terminator | US           |        0x1f |      31 | `` ^_ ``  | `0b00011111` |
| ␠      | Space                     | Word Separator   | SP           |        0x20 |      32 | `` ^` ``  | `0b00100000` |


## Table Row Format

**Unit Separator** (**US**) and **Row Separator** (**RS**) can be used to construct a table row.

For example, if we wanted to have a table row with 3 fields: “`joe`”, “`blow`”, and “`root beer`”. I.e,. —
 
| | | |
|-|-|-|
| joe | blow | root beer |
 

Then the result would be this:
```go
const US = 0x1f
const RS = 0x1e

[]byte{
	'j','o','e', 
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'b','l','o','w',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'r','o','o','t',' ','b','e','e','r',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	RS, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Row Terminator
}
```

(Note this is just a single row.
And _not_ a whole table.
A whole table would have a `GS` control code character at the end of it.)

⚠️ Notice that we are using the `US` control code characters in the Unix/Linux style — as a field **terminator** (and _not_ just a field **separator**).
I.e., the last field gets a `US` after it too.

⚠️ Notice also that we are using the `RS` control code character in the Unix/Linux style too — as a row **terminator** (and _not_ just a row **separator**).
I.e., the last row gets a `RS` after it too.

## Table Format

Let's make it more obvious how `RS` is used by showing a whole table encoded (and not just a row).
Let's encode this table:

| | | |
|-|-|-|
| joe | blow | root beer |
| john | doe | caramel apple |
| jane | doe | cotton candy |

```go
const US = 0x1f
const RS = 0x1e
const GS = 0x1d

[]byte{
	'j','o','e', 
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'b','l','o','w',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'r','o','o','t',' ','b','e','e','r',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	RS, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Row Terminator



	'j','o','h','n',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'd','o','e',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'c','a','r','a','m','e','l',' ','a','p','p','l','e',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	RS, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Row Terminator



	'j','a','n','e',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'd','o','e',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	'c','o','t','t','o','n',' ','c','a','n','d','y',
	US, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Field Terminator
	
	RS, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Row Terminator



	GS, // ⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚⇚ Table Terminator
}
```

⚠️ Notice that we are using the `GS` control code characters in the Unix/Linux style — as a table **terminator** (and _not_ just a table **separator**).
I.e., the last rows gets a `GS` after it.

## Escaping

One issue that can arise is — what if the data inside of a **unit** contains a **Unit Separator** (**US**), a **Row Separator** (**RS**), a **Group Separator** (**GS**), or a **File Separator** (**FS**)‽

How is that situation handled‽

The answer is that — Unicode inherited a **control code** character for **escaping**.
The aptly named **Escape** (**ESC**) **control code** character:


| Name    | Abbreviation | Hexadecimal | Decimal | Caret     | UTF-8        |
|---------|--------------|-------------|---------|-----------|--------------|
| Escape  | ESC          |        0x1b |      27 | `` ^[ ``  | `0b00011011` |


An **ESC** chararacter is stuffed before any **Unit Separator** (**US**), **Row Separator** (**RS**), **Group Separator** (**GS**), or **File Separator** (**FS**) that appears inside of a **unit**.
