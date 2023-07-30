package seed

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var (
// 	Database string = "vanus-cloud-test"
// 	Username string = "vanus-cloud-test"
// 	Password string = "wz747Ibql4R7q4Cl"
// 	Address  string = "cluster0.ywakulp.mongodb.net"
// )

// var (
// 	Database string = "vanus-cloud-prod"
// 	Username string = "vanus-cloud-prod-rw"
// 	Password string = "K1Cr0WGQca396QLu"
// 	Address  string = "cluster1.odfrc.mongodb.net"
// )

type MongoDB struct {
	Database string
	Username string
	Password string
	Address  string
}

func Connect(ctx context.Context, db *MongoDB) (*mongo.Client, error) {
	var (
		err error
	)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", db.Username, db.Password, db.Address)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err = cli.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	fmt.Printf("Success to connect to MongoDB, address: %s\n", db.Address)

	if cli != nil {
		return cli, nil
	}
	return cli, nil
}
