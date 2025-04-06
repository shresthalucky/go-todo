package data

const (
	CreateTodoAction  = "create_todo"
	ListTodosAction   = "list_todos"
	GetTodoByIDAction = "get_todo_by_id"
	UpdateTodoAction  = "update_todo"
)

type BrokerRequest struct {
	Action      string      `json:"action"`
	TodoPayload TodoPayload `json:"todoPayload,omitempty"`
}

type TodoPayload struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
	ID    string `json:"id,omitempty"`
}

type LogEntry struct {
	Level       string `json:"level"`
	Message     string `json:"message"`
	ServiceName string `json:"serviceName"`
	RequestID   string `json:"requestId,omitempty"`
	Error       string `json:"error,omitempty"`
}
