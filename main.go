package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/KatsutoshiOtogawa/batch/lib/config"
	"github.com/KatsutoshiOtogawa/batch/model/pkg"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

var REVISION string
var BUILD_TIME string

func main() {

	var pkgName string
	var funcName string

	db, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		log.Println(err.Error())
		return
	}
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "pkg",
				Value:       "",
				Usage:       "Write you want to use pkg name.",
				Destination: &pkgName,
			},
			&cli.StringFlag{
				Name:        "func",
				Value:       "",
				Usage:       "Write you want to use func name in pkg.",
				Destination: &funcName,
			},
		},
		Action: func(c *cli.Context) error {
			args := config.Args{
				FuncName: funcName,

				Db: db,
			}

			pkg.Invoke(pkgName, &args)

			return nil
		},
	}

	app.Run(os.Args)

}
