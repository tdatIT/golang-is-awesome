package oop

import (
	"fmt"
	"strconv"
	"strings"
)

type Coder struct {
	Experiences         string
	ProgrammingLanguage []string
	Salary              int
	User                //embedding
}

func (coder *Coder) CreateCV() {
	fmt.Println(coder)
}

func (coder *Coder) String() string {
	return "{" + "exp:" + coder.Experiences +
		",programming_language:" + strings.Join(coder.ProgrammingLanguage, ",") + "salary:" +
		strconv.Itoa(coder.Salary) + "}"
}
