package main

import (
    _ "bitbucket.org/phiggins/go-db2-cli"
    "fmt"
    "database/sql"
    "os"
    "flag"
)

type JOB_REF_T struct {
    JOB_ID      int
    REF_JOB_ID  int
    REF_LEVEL   int
    REF_GROUP   int
    CON_DEF     int
}

var (
    database = flag.String("dbname", "dbctl", "DB2 database name")
    uid = flag.String("uid", "", "DB2 user name")
    pwd = flag.String("pwd", "", "DB2 user password")
)

func main() {
    flag.Parse()
    if *database == "" || *uid == "" || *pwd == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

    connStr := fmt.Sprintf("DATABASE=%s; HOSTNAME=localhost; PORT=50000; PROTOCOL=TCPIP; UID=%s; PWD=%s;", *database, *uid, *pwd)
    db, err := sql.Open("db2-cli", connStr)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer db.Close()

    st, err := db.Prepare("SELECT * FROM CTL.JOB_REF")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer st.Close()

    rows, err := st.Query()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rows.Close()

    for rows.Next() {
        var ref JOB_REF_T
        err = rows.Scan(&ref.JOB_ID, &ref.REF_JOB_ID, &ref.REF_LEVEL, &ref.REF_GROUP, &ref.CON_DEF)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(ref)
    }
    if rows.Err() != nil {
        fmt.Println(rows.Err())
        os.Exit(1)
    }
}
