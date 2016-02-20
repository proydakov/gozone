package main

import (
    "os"
    "fmt"
    "strings"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type FreqCalculator struct{
    alphabet []rune
    dict map[string]int
}

func (c *FreqCalculator) InitDict(alphabet []rune) {
    c.alphabet = alphabet
    c.dict = make(map[string]int)
    for _, letter := range c.alphabet {
        text_letter := string(letter)
        //fmt.Println("Insert:", text_letter)
        c.dict[text_letter] = 0
    }
}

func (c *FreqCalculator) Process(r rune) {
    letter := strings.ToLower( string(r) )

    if v, ok := c.dict[letter]; ok {
        c.dict[letter] = v + 1
        //fmt.Printf("Process: '%s'\n", letter)
    } else {
        //fmt.Printf("Skip: '%s'\n", letter)
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: application filename")
        return;
    }

    testFileName := os.Args[1]
    fmt.Printf("Input file: %s\n", testFileName)

    data, err := ioutil.ReadFile(testFileName)
    check(err)
    runedata := []rune(string(data))

    calc := new(FreqCalculator)

    en_alphabet := string("abcdefghijklmnopqrstuvwxyz")
    calc.InitDict( []rune( en_alphabet ) )

    for _, element := range runedata {
        calc.Process(element)
    }

    /*
    for k, v := range calc.dict {
        fmt.Println("Key:", k, "Value:", v)
    }
    fmt.Printf("InitDict: %v\n", len(calc.dict))
    */
    fmt.Println(calc)
}
