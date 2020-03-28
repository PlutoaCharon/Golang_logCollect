package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"logCollect/logAgent/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string) (collectConf []tailf.CollectConf, err error) {
	// 初始化连接etcd
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("连接etcd失败:", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
	}

	// 如果Key不是以"/"结尾, 则自动加上"/"
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	for _, ip := range localIPArray {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("etcd get请求失败:", err)
			continue
		}
		cancel()
		logs.Debug("resp from etcd:%v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				// 反序列化为结构体
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("反序列化失败:", err)
					continue
				}
				logs.Debug("日志设置为%v", collectConf)
			}
		}
	}
	initEtcdWatcher(logConfig.etcdAddr)
	logs.Debug("连接etcd成功")
	return
}

// 初始化多个watch监控etcd中配置节点
func initEtcdWatcher(addr string) {
	for _, key := range etcdClient.keys {
		go watchKey(addr, key)
	}
}

func watchKey(addr string, key string) {

	// 初始化连接etcd
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("连接etcd失败:", err)
		return
	}

	logs.Debug("开始监控key:", key)

	// Watch操作
	for {
		var collectConf []tailf.CollectConf
		var getConfSucc = true
		wch := cli.Watch(context.Background(), key)
		for resp := range wch {
			for _, ev := range resp.Events {
				// DELETE处理
				if ev.Type == mvccpb.DELETE {
					logs.Warn("删除Key[%s]配置", key)
					continue
				}
				// PUT处理
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("反序列化key[%s]失败:", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd ,Type: %v, Key:%v, Value:%v\n", ev.Type, string(ev.Kv.Key), string(ev.Kv.Value))
			}
			if getConfSucc {
				logs.Debug("get config from etcd success, %v", collectConf)
				_ = tailf.UpdateConfig(collectConf)
			}
		}
	}
}
