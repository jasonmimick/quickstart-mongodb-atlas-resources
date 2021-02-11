package util

import (
	"github.com/Sectorbob/mlab-ns2/gae/ns/digest"
	"go.mongodb.org/atlas/mongodbatlas"
   logrus "github.com/sirupsen/logrus"
   "log"
	"strings"
    "os"
)

const (
	Version = "beta"
)

// This takes either "us-east-1" or "US_EAST_1"
// and returns "US_EAST_1" -- i.e. a valid Atlas region
func EnsureAtlasRegion(region string) string {
	r := strings.ToUpper(strings.Replace(string(region), "-", "_", -1))
    log.Printf("EnsureAtlasRegion--- region:%s r:%s", region,r)
	return r
}
// This takes either "us-east-1" or "US_EAST_1"
// and returns "us-east-1" -- i.e. a valid AWS region
func EnsureAWSRegion(region string) string {
	r := strings.ToLower(strings.Replace(string(region), "_", "-", -1))
    log.Printf("EnsureAWSRegion--- region:%s r:%s", region,r)
	return r
}

func CreateMongoDBClient(publicKey, privateKey string) (*mongodbatlas.Client, error) {
	// setup a transport to handle digest
	log.Printf("CreateMongoDBClient--- publicKey:%s", publicKey)
	logrus.Debugf("CreateMongoDBClient--- publicKey:%s", publicKey)
	transport := digest.NewTransport(publicKey, privateKey)

	// initialize the client
	client, err := transport.Client()
	if err != nil {
		return nil, err
	}

	//Initialize the MongoDB Atlas API Client.
	atlas := mongodbatlas.NewClient(client)
	atlas.UserAgent = "mongodbatlas-cloudformation-resources/" + Version
	return atlas, nil
}

// call this from each resource init() to setup
// the logger consistently
func InitLogger() {
  // Log as JSON instead of the default ASCII formatter.
  logrus.SetFormatter(&logrus.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  logrus.SetOutput(os.Stdout)


  // Only log the warning severity or above.
  logrus.SetLevel(logrus.DebugLevel)
  //log.SetLevel(log.WarnLevel)
}

