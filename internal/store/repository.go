package store

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateCollection(title string) error {
	newCollection := `INSERT INTO collection(title) VALUES(?);`
	result, err := r.db.Exec(newCollection, title)
	if err != nil {
		return fmt.Errorf("create collection: %w", err)
	}
	newId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("last insert id for collection: %w", err)
	}
	fmt.Println("New collection created. Collection ID:", newId)

	return nil
}

type Collection struct {
	ID            int
	Title         string
	NumberOfLinks int
}

func (r *Repository) ReadCollections() ([]Collection, error) {
	getCollections := `
	SELECT c.id, c.title, COUNT(l.id) as NumberOfLinks
	FROM collection c
	LEFT JOIN link l ON l.collection_id = c.id
	GROUP BY c.id, c.title
	ORDER BY c.id
	`

	rows, err := r.db.Query(getCollections)
	if err != nil {
		return nil, fmt.Errorf("read collections: %w", err)
	}
	defer rows.Close()

	var collections []Collection
	for rows.Next() {
		var id int
		var title string
		var numberOfLinks int
		err := rows.Scan(&id, &title, &numberOfLinks)
		if err != nil {
			return nil, fmt.Errorf("read row of collection: %w", err)
		}
		collections = append(collections, Collection{ID: id, Title: title, NumberOfLinks: numberOfLinks})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read rows of collection: %w", err)
	}

	return collections, nil
}

func (r *Repository) UpdateCollection(oldTitle,newTitle string)error{
 updateSql:=`UPDATE collection SET title=(?) WHERE title=(?);`
 result,err:=r.db.Exec(updateSql,newTitle,oldTitle)
 if err!=nil{
	return fmt.Errorf("Update collection row: %w", err)
 }
 affectedRows,err:=result.RowsAffected()
 if err!=nil{
	return fmt.Errorf("Read affected rows: %w", err)
 }
 fmt.Printf("%d row updated successfully in collection.",affectedRows)
 return nil
}

func (r *Repository) DeleteCollection(title string) error{
	deleteSql:=`DELETE FROM collection WHERE title=(?);`
	result,err:=r.db.Exec(deleteSql,title)
	if err!=nil{
		return fmt.Errorf("Delete collection row: %w", err)
	}
 affectedRows,err:=result.RowsAffected()
 if err!=nil{
	return fmt.Errorf("Read affected rows: %w", err)
 }
 	fmt.Printf("%d row deleted successfully in collection.",affectedRows)
 return nil
}