# Log Parser in Golang

## `bufio` Scanners

`bufio`'s scanner reads any input stream (terminal, network, file) into a buffer line-by-line. The `.Scan()` method reads a line into a `[]byte` buffer. Then the `.Text()` method retrieves the _last_ line and converts it to a string. You can also get the bytes out with `.Bytes()`. `Scan` returns `false` when there are no more lines.

[Example](https://yourbasic.org/golang/read-file-line-by-line/):

```go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

To split on words instead of lines, you can do this: `scanner.Split(bufio.ScanWords)`.
