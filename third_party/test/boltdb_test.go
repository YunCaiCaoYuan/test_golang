package test

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"testing"
)

func Test_create(t *testing.T) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func Test_example(t *testing.T) {
	// step1：新建数据库DB
	db, err := bolt.Open("./test.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 往db里面插入数据
	err = db.Update(func(tx *bolt.Tx) error {
		// step2：新建桶Bucket
		bucket, err := tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			log.Fatalf("CreateBucketIfNotExists err:%s", err.Error())
			return err
		}
		//step3：存入key/value数据【数据的存储有一定的规则】
		if err = bucket.Put([]byte("hello"), []byte("world")); err != nil {
			log.Fatalf("bucket Put err:%s", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("db.Update err:%s", err.Error())
	}

	// step1：指定数据库DB【一定存在】
	err = db.View(func(tx *bolt.Tx) error {
		// step2：找到桶Bucket
		bucket := tx.Bucket([]byte("user"))
		//step3：根据key找到value数据
		val := bucket.Get([]byte("hello"))
		log.Printf("the get val:%s", val)
		val = bucket.Get([]byte("hello2"))
		log.Printf("the get val2:%s", val)
		return nil
	})
	if err != nil {
		log.Fatalf("db.View err:%s", err.Error())
	}
}
