package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kapibara824/uuid-generator/pkg/generator"
)

func main() {
	i, err := input()
	if err != nil {
		log.Fatal(err)
	}
	id, err := generator.GenUUID(i)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range id {
		fmt.Printf("%v\n", v)
	}

}
func input() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("生成するUUIDの数を入力してください:")
	scanner.Scan()
	str := scanner.Text()

	num, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		return int(num), err
	}
	return int(num), nil

}
