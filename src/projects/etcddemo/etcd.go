package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopractice/common"
	"time"
)

var client *clientv3.Client
var dir string = "/test/node1"
var leaseDir string = "/test/lease/node1"
var ttl int64 = 10

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		fmt.Printf("new etcd client fail. | err: %s\n", err)
		return
	}

	user := &User{
		Name: "user_name",
		Age:  20,
	}

	data, _ := json.Marshal(user)

	// 添加节点
	_ = AddNode(dir, string(data))

	// 获取节点
	_ = GetNode(dir)

	// 添加租约节点
	_ = AddNodeWithLease(string(data))

	// 删除节点
	_ = DeleteNode(dir)

	// 事务
	Tnx(string(data))

	common.Break()

}

func AddNode(key, value string) error {
	rsp, err := client.Put(context.Background(), key, value)

	if err != nil {
		fmt.Printf("put etcd node fail. | err: %s\n", err)
		return err
	}

	fmt.Printf("rsp: %+v\n", *rsp)
	return nil
}

func GetNode(key string) error {
	rsp, err := client.Get(context.Background(), key)

	if err != nil {
		fmt.Printf("put etcd node fail. | err: %s\n", err)
		return err
	}

	user := &User{}

	_ = json.Unmarshal(rsp.Kvs[0].Value, user)

	fmt.Printf("user: %+v\n", *user)
	return nil
}

func GetRangeNode(key string, endKey string) error {
	rsp, err := client.Get(context.Background(), key, clientv3.WithRange(endKey), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))

	if err != nil {
		fmt.Printf("put etcd node fail. | err: %s\n", err)
		return err
	}

	if len(rsp.Kvs) > 0 {
		for _, item := range rsp.Kvs {
			fmt.Printf("k: %s | v: %s\n", string(item.Key), string(item.Value))
		}
	}
	return nil
}

func DeleteNode(key string) error {
	_, err := client.Delete(context.Background(), key)

	if err != nil {
		fmt.Printf("DeleteNode delete node fail. | err: %s\n", err)
		return err
	}

	return nil
}

func AddNodeWithLease(value string) error {
	rsp, err := client.Grant(context.Background(), ttl)

	if err != nil {
		fmt.Printf("AddNodeWithLease Grant fail. | err: %s\n", err)
		return err
	}

	opts := clientv3.WithLease(rsp.ID)

	_, err = client.Put(context.Background(), leaseDir, value, opts)

	if err != nil {
		fmt.Printf("AddNodeWithLease Put fail. | err: %s\n", err)
		return err
	}

	ctx := context.TODO()
	ka, err := client.KeepAlive(ctx, rsp.ID)

	if err != nil {
		fmt.Printf("AddNodeWithLease KeepAlive fail. | err: %s\n", err)
		return err
	}

	go func() {
		for {
			select {
			case kaRsp := <-ka:
				if kaRsp != nil {
					fmt.Printf("keep alive lease continue.| id: %d | ttl: %d | time: %s\n", kaRsp.ID, kaRsp.TTL, time.Now())
				} else {
					fmt.Printf("keep alive lease continue. rsp nil\n")
					return
				}
			case <-ctx.Done():
				fmt.Printf("keep alive done\n")
				return
			}
		}
	}()

	wt := client.Watch(context.Background(), leaseDir)

	go func() {
		for {
			select {
			case wtRsp := <-wt:
				for _, event := range wtRsp.Events {
					if string(event.Kv.Key) == leaseDir && event.Type == mvccpb.DELETE {
						client.Revoke(context.Background(), rsp.ID)
						fmt.Printf("revoke id: %d", rsp.ID)
						AddNodeWithLease(value)
					}
				}
			}
		}
	}()

	return nil
}

func Tnx(value string) {
	//通过key的Create_Revision 是否为 0 来判断key是否存在。其中If，Then 以及 Else 分支都可以包含多个操作。
	//返回的数据包含一个successed字段，当为 true 时代表 If 为真
	txnRsp, err := client.Txn(context.Background()).If(clientv3.Compare(clientv3.CreateRevision(dir), "=", 0)).
		Then(clientv3.OpPut(dir, value)).Commit()

	if err != nil {
		fmt.Printf("txn fail. | err: %s\n", err)
		return
	}

	fmt.Printf("txnRsp: %t\n", txnRsp.Succeeded)
}
