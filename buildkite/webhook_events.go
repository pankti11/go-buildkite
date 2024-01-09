package buildkite

// agentEvent is a wrapper for an agent event notification
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/agent-events
type agentEvent struct {
	Event  *string `json:"event"`
	Agent  *Agent  `json:"agent"`
	Sender *User   `json:"sender"`
}

// AgentConnectedEvent is triggered when an agent has connected to the API
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/agent-events
type AgentConnectedEvent agentEvent

// AgentDisconnectedEvent is triggered when an agent has disconnected.
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/agent-events
type AgentDisconnectedEvent agentEvent

// AgentLostEvent is triggered when an agent has been marked as lost.
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/agent-events
type AgentLostEvent agentEvent

// AgentStoppedEvent is triggered when an agent has stopped.
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/agent-events
type AgentStoppedEvent agentEvent

// AgentStoppingEvent is triggered when an agent is stopping.
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/agent-events
type AgentStoppingEvent agentEvent

// buildEvent is a wrapper for a build event notification
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/build-events
type buildEvent struct {
	Event    *string   `json:"event"`
	Build    *Build    `json:"build"`
	Pipeline *Pipeline `json:"pipeline"`
	Sender   *User     `json:"sender"`
}

// BuildFailingEvent is triggered when a build enters a failing state
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/build-events
type BuildFailingEvent buildEvent

// BuildFinishedEvent is triggered when a build finishes
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/build-events
type BuildFinishedEvent buildEvent

// BuildRunningEvent is triggered when a build starts running
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/build-events
type BuildRunningEvent buildEvent

// BuildScheduledEvent is triggered when a build is scheduled
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/build-events
type BuildScheduledEvent buildEvent

// jobEvent is a wrapper for a job event notification
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/job-events
type jobEvent struct {
	Event    *string   `json:"event"`
	Build    *Build    `json:"build"`
	Job      *Job      `json:"job"`
	Pipeline *Pipeline `json:"pipeline"`
	Sender   *User     `json:"sender"`
}

// JobActivatedEvent is triggered when a job is activated
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/job-events
type JobActivatedEvent jobEvent

// JobFinishedEvent is triggered when a job is finished
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/job-events
type JobFinishedEvent jobEvent

// JobScheduledEvent is triggered when a job is scheduled
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/job-events
type JobScheduledEvent jobEvent

// JobStartedEvent is triggered when a job is started
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/job-events
type JobStartedEvent jobEvent

// PingEvent is triggered when a webhook notification setting is changed
//
// Buildkite API docs: https://buildkite.com/docs/apis/webhooks/ping-events
type PingEvent struct {
	Event        *string       `json:"event"`
	Organization *Organization `json:"organization"`
	Sender       *User         `json:"sender"`
}
