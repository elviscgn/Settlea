package ws

const (
	// wsserver actions
	PingAction       = "ping"
	CreateRoomAction = "create_room"
	JoinRoomAction   = "join_room"
	LeaveRoomAction  = "leave_room"
	UserJoinAction   = "user_join"
	UserLeaveAction  = "user_leave"

	// chat
	SendMessageAction = "send_message"

	// game
	StartGameAction = "start_game"
	SendGameAction  = "send_game"
)

type Message struct {
	Action   string `json:"action"`
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

// func (message *Message) encode() []byte {
// 	json, err := json.Marshal(message)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return json
// }
