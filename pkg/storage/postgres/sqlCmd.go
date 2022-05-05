package postgres

const (
	InsertPost =	`INSERT INTO posts (author_id, title, content, created_at) 
 						 VALUES ($1, $2, $3, $4)`

	QueryAuthorByID = "SELECT id, name FROM authors WHERE id = %1"

	InsertAuthor = "INSERT INTO authors(name) VALUES (%1) RETURNING id;" //returning id

	UpdatePost = "UPDATE posts SET content = %1, created_at = %2 WHERE id = %3;" // content, publisher

	SQLDeletePost = "DELETE FROM posts WHERE id = %1"

	QueryAllPosts =	`SELECT p.id, p.title, content, a.id, a.name, p.created_at
					FROM posts as p , authors as a
					WHERE p.author_id = a.id;`
					//порядок столбцов
					// &post.ID,
					//&post.Title,
					//&post.Content,
					//&post.AuthorID,
					//&post.AuthorName,
					//&post.CreatedAt,


	UpdateTaskByID = `UPDATE tasks SET content = $1 where id=$2;`

	DeleteTaskByID = `DELETE FROM tasks WHERE id = $1;`
)