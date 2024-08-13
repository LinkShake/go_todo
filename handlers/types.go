package handlers

type Todo struct {
	ID uint
	Text string
}

type User struct {
	Email string
	Pwd string
}

type ReqFailed struct {
	Ok bool
	Msg string
}