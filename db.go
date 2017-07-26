package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var database string = "./cities.db"

type Destination struct {
	Name string
	Lat  float32
	Lng  float32
}

func writeDistMatrix(matrix *maps.DistanceMatrixResponse) {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var city_id int
	err := db.QueryRow("select id from cities where name = ?; ", matrix.DestinationAddresses[0]).Scan(&city_id)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`
	insert into points(city_id, lat, lng, status, duration, duration_in_traffic, distance) values (?, ?, ?, ?, ?, ?, ?);
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i, element := range matrix.DistanceMatrixElementsRow {
		origin := matrix.OriginAddresses[i]
		pretty.Println(i, origin, element)
		//_, err :+ stmt.Exec(city_id, origin.Lat, origin.Lng, element.Status, element.Duration, element.DurationInTraffic, element.Distance)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	return

}

func createDB(target Destination) {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt := `
	create table cities (id integer primary key autoincrement, name text, lat float, lng float);
	create table points (id integer primary key autoincrement, city_id integer, lat float, lng float, status string,
	                     duration integer, duration_in_traffic integer, distance integer, FOREIGN KEY(city_id) REFERENCES cities(id));
	`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
		return
	}
	stmt = fmt.Sprintf(
		"insert into cities (name, lat, lng) values(\"%s\", %f, %f);",
		target.Name, target.Lat, target.Lng,
	)
	res, err := db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
	}
	fmt.Println(res)
}
