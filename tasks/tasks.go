package task

import "fmt"

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

// ListTasks prints all tasks
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas para mostrar")
		return
	}

	for _, task := range tasks {
		if task.Completed {
			fmt.Println("[✓]", task.Name, "ID:", task.ID)
		} else {
			fmt.Println("[ ]", task.Name, "ID:", task.ID)
		}
	}
}

// AddTask adds a task to the list of tasks
func AddTask(name string, tasks []Task) []Task {
	tasks = append(tasks, Task{
		ID:        len(tasks) + 1,
		Name:      name,
		Completed: false,
	})

	return tasks
}

// CompleteTask toggles the completion status of a task
func CompleteTask(id string, tasks []Task) []Task {
	for i, task := range tasks {
		if fmt.Sprint(task.ID) == id {
			tasks[i].Completed = !tasks[i].Completed
			if tasks[i].Completed {
				fmt.Printf("Ahora tu tarea ID %s está completa\n", id)
			} else {
				fmt.Printf("Ahora tu tarea ID %s está incompleta\n", id)
			}
			break
		}
	}

	return tasks
}

// DeleteTask deletes a task from the list of tasks
func DeleteTask(id string, tasks []Task) []Task {
	for i, task := range tasks {
		if fmt.Sprint(task.ID) == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Tarea ID %s eliminada\n", id)
			break
		}
	}

	return tasks
}
