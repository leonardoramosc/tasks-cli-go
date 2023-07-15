package cli

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"github.com/leonardoramosc/task-cli/pkg/user"
	"github.com/leonardoramosc/task-cli/pkg/task"
)

type menu struct {
	options []*option
}

type option struct {
	number     int
	title      string
	optionType any
}

func (o *option) handleSelection() {

	if m, ok := o.optionType.(*menu); ok {
		m.Display().GetUserSelection()
	}

	if t, ok := o.optionType.(string); ok {
		fmt.Println(t)
		return
	}

	f, ok := o.optionType.(func())

	if ok {
		f()
	}
}

func (m *menu) WithOption(opt option) *menu {
	o := m.FindOption(opt.number)
	if o != nil {
		fmt.Printf("Option with number %v already exist", opt.number)
	} else {
		m.options = append(m.options, &opt)
	}
	return m
}

func (m *menu) FindOption(optionNumber int) *option {
	for _, o := range m.options {
		if o.number == optionNumber {
			return o
		}
	}
	return nil
}

func (m *menu) Display() *menu {
	fmt.Println()
	for _, o := range m.options {
		fmt.Printf("%v: %v\n", o.number, o.title)
	}
	fmt.Println()
	return m
}

func (m *menu) GetUserSelection() {
	var option int

	fmt.Scanln(&option)

	o := m.FindOption(option)

	if o == nil {
		fmt.Println("Por favor selecciona una opción válida")
		m.Display().GetUserSelection()
	} else {
		o.handleSelection()
	}
}

func getLoginMenu() *menu {
	m := menu{}

	opt1 := option{number: 1, title: "Iniciar sesión", optionType: logUser}
	opt2 := option{number: 2, title: "Registrarme", optionType: createUser}

	m.WithOption(opt1).WithOption(opt2)

	return &m
}

func displayUserMenu(u *user.UserSchema) {
	um := menu{}
	um.WithOption(option{number: 1, title: "Ver mis tareas", optionType: listUserTasks(u)}).WithOption(option{number: 2, title: "Crear tarea", optionType: createUserTask(u)})

	um.Display().GetUserSelection()
}

func createUser() {
	fmt.Println("Ingresa tu nombre de usuario:")
	u := readUserInput()

	fmt.Println("Ingresa tu edad:")
	a := readUserInput()

	aint, e := strconv.ParseInt(a, 10, 8)

	if e != nil {
		fmt.Println("La edad debe ser un número")
		createUser()
	}

	fmt.Println("Ingresa tu contraseña")
	p := readUserInput()

	newUser, e := user.New(u, int8(aint), p)

	if e != nil {
		fmt.Println(e)
		createUser()
	}

	uc := user.GetUserCollection()

	us, e := uc.CreateUser(newUser)

	if e != nil {
		fmt.Println("No se pudo crear el usuario. razón:", e)
		return
	}

	displayUserMenu(us)
}

func logUser() {
	var username string
	var password string

	fmt.Println("Insertar nombre de usuario:")
	fmt.Scanln(&username)

	fmt.Println("Insertar contraseña:")
	fmt.Scanln(&password)

	uc := user.GetUserCollection()

	u, e := uc.GetByUsername(username)
	eMsg := "Usuario o contraseña invalidos"

	if e != nil || u.Password != password {
		fmt.Println(eMsg)
		logUser()
	} else {
		fmt.Printf("\nBienvenido %v. ¿Qué deseas hacer?\n", username)

		displayUserMenu(u)
	}
}

func listUserTasks(us *user.UserSchema) func() {
	return func() {
		tasks := us.ListTasks()

		if len(tasks) == 0 {
			fmt.Println("No tienes tareas creadas")
		}
		fmt.Println("Tienes las siguientes tareas pendientes:")
		for _, t := range tasks {
			fmt.Printf("- %v\n", t.Title)
		}
	}
}

func createUserTask(us *user.UserSchema) func() {
	return func() {
		fmt.Println("Inserta un título para tu tarea:")
		title := readUserInput()
	
		if len(title) < 3 {
			fmt.Println("El titulo de la tarea es muy corto. Debe contener al menos 3 caracteres")
		}
		t := &task.Task{Title: title, Completed: false}

		uc := user.GetUserCollection()
		uc.AppendTask(us.Username, t)
		fmt.Println("Tarea creada exitosamente")
		displayUserMenu(us)
	}
}

func Entrypoint() {
	uc := user.GetUserCollection()

	uc.Init().LoadData()

	fmt.Println("Bienvenido! qué deseas hacer?")
	getLoginMenu().Display().GetUserSelection()

	uc.Exec()
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	text, _, _ := reader.ReadLine()

	return string(text)
}

/*
{
	optionNumber: 1,
	optionDescription: "iniciar sesion",
	displaySelection
}
**/
