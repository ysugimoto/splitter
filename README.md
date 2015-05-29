# splitter

String / Byte split by sep string / byte ignore the string sequence.

## Installation

```
$ go get github.com/ysugimoto/splitter
```

## Usage

Simple string split same `strings.Split`:

```
str := "Lorem ipsum dolor sit amet"
splitted := splitter.Split(str, " ")

fmt.Println(splitted)
// => []string{"Lorem", "ipsum", "dolor", "sit", "amet"}
```

spitter correctly split when sep string found in string sequence:

```
// famous case, command args split
str := "echo \"foo | bar | baz\" | grep bar"
splitted := splitter.Split(str, "|")

// strings.Split works all sep string
fmt.Println(strings.Split(str, "|")
// => []string{"echo \"foo ", " bar ", " baz\" ", " grep bar"}

// splitter.Split works out of sequence only
fmt.Println(splitted)
// => []string{"echo \"foo | bar | baz\" ", " grep bar"}
```

## Author

Yoshiaki Sugimoto

## License

MIT



