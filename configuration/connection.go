package configuration

import (
        "database/sql"
        "encoding/json"
        "fmt"
        _ "github.com/lib/pq"
        "log"
        "os"
)

type configuration struct {
        DBServer string `json:"db_server"`
        DBPort int `json:"db_port"`
        DBName string `json:"db_name"`
        DBUser string `json:"db_user"`
        DBPassword string `json:"db_password"`
}

func getConfiguration() (configuration, error)  {
        config := configuration{}

        file, err := os.Open("./config.json")
        if err != nil {
                return config, err
        }
        defer file.Close()

        decoder := json.NewDecoder(file)
        err = decoder.Decode(&config)
        if err != nil {
                return config, err
        }

        return config, err
}

func GetConnectionPsql() *sql.DB  {
        config, err := getConfiguration()
        if err != nil {
                log.Fatalln(err)
        }

        dsn := fmt.Sprintf(
                "postgres://%s:%s@%s:%v/%s?sslmode=disable",
                config.DBUser,
                config.DBPassword,
                config.DBServer,
                config.DBPort,
                config.DBName)
        fmt.Println(dsn)

        db, err := sql.Open("postgres", dsn)
        if err != nil {
                log.Fatalln(err)
        }
        return db
}
