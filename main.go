package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	task "github.com/AgusMolinaCode/go-cli/tasks"
)

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() > 0 {

		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}

	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Nombre de la tarea: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(name, tasks)

		// Guardar las tareas actualizadas en el archivo
		file.Seek(0, 0)
		file.Truncate(0)
		bytes, err := json.Marshal(tasks)
		if err != nil {
			panic(err)
		}
		file.Write(bytes)
	case "complete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("ID de la tarea a completar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		tasks = task.CompleteTask(id, tasks)

		// Guardar las tareas actualizadas en el archivo
		file.Seek(0, 0)
		file.Truncate(0)
		bytes, err := json.Marshal(tasks)
		if err != nil {
			panic(err)
		}
		file.Write(bytes)
	case "delete":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("ID de la tarea a eliminar: ")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)

		tasks = task.DeleteTask(id, tasks)

		// Guardar las tareas actualizadas en el archivo
		file.Seek(0, 0)
		file.Truncate(0)
		bytes, err := json.Marshal(tasks)
		if err != nil {
			panic(err)
		}
		file.Write(bytes)
	}

}

func printUsage() {
	fmt.Println("Escoje un comando para ejecutar [add, list, complete, delete]")
}
