package couchbase

import (
	"errors"
	//"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

//type Client interface {
//	InitCouchbase(username string, password string, bucketName string) error
//	GetCollection() *gocb.Collection
//	EssClient() *client
//	//NewClient() *client
//}
//type client struct {
//	Cluster *gocb.Cluster
//	Bucket  *gocb.Bucket
//}
//
//func (c *client) EssClient() *client {
//	return &client{}
//}

// InitCouchbase initializes the connection to the Couchbase cluster and bucket
func InitCouchbase(username, password, bucketName string) (*gocb.Cluster, error) {
	//var err error
	maxRetries := 5
	retryDelay := 10 * time.Second

	for i := 0; i < maxRetries; i++ {
		// Connect to the Couchbase Cluster
		cluster, err := gocb.Connect("couchbase://couchbase", gocb.ClusterOptions{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Printf("Attempt %d: Failed to connect to Couchbase: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		// Connect to the specified Bucket
		bucket := cluster.Bucket(bucketName)
		if err = bucket.WaitUntilReady(5*time.Second, nil); err != nil {
			log.Printf("Attempt %d: Failed to connect to bucket: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}
		log.Println("Connected to Couchbase successfully!")
		return cluster, nil
	}

	return nil, errors.New("Failed to connect to Couchbase")
}

// GetCollection returns the default collection for performing document operations
// This can be used for CRUD operations
//func GetCollection() *gocb.Collection {
//	return gocb.Cluster.Bucket().DefaultCollection()
//}
