package domain

import "time"

type Task struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompleteAt  *time.Time

	AuthorID int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorUserID int,
) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   createdAt,
		CompleteAt:  completedAt,
		AuthorID:    authorUserID,
	}
}

func NewTaskUninitialized(
	title string,
	description *string,
	authorUserID int,
) Task {
	return NewTask(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		authorUserID,
	)
}
