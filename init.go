package gnet

import (
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var DefMsgQueTimeout int = 180 //默认连接超时时间s

var waitAll = &WaitGroup{} //等待所有goroutine
var waitAllForLog sync.WaitGroup
var waitAllForRedis sync.WaitGroup

var stopForLog int32 //
var stop int32       //停止标志

//静态数据结构
type Statis struct {
	GoCount     int
	MsgqueCount int
	StartTime   time.Time
	LastPanic   int //panic发生时间戳
	PanicCount  int32
	PoolGoCount int32
}

var statis = &Statis{}

var gocount int32 //goroutine数量
var goid uint32  //goroutine Id
var DefLog *Log //日志

var msgqueId uint32 //消息队列id
var msgqueMapSync sync.Mutex
var msgqueMap = map[uint32]IMsgQue{}

type gMsg struct {
	c   chan struct{}
	msg *Message
	fun func(msgque IMsgQue) bool
}

var gmsgId uint16
var gmsgMapSync sync.Mutex
var gmsgArray = [65536]*gMsg{}

type WaitGroup struct {
	count int64
}

func (r *WaitGroup) Add(delta int) {
	atomic.AddInt64(&r.count, int64(delta))
}

func (r *WaitGroup) Done() {
	atomic.AddInt64(&r.count, -1)
}

func (r *WaitGroup) Wait() {
	for atomic.LoadInt64(&r.count) > 0 {
		Sleep(1)
	}
}

func (r *WaitGroup) TryWait() bool {
	return atomic.LoadInt64(&r.count) == 0
}

var atexitId uint32
var atexitMapSync sync.Mutex
var atexitMap = map[uint32]func(){}

var stopChanForGo = make(chan struct{})//携程stop chanel
var stopChanForLog = make(chan struct{})
var stopChanForSys = make(chan os.Signal, 1)

var poolChan = make(chan func())
var poolGoCount int32  //go携程计数器

var StartTick int64
var NowTick int64
var Timestamp int64   // 当前秒数
var TimeString string // 当前时间 格式：2020-7-9 14:59:15

//全局配置
var Config = struct {
	AutoCompressLen uint32
	UdpServerGoCnt  int    //UDP
	PoolSize        int32  //携程池默认大小
	SSLCrtPath      string //crt路径
	SSLKeyPath      string //Key路径
	EnableWss       bool   //是否启用wss
	ReadDataBuffer  int    //读取数据大小
	StopTimeout     int    //超时
}{UdpServerGoCnt: 64, PoolSize: 50000, ReadDataBuffer: 1 << 12, StopTimeout: 3000} //默认配置

var stopCheckIndex uint64
var stopCheckMap = struct {
	sync.Mutex
	M map[uint64]string
}{M: map[uint64]string{}}

func init() {
	gmsgArray[gmsgId] = &gMsg{c: make(chan struct{})}
	runtime.GOMAXPROCS(runtime.NumCPU())
	DefLog = NewLog(10000, &ConsoleLogger{true})
	DefLog.SetLevel(LogLevelInfo)
	timerTick()
}
