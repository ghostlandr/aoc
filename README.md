# aoc - Advent of Code scaffolding generator

I found I was doing the exact same thing to start every day of Advent of Code, so I figured I could automate it.

## What does it do?

It creates a new folder in the current directory called d1 (if you're starting day 1). It will also create three new
files inside of that directory. `input.txt`, where you should put your puzzle input. `part1.<extension>` and `part2.<extension>`
where your code should go. In the part files it will setup an appropriate starting template that will read in the 
`input.txt` file and print each line of it. This is to get you started as I found I was doing this every day of AoC.

## Installation

`go get github.com/gholtslander-va/aoc`

## Usage

Go to the folder you are doing your Advent of Code in and run

`aoc start <day number>`

If you are using Golang, it will generate go by default. If you want something else, use the language flag:

`aoc start <day number> -l python`

And it will generate python files.

# License

This project is licensed by the Apache 2.0 license. See LICENSE for full details.
