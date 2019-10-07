package repo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"gitlab.ghn.vn/vinhnlk/gowithmongodb/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*model.Hero {
//     var heroes []*model.Hero
//     collection := client.Database("civilact").Collection("heroes")
//     // connect database and collection
//     cur, err := collection.Find(context.TODO(), filter)
//     if err != nil {
//         log.Fatal("Error on Finding all the documents", err)
//     }
//     for cur.Next(context.TODO()) {
//         var hero model.Hero
//         err = cur.Decode(&hero)
//         if err != nil {
//             log.Fatal("Error on Decoding the document", err)
//         }
//         heroes = append(heroes, &hero)
//     }
//     return heroes
// }
func InsertOneHeroes(hero model.Hero) interface{} {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//set client options
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("civilact").Collection("heroes")
	insertHeroes, err := collection.InsertOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	id := insertHeroes.InsertedID
	return id
}

func RemoveOneHeroes(hero model.Hero) interface{} {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//set client options
	fmt.Println(hero)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("civilact").Collection("heroes")
	removeHeroes, err := collection.DeleteOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on remove Hero", err)
	}
	return removeHeroes
}
func RemoveOneHeroesFilter(hero bson.M) interface{} {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//set client options
	fmt.Println(hero)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("civilact").Collection("heroes")
	removeHeroes, err := collection.DeleteOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on remove Hero", err)
	}
	return removeHeroes
}

// func updateOneHeroes(client *mongo.Client,filter bson.M,filterReplace bson.M) interface{}{
//     collection := client.Database("civilact").Collection("heroes")
//     atualizacao := bson.D{ {Key: "$set", Value: filterReplace} }
//     updateHeroes,err:= collection.UpdateOne(context.TODO(),filter,atualizacao)
//     if err != nil {
//         log.Fatalln("Error on update Hero", err)
//     }
//     return updateHeroes
// }
// func main() {
//     clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//     //set client options
//     client, err := mongo.Connect(context.TODO(), clientOptions)
//     //connect to mongodb
//     if err != nil {
//         log.Fatal(err)
//     }
//     err = client.Ping(context.TODO(), nil)
//     //check the connection
//     if err != nil {
//         log.Fatal(err)
//     }
//    fmt.Println("connect success")
//    hero := model.Hero{Name: "Stephen Strange", Alias: "Doctor Strange", Signed: true}
// //    removeHero:=Hero{Name: "Vision", Alias: "Vision", Signed: true}
//    insertOneHeroes(client,hero)// chen 1 dien vien
//    removeOneHeroes(client,bson.M{"alias": "Doctor Strange"}) // xoa 1 dien vien
//    updateOneHeroes(client,bson.M{"alias": "Doctor Strange"},bson.M{"alias": "Captain America"}) // cap nhat 1 dien vien
//    heroes := ReturnAllHeroes(client, bson.M{})
//     for _, hero := range heroes {
//     log.Println(hero.Name, hero.Alias, hero.Signed) // in ra database khi cap nhaat
//     }
// }
