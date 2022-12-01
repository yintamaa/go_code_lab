package raft

import "time"

const (
	ElectionTimeout  = time.Millisecond * 300 //选举超时时间
	HeartBeatTimeout = time.Millisecond * 150 // 心跳超时时间
)
