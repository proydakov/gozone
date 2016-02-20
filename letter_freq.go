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
    value_dict map[string]int
    stats_dict map[string]int
}

func (c *FreqCalculator) InitDict(alphabet []rune) {
    c.alphabet = alphabet
    c.value_dict = make(map[string]int)
    for _, letter := range c.alphabet {
        text_letter := string(letter)
        //fmt.Println("Insert:", text_letter)
        c.value_dict[text_letter] = 0
    }
}

func (c *FreqCalculator) Process(r rune) {
    letter := strings.ToLower( string(r) )

    if v, ok := c.value_dict[letter]; ok {
        c.value_dict[letter] = v + 1
        //fmt.Printf("Process: '%s'\n", letter)
    } else {
        //fmt.Printf("Skip: '%s'\n", letter)
    }
}

func (c *FreqCalculator) CalcStatistic() {
    counter := 0
    for _, value := range c.value_dict {
        counter += value
    }

    c.stats_dict = make(map[string]int)
    for key, value := range c.value_dict {
        c.stats_dict[key] = value * 100.0 / counter
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
    calc.CalcStatistic()

    for k, v := range calc.stats_dict {
        fmt.Println("Key:", k, "Value:", v, "%")
    }
}
