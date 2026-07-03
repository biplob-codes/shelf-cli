package store

import (
	"database/sql"
	"errors"
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

func (r *Repository) AddLink(url, tag, collection string) error {
	var collectionID sql.NullInt64

	if collection != "" {
		var id int64
		err := r.db.QueryRow(`SELECT id FROM collection WHERE title = ?`, collection).Scan(&id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("collection %q does not exist", collection)
			}
			return fmt.Errorf("looking up collection %q: %w", collection, err)
		}
		collectionID = sql.NullInt64{Int64: id, Valid: true}
	}

	insertLink := `
	INSERT INTO link (url, tag, collection_id)
	VALUES (?, ?, ?);`
	result, err := r.db.Exec(insertLink, url, tag, collectionID)
	if err != nil {
		return fmt.Errorf("add link: %w", err)
	}

	newId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("last insert id for link: %w", err)
	}
	fmt.Println("Created link id:", newId)

	return nil
}

func (r *Repository) ReadCollections() ([]Collection, error) {
	getCollections := `
	SELECT c.id, c.title, COUNT(l.id) as NumberOfLinks,c.created_at
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
		var c Collection
		err := rows.Scan(&c.ID, &c.Title, &c.LinkCount, &c.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("read row of collection: %w", err)
		}
		collections = append(collections, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read rows of collection: %w", err)
	}

	return collections, nil
}
func (r *Repository) GetLinks(collection string) ([]Link, error) {
	var getLinks string
	if collection == "" {
		getLinks = `SELECT id,url,tag,created_at FROM link WHERE collection_id IS NULL ORDER BY id`
	} else {
		getLinks = `
	SELECT id,url,tag,created_at FROM link
	WHERE collection_id=(SELECT id FROM collection WHERE title=(?))
	ORDER BY id
	`
	}

	rows, err := r.db.Query(getLinks, collection)
	if err != nil {
		return nil, fmt.Errorf("read links: %w", err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var l Link
		err := rows.Scan(&l.ID, &l.URL, &l.Tag, &l.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("read row of link: %w", err)
		}
		links = append(links, l)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read rows of link: %w", err)
	}

	return links, nil
}

func (r *Repository) UpdateCollection(oldTitle, newTitle string) error {
	updateSql := `UPDATE collection SET title=(?) WHERE title=(?);`
	result, err := r.db.Exec(updateSql, newTitle, oldTitle)
	if err != nil {
		return fmt.Errorf("Update collection row: %w", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Read affected rows: %w", err)
	}
	fmt.Printf("%d row updated successfully in collection.", affectedRows)
	return nil
}
func (r *Repository) UpdateLink(id int, tag string) error {
	updateSql := `UPDATE link SET tag=(?) WHERE id=(?);`
	result, err := r.db.Exec(updateSql, tag, id)
	if err != nil {
		return fmt.Errorf("Update link row: %w", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Read affected rows: %w", err)
	}
	fmt.Printf("%d row updated successfully in link.", affectedRows)
	return nil
}
func (r *Repository) DeleteCollection(title string) error {
	deleteSql := `DELETE FROM collection WHERE title=(?);`
	result, err := r.db.Exec(deleteSql, title)
	if err != nil {
		return fmt.Errorf("Delete collection row: %w", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Read affected rows: %w", err)
	}
	fmt.Printf("%d row deleted successfully in collection.", affectedRows)
	return nil
}
func (r *Repository) DeleteLink(id int) error {
	deleteSql := `DELETE FROM link WHERE id=(?);`
	result, err := r.db.Exec(deleteSql, id)
	if err != nil {
		return fmt.Errorf("Delete link row: %w", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Read affected rows: %w", err)
	}
	fmt.Printf("%d row deleted successfully in link.", affectedRows)
	return nil
}
