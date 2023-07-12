package main

import (
	"fmt"
	"github.com/leonardoramosc/task-cli/pkg/user"
	"github.com/leonardoramosc/task-cli/pkg/task"
)

func main() {
	fmt.Println("Init app")

	u, e := user.CreateUser("leorrc", 26, "12345")
	t := &task.Task{Title: "estudiar", Completed: false}

	if e != nil {
		fmt.Println("No fue posible crear el usuario: ", e)
	}

	uc := user.UserCollection{}

	uc.Init().LoadData()

	uc.CreateUser(u)
	uc.AppendTask(u.Username, t)

	uc.Exec()
}