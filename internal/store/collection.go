package store

import (
	"database/sql"
	"errors"
	"fmt"
)

type CollectionRepository struct {
	db *sql.DB
}

func NewCollectionRepository(db *sql.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}

func (r *CollectionRepository) Create(title string) error {
	newCollection := `INSERT INTO collection(title) VALUES(?);`
	result, err := r.db.Exec(newCollection, title)
	if err != nil {
		return fmt.Errorf("create collection %q: %w", title, err)
	}
	newId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("last insert id for collection: %w", err)
	}
	fmt.Println("New collection created. Collection ID:", newId)

	return nil
}

func (r *CollectionRepository) ReadAll() ([]Collection, error) {
	getCollections := `
	SELECT c.id, c.title, COUNT(l.id) as NumberOfLinks, c.created_at
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
		if err := rows.Scan(&c.ID, &c.Title, &c.LinkCount, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("read row of collection: %w", err)
		}
		collections = append(collections, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read rows of collection: %w", err)
	}

	return collections, nil
}

// idByTitle looks up a collection's id by title. Returns ErrNotFound if
// no collection with that title exists.
func (r *CollectionRepository) idByTitle(title string) (int64, error) {
	var id int64
	err := r.db.QueryRow(`SELECT id FROM collection WHERE title = ?`, title).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("collection %q: %w", title, ErrNotFound)
		}
		return 0, fmt.Errorf("looking up collection %q: %w", title, err)
	}
	return id, nil
}

func (r *CollectionRepository) Update(oldTitle, newTitle string) error {
	updateSql := `UPDATE collection SET title = ? WHERE title = ?;`
	result, err := r.db.Exec(updateSql, newTitle, oldTitle)
	if err != nil {
		return fmt.Errorf("update collection %q: %w", oldTitle, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("read affected rows for collection update: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("collection %q: %w", oldTitle, ErrNotFound)
	}
	return nil
}

func (r *CollectionRepository) Delete(title string) error {
	deleteSql := `DELETE FROM collection WHERE title = ?;`
	result, err := r.db.Exec(deleteSql, title)
	if err != nil {
		return fmt.Errorf("delete collection %q: %w", title, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("read affected rows for collection delete: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("collection %q: %w", title, ErrNotFound)
	}
	return nil
}