package main

import "database/sql"

type tv struct {
	Id           int    `json:"id"`
	Brand        string `json:"brand"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
}

func (tv *tv) getTV(db *sql.DB) error {
	return db.QueryRow("SELECT brand, manufacturer, model, year  FROM tv WHERE id=?",
		tv.Id).Scan(&tv.Brand, &tv.Manufacturer, &tv.Model, &tv.Year)
}

func (tv *tv) updateTv(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE tv SET brand=?, manufacturer=?, model=?, year =? WHERE id=?",
			tv.Brand, tv.Manufacturer, tv.Model, tv.Year, tv.Id)

	return err
}

func (tv *tv) deleteTv(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tv WHERE id = ?", tv.Id)

	return err
}

func (tv *tv) createTv(db *sql.DB) error {
	_, err := db.Query(
		"INSERT INTO tv(brand, manufacturer, model, year) VALUES(?, ?, ?, ?)",
		tv.Brand, tv.Manufacturer, tv.Model, tv.Year)

	if err != nil {
		return err
	}
	return nil
}

func getAllTv(db *sql.DB) ([]tv, error) {
	rows, err := db.Query(
		"SELECT id, brand, manufacturer, model, year FROM tv")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tvs := []tv{}

	for rows.Next() {
		var tv tv
		if err := rows.Scan(&tv.Id, &tv.Brand, &tv.Manufacturer, &tv.Model, &tv.Year); err != nil {
			return nil, err
		}
		tvs = append(tvs, tv)
	}

	return tvs, nil
}
