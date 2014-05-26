package message

import (
    "testing"
)

func TestInterface(t *testing.T) {
    abc := NewMessage(1002)
    abc.Serialize()
    abc.Deserialize("")
    abc.SetParameter("name", "yuki")
    abc.SetParameter("number", 2007081013)
    t.Log("abc.ParameterSize() ", abc.ParameterSize())
    t.Log(abc)
    t.Log(abc.Parameter("name"))
    t.Log(abc.Parameter("abc"))
}
