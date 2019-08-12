package etcd

import (
	"testing"
	"github.com/coreos/etcd/clientv3"
)

func Test_Etcd(t *testing.T) {
	cli, _ := Cli()
	defer cli.Close()
	Put(cli, "a", "1")
	Put(cli, "key", "valddd")
	Put(cli, "key1", "val1")
	Put(cli, "key2", "val2")
	Put(cli, "jkey2", "jval2")
	Put(cli, "kfy", "val")

	opts := make([]clientv3.OpOption, 0)

	opts = append(opts, clientv3.WithLimit(20))
	opts = append(opts, clientv3.WithFromKey())
	opts = append(opts, clientv3.WithRange("key3"))
	Get(cli,"a", opts...)
}


func Test_Watch(t *testing.T) {
	cli, _ := Cli()
	defer cli.Close()
	Watch(cli, "key")
}