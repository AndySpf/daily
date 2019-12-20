package main

import (
	"fmt"
	"go.etcd.io/bbolt"
	"reflect"
)

// bbolt key-value数据库试用
func main() {
	db, err := bbolt.Open("/Users/qijing.fqj/go/src/daily/bbolt/metadata.db", 0666, nil)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		//
		//if err = tx.DeleteBucket([]byte("MyBucket")); err != nil {
		//	return err
		//}

		return nil
	})

	db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})

	db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("noexists"))
		fmt.Println(reflect.DeepEqual(v, nil)) // false
		fmt.Println(v == nil)                  // true

		v = b.Get([]byte("zero"))
		fmt.Println(reflect.DeepEqual(v, nil)) // false
		fmt.Println(v == nil)                  // true
		return nil
	})

	db.View(func(tx *bbolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("MyBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}
