package etcd

import (
	"time"
	"github.com/coreos/etcd/clientv3"
	"fmt"
	"encoding/json"
	"context"
)

func main() {

}
func Cli() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println(err)
	} else {
		fmt.Println("成功")
	}
	return cli, err

}

func Put(cli *clientv3.Client, key, val string, opts ...clientv3.OpOption) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := cli.Put(ctx, key, val, opts...)

	if err != nil {
		// handle error!
		fmt.Println(err)
		return err
	} else {
		fmt.Println("put success")
	}

	printJson(resp.Header)
	return nil
}

func Get(cli *clientv3.Client, key string, opts ...clientv3.OpOption) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := cli.Get(ctx, key, opts...)

	if err != nil {
		// handle error!
		fmt.Println(err)
		return err
	} else {
		fmt.Println("put success")
	}

	printJson(resp.Header)

	for _, item := range resp.Kvs {
		fmt.Println(string(item.Key), string(item.Value))
	}
	return nil
}

func Watch(cli *clientv3.Client, key string, opts ...clientv3.OpOption) {
	ctx := context.Background()
	resChan := cli.Watch(ctx, key)

	for {
		select {
		case val, ok := <-resChan:
			if !ok {
				fmt.Println("已取消")
				return
			}
			for _, item := range val.Events {
				fmt.Println("type:", item.Type, "key:", string(item.Kv.Key), "val:", string(item.Kv.Value))
			}
		}
	}
}

func printJson(mo interface{}) {
	by, _ := json.Marshal(mo)
	fmt.Println(string(by))
}
