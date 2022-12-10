package handling

import (
	"errors"
	"fmt"
	"strconv"
)

// 模拟枚举 enum 使用的一种新尝试
var (
	LOAD    = EventType{0, true}
	UNLOAD  = EventType{1, true}
	RECEIVE = EventType{2, false}
	CLAIM   = EventType{3, false}
	CUSTOMS = EventType{4, false}
)

var availabilityValues = []*EventType{
	&LOAD,
	&UNLOAD,
	&RECEIVE,
	&CLAIM,
	&CUSTOMS,
}

// EventType 是 enum
//
// 使用 struct 来 代替 type EventType int 是因为我们可以确保 类型值的安全性，不能随意赋值
// 另外，具有更强的灵活性，可以 添加 任意 需要的字段和新的值
type EventType struct {
	code                int  // 标识值
	needCarrierMovement bool // 是否需要 a carrier movement
}

func (t EventType) Code() int {
	return t.code
}

func (t EventType) NeedCarrierMovement() bool {
	return t.needCarrierMovement
}

func NewEventTypeFromCode(code int) (EventType, error) {
	for _, t := range availabilityValues {
		if t.code == code {
			return *t, nil
		}
	}

	return EventType{}, errors.New("unknown EventType code " + strconv.Itoa(code))
}

// IsZero 表示是否是零值，这里枚举业务上是不允许零值的
func (t EventType) IsZero() bool {
	return t == EventType{}
}

func (t EventType) String() string {
	return fmt.Sprintf("code: %d, needCarrierMovement:%s", t.Code(), t.NeedCarrierMovement())
}
