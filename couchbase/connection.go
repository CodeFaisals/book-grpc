package couchbase

import (
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

var (
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
)

// InitCouchbase initializes the connection to the Couchbase cluster and bucket
func InitCouchbase(username, password, bucketName string) {
	var err error
	maxRetries := 5
	retryDelay := 10 * time.Second

	for i := 0; i < maxRetries; i++ {
		// Connect to the Couchbase Cluster
		Cluster, err = gocb.Connect("couchbase://couchbase", gocb.ClusterOptions{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Printf("Attempt %d: Failed to connect to Couchbase: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		// Connect to the specified Bucket
		Bucket = Cluster.Bucket(bucketName)
		if err = Bucket.WaitUntilReady(5*time.Second, nil); err != nil {
			log.Printf("Attempt %d: Failed to connect to bucket: %v", i+1, err)
			time.Sleep(retryDelay)
			continue
		}

		log.Println("Connected to Couchbase successfully!")
		return
	}

	log.Fatalf("Failed to connect after %d attempts", maxRetries)
}

// GetCollection returns the default collection for performing document operations
// This can be used for CRUD operations
func GetCollection() *gocb.Collection {
	return Bucket.DefaultCollection()
}
