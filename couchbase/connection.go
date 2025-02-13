package couchbase

import (
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

type Client interface {
	InitCouchbase(username string, password string, bucketName string) error
	GetCollection() *gocb.Collection
	NewClient() *client
}
type client struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
}

func (c *client) NewClient() *client {
	return &client{
		Cluster: c.Cluster,
		Bucket:  c.Bucket,
	}
}

// InitCouchbase initializes the connection to the Couchbase cluster and bucket
func (c *client) InitCouchbase(username, password, bucketName string) error {
	var err error
	maxRetries := 5
	retryDelay := 10 * time.Second

	for i := 0; i < maxRetries; i++ {
		// Connect to the Couchbase Cluster
		c.Cluster, err = gocb.Connect("couchbase://couchbase", gocb.ClusterOptions{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Printf("Attempt %d: Failed to connect to Couchbase: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		// Connect to the specified Bucket
		c.Bucket = c.Cluster.Bucket(bucketName)
		if err = c.Bucket.WaitUntilReady(5*time.Second, nil); err != nil {
			log.Printf("Attempt %d: Failed to connect to bucket: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		log.Println("Connected to Couchbase successfully!")

	}

	log.Fatalf("Failed to connect after %d attempts", maxRetries)
	return err
}

// GetCollection returns the default collection for performing document operations
// This can be used for CRUD operations
func (c *client) GetCollection() *gocb.Collection {
	return c.Bucket.DefaultCollection()
}
