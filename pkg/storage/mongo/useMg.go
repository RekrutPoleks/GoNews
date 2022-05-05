package mongo

import (
	ctx "context"	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
	"GoNews/pkg/storage"
)

const (
	nameDB = "GoNews"
	nameColections = "posts"
	conn = "mongodb://185.251.91.177:27017/"
)

type dbStorage struct {
	db *mongo.Client 
	collectionGoNews *mongo.Collection
}

func New(connect string) (*dbStorage, error){

	mongoOPTS := options.Client().ApplyURI(connect)
	client, err := mongo.Connect(ctx.Background(), mongoOPTS)
	if err != nil{
		return nil, err
	}
	err = client.Ping(ctx.Background(), nil)
	if err != nil{
		return nil, err
	}
	return &dbStorage{client ,client.Database(nameDB).Collection(nameColections)}, nil
}

func (ds *dbStorage) AddPost(post storage.Post) error {
	_, err := ds.collectionGoNews.InsertOne(ctx.Background(), post)
	if err != nil {
		return err
	}
	return nil
}

func(ds *dbStorage) Posts() ([]storage.Post, error){
	filter := bson.D{}
	cur, err := ds.collectionGoNews.Find(ctx.Background(), filter)
	if err != nil{
		return nil, err
	}
	defer  cur.Close(ctx.Background())
	var data []storage.Post
	for cur.Next(ctx.Background()){
		var item storage.Post
		err := cur.Decode(&item)
		if err != nil{
			return nil, err
		}
		data = append(data, item)
	}
	return data, cur.Err()
}
func (ds dbStorage) UpdatePost(post storage.Post) error  {
	filter := bson.D{{"id", post.ID}}
	_, err := ds.collectionGoNews.UpdateOne(ctx.Background(),filter,post)
	if err != nil{
		return err
	}
	return nil	
}

func (ds dbStorage) DeletePost(post storage.Post) error{
	filter := bson.D{{"id", post.ID}}
	_, err := ds.collectionGoNews.DeleteOne(ctx.Background(),filter)
	if err != nil{
		return err
	}
	return nil
}