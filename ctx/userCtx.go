package ctx

type UserContext struct {
	UserId string
}

var UserCtx UserContext = UserContext{}