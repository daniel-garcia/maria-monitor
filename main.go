package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

func main() {

	go metrics.Log(metrics.DefaultRegistry, time.Second * 10, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	for {
		func() {
			db, err := sql.Open("mysql", "")
			defer db.Close()
			if err != nil {
				log.Fatalf("Could not open connection to db: %s", err)
			}

			interval := time.Tick(time.Second * 10)

			for {

				select {
				case <-interval:
					rows, err := db.Query("SHOW /*!50002 GLOBAL */ STATUS")
					if err != nil {
						log.Printf("%s", err)
						break
					}

					for rows.Next() {
						var name string
						var val int64
						err := rows.Scan(&name, &val)
						if err != nil {
							continue
						} else {
							metrics.GetOrRegisterGauge(fmt.Sprintf("mariadb.%s", name), metrics.DefaultRegistry).Update(val)
						}
					}
				}
			}
		}()
	}
}
