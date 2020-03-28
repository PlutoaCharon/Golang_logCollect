package tailf

import (
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"sync"
	"time"
)

// 定义常量
const (
	StatusNormal = 1
	StatusDelete = 2
)

// 将日志收集配置放在tailf包下,方便其他包引用
type CollectConf struct {
	LogPath string `json:"logpath"`
	Topic   string `json:"topic"`
}

// 存入Collect
type TailObj struct {
	tail     *tail.Tail
	conf     CollectConf
	status   int
	exitChan chan int
}

// 定义Message信息
type TextMsg struct {
	Msg   string
	Topic string
}

// 管理系统所有tail对象
type TailObjMgr struct {
	tailsObjs []*TailObj
	msgChan   chan *TextMsg
	lock      sync.Locker
}

// 定义全局变量
var (
	tailObjMgr *TailObjMgr
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

// 新增etcd配置项
func UpdateConfig(confs []CollectConf) (err error) {
	//// 加入锁, 防止多个goroutine工作
	//tailObjMgr.lock.Lock()
	//defer tailObjMgr.lock.Unlock()

	// 创建新的tailtask
	for _, oneConf := range confs {
		// 对于已经运行的所有实例, 路径是否一样
		var isRuning = false
		for _, obj := range tailObjMgr.tailsObjs {
			// 路径一样则证明是同一实例
			if oneConf.LogPath == obj.conf.LogPath {
				isRuning = true
				obj.status = StatusNormal
				break
			}
		}

		// 检查是否已经存在
		if isRuning {
			continue
		}

		// 如果不存在该配置项 新建一个tailtask任务
		createNewTask(oneConf)
	}

	// 遍历所有查看是否存在删除操作
	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailsObjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}
		// 如果status为删除, 则将exitChan置为1
		if obj.status == StatusDelete {
			obj.exitChan <- 1
		}
		// 将obj存入临时的数组中
		tailObjs = append(tailObjs, obj)
	}
	// 将临时数组传入tailsObjs中
	tailObjMgr.tailsObjs = tailObjs
	return
}

func createNewTask(conf CollectConf) {
	// 初始化Tailf实例
	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if errTail != nil {
		logs.Error("收集文件[%s]错误: %v", conf.LogPath, errTail)
		return
	}
	// 导入配置项
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}

	obj.tail = tails
	tailObjMgr.tailsObjs = append(tailObjMgr.tailsObjs, obj)

	go readFromTail(obj)
}

// 初始化tail
func InitTail(conf []CollectConf, chanSize int) (err error) {

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize), // 定义Chan管道
	}

	// 加载配置项
	if len(conf) == 0 {
		logs.Error("无效的日志collect配置: ", conf)
	}

	// 循环导入
	for _, v := range conf {
		createNewTask(v)
	}

	return
}

// 读入日志数据
func readFromTail(tailObj *TailObj) {
	for true {
		select {

		case msg, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("Tail file close reopen, filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{
				Msg:   msg.Text,
				Topic: tailObj.conf.Topic,
			}
			// 放入chan里
			tailObjMgr.msgChan <- textMsg

		// 如果exitChan为1, 则删除对应配置项
		case <-tailObj.exitChan:
			logs.Warn("tail obj 退出, 配置项为conf:%v", tailObj.conf)
			return
		}
	}
}
