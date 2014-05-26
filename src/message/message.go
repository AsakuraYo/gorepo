package message

import (
    "time"
)

//type Serializable interface {
//    Serialize() string
//    Deserialize(obj string) interface{}
//}

const (
    SM_SCHED_EVENT = 10
    SM_SCHED_TRIGGER = 11
    SM_TASK_START = 100
    SM_TASK_RETURN = 101
    SM_ROUTINE_RETURN = 102
    SM_FILE_FOUND = 200
    SM_USER_REFRESH = 301
    SM_USER_SHUTDOWN = 302
    SM_USER_PAUSE = 303
    SM_USER_RESUME = 304
    SM_USER_UPDATE = 305
    SM_USER_RESCHED = 306
    SM_USER_SETRET = 307
    SM_USER_CHECKPOINT = 308
    SM_USER_LISTRUNNING = 309
    SM_USER_UPDTINST = 310
    SM_USER_FORCE_RESCHED = 311
    SM_USER_REMOVE = 312
    SM_USER_ENABLE_ROUTINE = 313
    SM_USER_DISABLE_ROUTINE = 314
    SM_USER_LISTROUTINES = 315
    SM_TIMER = 400
    SM_CKPT_COTEAU = 500
    SM_HEART_BEAT = 501
    SM_WARNING = 502
    SM_INFO = 503
    SM_ERROR = 504
    SM_FATAL = 505
    SM_ROUTINE_CYCLE_START = 700
    SM_ROUTINE_CYCLE_END = 701
)

type Message struct {
    messageID   int
    messageDate int
    messageTime int
    messageSeq  int
    parameters  map[string] interface{}
}

func NewMessage(_id int) *Message {
    t       := time.Now()
    year    := int(t.Year())
    month   := int(t.Month())
    day     := int(t.Day())
    hour    := int(t.Hour())
    minute  := int(t.Minute())
    second  := int(t.Second())
    return &Message {
        messageID   : _id,
        messageDate : year * 10000 + month * 100 + day,
        messageTime : hour * 10000 + minute * 100 + second,
        messageSeq  : 0,
        parameters  : make(map[string] interface{}),
    }
}

func CopyMessage(msg Message) *Message {
    return &Message {
        messageID   : msg.messageID,
        messageDate : msg.messageDate,
        messageTime : msg.messageTime,
        messageSeq  : msg.messageSeq,
        parameters  : msg.parameters,
    }
}

func (m *Message) Serialize() string {
    return ""
}

func (m *Message) Deserialize(obj string) interface{} {
    return nil
}

func (m *Message) ParameterSize() int {
    return len(m.parameters)
}

func (m *Message) SetParameter(name string, value interface{}) {
    m.parameters[name] = value
}

func (m *Message) Parameter(name string) (value interface{}, ok bool) {
    val, ok := m.parameters[name]
    return val, ok
}
