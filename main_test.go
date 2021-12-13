package main

import (
	"testing"
)

func TestValidStatus(t *testing.T) {

	got := validStatus(PENDING_STATUS)
	want := true

	if got != want {
		t.Errorf("Expect pass status: " + PENDING_STATUS)
	}

	got = validStatus(FINISHED_STATUS)
	want = true

	if got != want {
		t.Errorf("Expect pass status: " + FINISHED_STATUS)
	}

	got = validStatus("")
	want = false

	if got != want {
		t.Errorf("Expect not pass ''")
	}
}

func TestUpdateTask(t *testing.T) {
	var task Task
	task.Status = FINISHED_STATUS
	task.TimesFinished = 0

	got := updateTaskStatus(task, PENDING_STATUS)

	task.Status = PENDING_STATUS
	task.TimesFinished = 1
	want := task

	if got != want {
		t.Errorf("Expect update status and increment finished")
	}

}
