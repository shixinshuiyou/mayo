package mission

import "time"

// 任务活动处理
type MissionProcess interface {
	InitMission()
	DoMission()
	PresentProgress()
}

// 任务状态
type MissionStates struct {
}

// 任务实例
type Mission struct {
	TimeLimit time.Duration // 时间限制
	NumLimit  int32         // 数量限制
}
