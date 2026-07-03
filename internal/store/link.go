package store

import (
	"database/sql"
	"fmt"
)

type LinkRepository struct {
	db          *sql.DB
	collections *CollectionRepository
}

func NewLinkRepository(db *sql.DB, collections *CollectionRepository) *LinkRepository {
	return &LinkRepository{db: db, collections: collections}
}

func (r *LinkRepository) Add(url, tag, collection string) error {
	var collectionID sql.NullInt64

	if collection != "" {
		id, err := r.collections.idByTitle(collection)
		if err != nil {
			return fmt.Errorf("add link: %w", err)
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

func (r *LinkRepository) GetAll(collection string) ([]Link, error) {
	var getLinks string
	var args []any

	if collection == "" {
		getLinks = `SELECT id, url, tag, created_at FROM link WHERE collection_id IS NULL ORDER BY id`
	} else {
		id, err := r.collections.idByTitle(collection)
		if err != nil {
			return nil, fmt.Errorf("list links: %w", err)
		}
		getLinks = `SELECT id, url, tag, created_at FROM link WHERE collection_id = ? ORDER BY id`
		args = append(args, id)
	}

	rows, err := r.db.Query(getLinks, args...)
	if err != nil {
		return nil, fmt.Errorf("read links: %w", err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var l Link
		if err := rows.Scan(&l.ID, &l.URL, &l.Tag, &l.CreatedAt); err != nil {
			return nil, fmt.Errorf("read row of link: %w", err)
		}
		links = append(links, l)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read rows of link: %w", err)
	}

	return links, nil
}

func (r *LinkRepository) Update(id int, tag string) error {
	updateSql := `UPDATE link SET tag = ? WHERE id = ?;`
	result, err := r.db.Exec(updateSql, tag, id)
	if err != nil {
		return fmt.Errorf("update link %d: %w", id, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("read affected rows for link update: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("link %d: %w", id, ErrNotFound)
	}
	return nil
}

func (r *LinkRepository) Delete(id int) error {
	deleteSql := `DELETE FROM link WHERE id = ?;`
	result, err := r.db.Exec(deleteSql, id)
	if err != nil {
		return fmt.Errorf("delete link %d: %w", id, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("read affected rows for link delete: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("link %d: %w", id, ErrNotFound)
	}
	return nil
}