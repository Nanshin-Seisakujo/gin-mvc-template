package db

import (
	"fmt"
)

/*
	global db object.

	var db *mongo.Client
*/

func Init() {
	/*
		some db connection settings.

		ex. mongo db

		uri := os.Getenv("DB_URI")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}

		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic(err)
		}

		db = client

	*/
	fmt.Println("Successfully db connected.")
}
