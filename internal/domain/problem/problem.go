package problem

// Prolem represents a problem that the users can post to have solved

import (
	"byte-bird/internal/domain/user"
	"time"
)

type Problem struct {
	ID          string
	OwnerID     string
	Author      Author
	Title       string
	Description string
	Timestamp   time.Time
	Tasks       []Task
	Solved      bool
}

type Author struct {
	// embed the user type for now, may not need to change
	user.User
}

// tasks are the division of the problem into smaller parts (duh but trying to keep my thhoughts straight)
type Task struct {
	ID              string
	ProblemID       string
	Author          Author
	Title           string
	Description     string
	ImportanceValue int
	Completed       bool
}

func NewProblem(ownerID string, author Author, title string, description string, tasks []Task) Problem {
	return Problem{
		OwnerID:     ownerID,
		Author:      author,
		Title:       title,
		Description: description,
		Tasks:       tasks,
		Timestamp:   time.Now(),
		Solved:      false,
	}
}

// users can add a task to an overall problem. This will be a way to break the problem down into smaller parts
func (p *Problem) AddTask(description string, importanceValue int) {
	task := Task{
		ProblemID: p.ID,
	}

	p.Tasks = append(p.Tasks, task)
}

func (p *Problem) RemoveTask(taskID string) {
	for i, task := range p.Tasks {
		if task.ID == taskID {
			p.Tasks = append(p.Tasks[:i], p.Tasks[i+1:]...)
		}
	}
}

func (p *Problem) MarkTaskComplete(taskID string) {
	for i, task := range p.Tasks {
		if task.ID == taskID {
			p.Tasks[i].Completed = true
		}
	}
}

// update functions for problems
func (p *Problem) UpdateTitle(title string) {
	p.Title = title
}

func (p *Problem) UpdateDescription(description string) {
  p.Description = description
}

func (p *Problem) UpdateAuthor(author Author) {
  p.Author = author
}

func (p *Problem) UpdateSolved(solved bool) {
  p.Solved = solved
}

func (p *Problem) UpdateTaskByID(taskID string, task Task) {
  for i, t := range p.Tasks {
    if t.ID == taskID {
      p.Tasks[i] = task
    }
  }
}

