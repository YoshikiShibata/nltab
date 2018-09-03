// Copyright © 2018 Yoshiki Shibata. All rights reserved.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\\n", "\n", -1)
		line = strings.Replace(line, "\\t", "\t", -1)
		fmt.Println(line)
	}
}
