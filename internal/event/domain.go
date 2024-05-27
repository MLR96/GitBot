package event

import (
	"io"
	"net/http"
)

type EventType int

const (
	// EventType
	EventTypeUnknown   EventType = -1
	EventTypeOpened    EventType = 0
	EventTypeUpdated   EventType = 1
	EventTypeDeclined  EventType = 2
	EventTypeMerged    EventType = 3
	EventTypeCommented EventType = 4
)

type Action int

const (
	ActionNothing Action = iota
	ActionLock
	ActionUnlock
)

type PullRequest struct {
	Id                int
	SourceBranch      string
	DestinationBranch string
	Approved          bool
	FilesChanged      []string
}

type Event struct {
	Type        EventType
	Repository  string
	Author      string
	Comment     string
	CommentId   int
	PullRequest PullRequest
}

type Provider interface {
	//Name() string
	ParseEvent(headers http.Header, body io.ReadCloser) (Event, error)
	WriteComment(repo string, prId int, parentId int, msg string) error
	//	RespondEvent(e GitEvent, msg string) error
}

type QueueItem struct {
	Event    Event
	Provider Provider
}

type Queue interface {
	Enqueue(item QueueItem)
	NextItem() *QueueItem
	Dequeue() *QueueItem
	Size() int
}

type SecurityRule struct {
	Repository   string
	FilePatterns []string
	Actions      []string
	Users        []string
}