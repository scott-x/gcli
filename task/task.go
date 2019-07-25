/*
* @Author: scottxiong
* @Date:   2019-07-25 19:51:24
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-25 22:54:31
 */
package task

import (
	"github.com/fatih/color"
	"github.com/scott-x/gcmd"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"path"
	"strconv"
	"time"
)

var (
	cyan   = color.New(color.FgCyan, color.Bold)
	red    = color.New(color.FgRed, color.Bold)
	yellow = color.New(color.FgYellow, color.Bold)
	blue   = color.New(color.FgBlue, color.Bold)
)

func Run() {
	regestTask()
	render()
}

func (t task) takes() time.Duration {
	t1 := time.Now()
	t.fn()
	t2 := time.Now()
	return t2.Sub(t1)
}

func regestTask() {
	task1 := task{
		name: "react",
		fn: func() {
			cmd.AddQuestion("name", "What's your react name ? ", "Please input the correct name: ", "^[a-z]+")
			answers := cmd.Exec()
			cliName := answers["name"]
			fs.CopyFolder(path.Join(PROJECT_FOLDER, "react"), "/Users/apple/desktop/"+cliName)
		},
	}
	task2 := task{
		name: "angular",
		fn: func() {
			cmd.AddQuestion("name", "What's your angular name ? ", "Please input the correct name: ", "^[a-z]+")
			answers := cmd.Exec()
			cliName := answers["name"]
			fs.CopyFolder(path.Join(PROJECT_FOLDER, "react"), "/Users/apple/desktop/"+cliName)
		},
	}
	addTask(task1, task2)
}

func render() {
	if len(tasks) == 1 {
		yellow.Printf("Now is running task: %s\n", tasks[0].name)
		//printTime(tasks[0])
		tasks[0].fn()
	} else {
		yellow.Printf("Please select the task [1-%d]: \n", len(tasks))
		for k, t := range tasks {
			cyan.Printf("%d, %s\n", k+1, t.name)
		}
		gcmd.AddQuestion("select_number", "you selected: ", "Please input the correct number: ", "^[1-9]+")
		answers := gcmd.Exec()
		select_number, _ := strconv.Atoi(answers["select_number"])
		if select_number > 0 && select_number <= len(tasks) {
			//printTime(tasks[select_number-1])
			tasks[select_number-1].fn()
		} else {
			red.Println("The number is out of range, the process was killed. ")
		}
	}
}

func addTask(ts ...task) {
	for _, t := range ts {
		tasks = append(tasks, t)
	}
}
