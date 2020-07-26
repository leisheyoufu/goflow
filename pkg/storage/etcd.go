package storage

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/clientv3"
)

func get(cli *clientv3.Client, key string) ([]byte, error) {
	var err error
	var resp *clientv3.GetResponse
	printFunc()
	if resp, err = cli.Get(context.TODO(), key); err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("Key %s: not found", key)
	}
	return resp.Kvs[0].Value, nil
}

func put(cli *clientv3.Client, key string, value string) error {
	var err error
	printFunc()
	if _, err = cli.Put(context.TODO(), key, value); err != nil {
		return err
	}
	return nil
}

func delete(cli *clientv3.Client, key string) error {
	var err error
	printFunc()
	if _, err = cli.Delete(context.TODO(), key); err != nil {
		return err
	}
	return nil
}
