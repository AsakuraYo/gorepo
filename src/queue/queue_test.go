package queue

import (
    "testing"
)

func TestNewMultiRoutineQueue(t *testing.T) {
    q := NewMultiRoutineQueue()
    if q == nil {
        t.Errorf("q is expected not nil\n")
    }
    var arr []interface{} = []interface{}{ 1, 2, 4 }
    q = NewMultiRoutineQueueWith(arr)
    if q == nil {
        t.Errorf("q is expected not nil\n")
    }
}

func TestEmpty(t *testing.T) {
    q := NewMultiRoutineQueue()
    if q.Empty() == false {
        t.Errorf("q is expected empty\n")
    }

    var arr []interface{} = []interface{}{ 1, 2, 4 }
    q = NewMultiRoutineQueueWith(arr)
    if q.Empty() {
        t.Errorf("q is expected not empty\n")
    }
}

func TestClear(t *testing.T) {
    q := NewMultiRoutineQueue()
    q.Clear()
    if q.Empty() == false {
        t.Errorf("q is expected empty\n")
    }

    var arr []interface{} = []interface{}{ 1, 2, 4 }
    q = NewMultiRoutineQueueWith(arr)
    q.Clear()
    if q.Empty() == false {
        t.Errorf("q is expected empty\n")
    }
}

func TestLen(t *testing.T) {
    q := NewMultiRoutineQueue()
    if q.Len() != 0 {
        t.Errorf("Len error %d, EXCEPT 0\n", q.Len())
    }

    var arr []interface{} = []interface{}{ 1, 2, 4 }
    q = NewMultiRoutineQueueWith(arr)
    if q.Len() != 3 {
        t.Errorf("Len error %d, EXCEPT 3\n", q.Len())
    }
    q.Push(16)
    if q.Len() != 4 {
        t.Errorf("Len error(%d), EXCEPT 4\n", q.Len())
    }
}

func TestPop(t *testing.T) {
    var arr []interface{} = []interface{}{ 1, 2, 4 }
    q := NewMultiRoutineQueueWith(arr)
    item := q.Pop()
    if item != 1 {
        t.Errorf("item = %d, EXCEPT %d\n", item, 1)
    }
    item = q.Pop()
    if item != 2 {
        t.Errorf("item = %d, EXCEPT %d\n", item, 2)
    }
    item = q.Pop()
    if item != 4 {
        t.Errorf("item = %d, EXCEPT %d\n", item, 4)
    }
    if q.Empty() == false {
        t.Errorf("not empty, EXCEPT empty\n")
    }
}

func BenchmarkPush(b *testing.B) {
    var arr []interface{} = []interface{}{ 1, 2, 4 }
    q := NewMultiRoutineQueueWith(arr)
    for i := 0; i < 100; i++ {
        q.Push(i)
        b.Log(q.Len())
    }
}
