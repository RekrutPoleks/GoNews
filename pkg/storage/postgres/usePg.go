package postgres

import (
	"time"
	ctx "context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"GoNews/pkg/storage"
)

type author struct{
	id int
	name string
}


type dbStorage struct {
	db *pgxpool.Pool
}



func New(conn string ) (*dbStorage, error){
	db, err := pgxpool.Connect(ctx.Background(), conn)
	if err != nil{
		return nil, err
	}
	if err = db.Ping(ctx.Background()); err != nil{
		return nil, err
	}
	return &dbStorage{db: db}, nil
}

func (ds dbStorage) Posts() ([]storage.Post, error) {
	var post storage.Post
	var posts []storage.Post
	rows, err := ds.db.Query(
		ctx.Background(),
		QueryAllPosts,
		)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (ds *dbStorage) AddPost(post storage.Post) error{

	if !ds.AuthorExistById(post.AuthorID){
			return fmt.Errorf("Author not exist")
	}
	_, err := ds.db.Exec(
		ctx.Background(),
		InsertPost,
		post.AuthorID,
		post.Title,
		post.Content,
		time.Now().Unix(),
		)
	if err != nil{
		return err
	}
	return nil
}

func (ds dbStorage) CreateAuthor(name string) (*author, error) {
	var id int
	//Returning id author
	row:= ds.db.QueryRow(
		ctx.Background(),
		InsertAuthor,
		name,
	)
	if err := row.Scan(&id);err == nil {
		return &author{id, name}, err
	}else{
		return nil, err
	}
}

func (ds *dbStorage) AuthorById(id int)(*author, error){
	var a = author{}
	row := ds.db.QueryRow(
		ctx.Background(),
		QueryAuthorByID,
		id)
	err := row.Scan(
		&a.id,
		&a.name,
		)
	if err != nil{
		return nil, err
	}
	return &a, err
}

func (ds *dbStorage) AuthorExistById(id int) bool {
	var exist bool
	if _, err := ds.AuthorById(id); err != nil {
		exist = true
	}else{
		exist = false
	}
	return exist
}

func (ds dbStorage) UpdatePost(post storage.Post) error {
	_, err := ds.db.Exec(
		ctx.Background(),
		UpdatePost,		
		post.Content,
		time.Now().Unix(),
		post.ID
		)
	if err != nil{
		return err
	}
	return nil
}

func (ds *dbStorage) DeletePost(post storage.Post) error {
	_, err := ds.db.Exec(
		ctx.Background(),
		SQLDeletePost,
		post.ID)
	if err != nil {
		return err
	}
	return nil
}