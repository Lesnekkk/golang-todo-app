package domain

import "time"

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreated int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}

func CreateStatistics(tasks []Task) Statistics {
	if len(tasks) == 0 {
		return NewStatistics(0, 0, nil, nil)
	}

	tasksCreated := len(tasks)
	tasksCompleted := 0
	var totalDuration time.Duration

	for _, task := range tasks {
		if task.Completed {
			tasksCompleted++
		}
		if d := task.CompletionDuration(); d != nil {
			totalDuration += *d
		}
	}

	rate := float64(tasksCompleted) / float64(tasksCreated) * 100

	var avgTime *time.Duration
	if tasksCompleted > 0 && totalDuration != 0 {
		avg := totalDuration / time.Duration(tasksCompleted)
		avgTime = &avg
	}

	return NewStatistics(tasksCreated, tasksCompleted, &rate, avgTime)
}
