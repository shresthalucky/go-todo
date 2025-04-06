package data

const (
	CreateTodoAction  = "create_todo"
	ListTodosAction   = "list_todos"
	GetTodoByIDAction = "get_todo_by_id"
	UpdateTodoAction  = "update_todo"
)

type BrokerRequest struct {
	Action      string      `json:"action"`
	TodoPayload TodoPayload `json:"todoPayload"`
}

type TodoPayload struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
	ID    string `json:"id,omitempty"`
}
