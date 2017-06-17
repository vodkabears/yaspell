# yaspell
[![Build Status](https://travis-ci.org/vodkabears/yaspell.svg?branch=master)](https://travis-ci.org/VodkaBears/yaspell)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodkabears/yaspell)](https://goreportcard.com/report/github.com/vodkabears/yaspell)
[![Coverage Status](https://coveralls.io/repos/github/VodkaBears/yaspell/badge.svg)](https://coveralls.io/github/VodkaBears/yaspell)

Yaspell checks spelling of different texts with Yandex.Speller API.
The tool is targeted to people, who works with texts and wants to get fast feedback about the quality.
If you need to use API for your app, please, use [Yandex.Speller API](https://tech.yandex.ru/speller/doc/dg/concepts/api-overview-docpage/) directly or something else.

## Installation

Next platforms are supported: Windows, Mac, Linux.

Go the the release page and download a binary for your platform and architecture from the latest release: https://github.com/vodkabears/yaspell/releases/latest. Than you can rename a binary from `yaspell_{OS}_{ARCH}` to `yaspell`.

If you have installed golang environment, just make:
```
go get -u github.com/vodkabears/yaspell
```

## Usage

Run it in the terminal:
```
yaspell [flags] [files ...]
```

Example:
```
yaspell -opts=IGNORE_UPPERCASE,IGNORE_DIGITS,IGNORE_ROMAN_NUMERALS -dict=dict.txt text1.txt text2.txt
```

### Flags

#### -opts

Yandex.Speller options.

Example: `-opts=IGNORE_UPPERCASE,IGNORE_DIGITS`

`IGNORE_UPPERCASE` ignores uppercase words  
`IGNORE_DIGITS` ignores words with digits  
`IGNORE_URLS` ignores URLs, emails, filenames  
`FIND_REPEAT_WORDS` highlights repetitions of words, consecutive  
`IGNORE_LATIN` ignores Latin words  
`NO_SUGGEST` disables suggestions for incorrect words  
`FLAG_LATIN` marks Latin words as incorrect  
`BY_WORDS` ignores dictionary context  
`IGNORE_CAPITALIZATION` ignores the incorrect use of UPPERCASE/lowercase letters  
`IGNORE_ROMAN_NUMERALS` ignores Roman numerals  

#### -dict

Dictionary file with regular expressions.

Example: `-dict=dict.txt`  
Syntax of regular expressions : https://golang.org/pkg/regexp/syntax/#hdr-Syntax  

dict.txt content:
```
^lang$
^(G|g)olang$
```

#### -lang

Language to check.

Values: `en`, `ru`, `uk`  
Default: `ru,en`  
Example: `-lang=en,ru,uk`  

#### -format

Text format.

Values: `html`, `plain`  
Default: `plain`  
Example: `-format=html`

#### -version

Prints current version.

### Problems

```
2017/06/17 22:44:44 open ./i18n/en.js: too many open files
```

Solution:
```
ulimit -n 10000
```


