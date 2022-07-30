package main

import (
	"encoding/json"
	"os"
	"strconv"
)

const INPUT_PATH string = "../pipes/input"
const OUTPUT_PATH string = "../pipes/output"

type response struct {
	Rtn string `json:"rtn"`
	Ok  bool   `json:"ok"`
}

const (
	ACTION int = iota
	TEXT       = iota
)

const (
	EMPTY_CLIPBOARD string = "empty clipboard"
	FULL_CLIPBOARD  string = "clipboard is full"
	NEGATIVE_SIZE   string = "negative size"
)

func readArgs() ([]string, error) {
	encoded, _ := os.ReadFile(INPUT_PATH)
	args := []string{}
	json.Unmarshal(encoded, &args)
	return args, nil
}

func writeResponse(rtn string, ok bool) {
	output, _ := os.OpenFile(OUTPUT_PATH, os.O_WRONLY, 0)
	encoded, _ := json.Marshal(response{rtn, ok})
	output.Write(encoded)
	output.Close()
}

func main() {
	var s Stack = Stack{sizeLimit: 1}

	for {
		args, _ := readArgs()

		var rtn string
		var ok bool

		switch action := args[ACTION]; action {
		case "add":
			if s.sizeLimit == 0 || s.size < s.sizeLimit {
				s.push(args[TEXT])
				ok = true
			} else {
				rtn = FULL_CLIPBOARD
			}
		case "list":
			if s.size == 0 {
				rtn = EMPTY_CLIPBOARD
			} else {
				ok = true
				rtn = s.print()
			}
		case "pop":
			if s.size == 0 {
				rtn = EMPTY_CLIPBOARD
			} else {
				rtn = s.pop()
				ok = true
			}
		case "front":
			if s.size == 0 {
				rtn = EMPTY_CLIPBOARD
			} else {
				rtn = s.front()
				ok = true
			}
		case "clear":
			s.clear()
			ok = true
		case "setLimit":
			n, err := strconv.Atoi(args[1])
			if err != nil {
				rtn = err.Error()
			}
			if n < 0 {
				rtn = NEGATIVE_SIZE
			}
			s.sizeLimit = n
			ok = true
		}
		writeResponse(rtn, ok)
	}
}
