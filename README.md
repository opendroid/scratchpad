# Scrachpad for various go-lang examples

## ex0

Spins nFuncs go-routines that increments a shared variable safely.

## ex1

Test impact of golang version in `go.mod` file. Example is a generics based function introduced in go1.18.
It fails compilation when go version is set to < 1.18

## ex2

Example flag project. flag.Parse() parses flags until first non-flag.
`ex2 -flagBool -flagString="hi" subcommand -flag2Bool` will not parse flag2Bool and beyond.

## ex3

Example designed multiple subcommands and their flags.

## js

Contains JS examples.
 - t1.js tests JSON parsing and stringifying.

## OpenAI

Contains examples of using OpenAI API.

### ChatGPT

ChatGPT is a chatbot that uses GPT-3 to generate responses to user input. This example tests the API and trains a model.

