package main

import (
	"fmt"
	"github.com/leonardoramosc/task-cli/pkg/cli"
	// "github.com/leonardoramosc/task-cli/pkg/user"
	// "github.com/leonardoramosc/task-cli/pkg/task"
)

func main() {
	fmt.Println("Init app")

	cli.Entrypoint()
}

// func createUser() {
// 	var username, password string
// 	var age int8

// 	fmt.Println("Ingrese username:")
// 	fmt.Scanln(&username)

// 	fmt.Println("Ingrese contraseña:")
// 	fmt.Scanln(&password)

// 	fmt.Println("Ingrese edad:")
// 	fmt.Scanln(&age)

// 	u, e := user.CreateUser(username, age, password)
// 	// t := &task.Task{Title: "estudiar", Completed: false}

// 	if e != nil {
// 		fmt.Println("No fue posible crear el usuario: ", e)
// 		return
// 	}

// 	fmt.Println("usuario creado:", u)
// }
