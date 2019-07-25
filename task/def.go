/*
* @Author: scottxiong
* @Date:   2019-07-25 19:58:24
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 21:11:44
 */
package task

import (
	"fmt"
	"time"
)

const (
	PROJECT_FOLDER = "/Users/apple/go/src/github.com/scott-x/templates"
)

var tasks []task

type timeout interface {
	takes() time.Duration
}

type task struct {
	name string
	fn   func()
}

func printTime(t timeout) {
	fmt.Printf("It takes %v in total...\n", t.takes())
}
