// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ScheduleNew struct {
	Id *Id `json:"id,omitempty"`
	// An optional name of the Schedule. This string must not contain more than 100 characters.
	Name *string `json:"name,omitempty"`
	// The name of the of the queue to schedule the Task on. This string must not contain more than 100 characters.
	Queue *string `json:"queue,omitempty"`
	// An optional description of the Schedule. This string must not contain more than 500 characters.
	Description *string `json:"description,omitempty"`
	// A [cron expression](https://crontab.guru/examples.html) describing the
	// Schedule on which Tasks will run (UTC).
	// Note: execution n + 1 of a Task will not begin until execution n has
	// completed successfully.
	// You must pass either `cron` or `rrule` when creating a new Schedule.
	Cron *string `json:"cron,omitempty"`
	// An [iCal RRule expression](https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html) describing the Schedule on which Tasks will run (UTC). The time of Schedule creation will be used as the start of the recurrence interval (i.e. `DTSTART`). Note: execution n + 1 of a Task will not begin until execution n has completed successfully. You must pass either `cron` or `rrule` when creating a new Schedule.
	Rrule *string `json:"rrule,omitempty"`
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations) denoting the start of an RRULE schedule. Example: "2021-10-01T15:53:05Z". When not set, it will be set to the current time, and the first Task will be scheduled immediately. Ignored for `cron`-type Schedules.
	Dtstart *string `json:"dtstart,omitempty"`
	// If `true`, the Schedule will be paused immediately. If `false`, a paused Schedule will be resumed.
	Paused    *bool      `json:"paused,omitempty"`
	Request   *Request   `json:"request,omitempty"`
	CreatedAt *CreatedAt `json:"created_at,omitempty"`
}

type Error struct {
	// A human-readable message providing more details about the error(s).
	Message *string `json:"message,omitempty"`
	// If the error is parameter-specific, the parameter related to the error.
	Param *string `json:"param,omitempty"`
	// If multiple errors occured (e.g., with param validation), the list of errors that occured.
	Errors []*Error `json:"errors,omitempty"`
}

type Request struct {
	// The URL that the POST request will be sent to.
	// For localhost development, use something like ngrok to get a publicly
	// accessible URL for your local service. See https://docs.mergent.co for
	// more info.
	Url string `json:"url"`
	// The headers that will accompany any Task's HTTP request. For
	// example, you can use this to set Content-Type to "application/json"
	// or "application/octet-stream".
	Headers map[string]any `json:"headers,omitempty"`
	// The HTTP request body as a string.
	Body *string `json:"body,omitempty"`
}

type Schedule struct {
	Id *Id `json:"id,omitempty"`
	// An optional name of the Schedule. This string must not contain more than 100 characters.
	Name *string `json:"name,omitempty"`
	// The name of the of the queue to schedule the Task on. This string must not contain more than 100 characters.
	Queue *string `json:"queue,omitempty"`
	// An optional description of the Schedule. This string must not contain more than 500 characters.
	Description *string `json:"description,omitempty"`
	// A [cron expression](https://crontab.guru/examples.html) describing the
	// Schedule on which Tasks will run (UTC).
	// Note: execution n + 1 of a Task will not begin until execution n has
	// completed successfully.
	// You must pass either `cron` or `rrule` when creating a new Schedule.
	Cron *string `json:"cron,omitempty"`
	// An [iCal RRule expression](https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html) describing the Schedule on which Tasks will run (UTC). The time of Schedule creation will be used as the start of the recurrence interval (i.e. `DTSTART`). Note: execution n + 1 of a Task will not begin until execution n has completed successfully. You must pass either `cron` or `rrule` when creating a new Schedule.
	Rrule *string `json:"rrule,omitempty"`
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations) denoting the start of an RRULE schedule. Example: "2021-10-01T15:53:05Z". When not set, it will be set to the current time, and the first Task will be scheduled immediately. Ignored for `cron`-type Schedules.
	Dtstart *string `json:"dtstart,omitempty"`
	// If `true`, the Schedule will be paused immediately. If `false`, a paused Schedule will be resumed.
	Paused    *bool      `json:"paused,omitempty"`
	Request   *Request   `json:"request,omitempty"`
	CreatedAt *CreatedAt `json:"created_at,omitempty"`
}

type Task struct {
	Id *Id `json:"id,omitempty"`
	// An optional name of the Task. This string must not contain more than 100 characters.
	Name *string `json:"name,omitempty"`
	// The name of the of the Task queue. This string must not contain more than 100 characters.
	Queue *string `json:"queue,omitempty"`
	// The status of this Task.
	Status  *TaskStatus `json:"status,omitempty"`
	Request *Request    `json:"request,omitempty"`
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations) for when the Task is scheduled to be executed. Example: "2021-10-01T15:53:05Z". When not present, the Task will be scheduled for immediate execution.
	ScheduledFor *string `json:"scheduled_for,omitempty"`
	// A duration string containing numbers and a unit suffix of "s" for seconds, "m" for minutes, and "h" for hours. Examples: "5s"; "1.5h"; "2h45m" When both `delay` and `scheduled_for` are present, `delay` will be added to `scheduled_for`.
	Delay     *string    `json:"delay,omitempty"`
	CreatedAt *CreatedAt `json:"created_at,omitempty"`
}

type TaskNew struct {
	Id *Id `json:"id,omitempty"`
	// An optional name of the Task. This string must not contain more than 100 characters.
	Name *string `json:"name,omitempty"`
	// The name of the of the Task queue. This string must not contain more than 100 characters.
	Queue *string `json:"queue,omitempty"`
	// The status of this Task.
	Status  *TaskStatus `json:"status,omitempty"`
	Request *Request    `json:"request,omitempty"`
	// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations) for when the Task is scheduled to be executed. Example: "2021-10-01T15:53:05Z". When not present, the Task will be scheduled for immediate execution.
	ScheduledFor *string `json:"scheduled_for,omitempty"`
	// A duration string containing numbers and a unit suffix of "s" for seconds, "m" for minutes, and "h" for hours. Examples: "5s"; "1.5h"; "2h45m" When both `delay` and `scheduled_for` are present, `delay` will be added to `scheduled_for`.
	Delay     *string    `json:"delay,omitempty"`
	CreatedAt *CreatedAt `json:"created_at,omitempty"`
}

// The status of this Task.
type TaskStatus uint

const (
	TaskStatusQueued TaskStatus = iota + 1
	TaskStatusWorking
	TaskStatusSuccess
	TaskStatusFailure
)

func (t TaskStatus) String() string {
	switch t {
	default:
		return strconv.Itoa(int(t))
	case TaskStatusQueued:
		return "queued"
	case TaskStatusWorking:
		return "working"
	case TaskStatusSuccess:
		return "success"
	case TaskStatusFailure:
		return "failure"
	}
}

func (t TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.String())), nil
}

func (t *TaskStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "queued":
		value := TaskStatusQueued
		*t = value
	case "working":
		value := TaskStatusWorking
		*t = value
	case "success":
		value := TaskStatusSuccess
		*t = value
	case "failure":
		value := TaskStatusFailure
		*t = value
	}
	return nil
}

// The [ISO 8601 timestamp](https://en.wikipedia.org/wiki/ISO_8601#Combined_date_and_time_representations) representing when the object was created.
type CreatedAt = string

// A unique ID assigned upon creation.
type Id = string
