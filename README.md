# go-test-mysql-goose

This packege is just a wrapper for [go-test-mysqld](https://github.com/lestrrat/go-test-mysqld) and [goose](https://bitbucket.org/liamstask/goose).

At first launch MySQL instance by go-test-mysqld, and then automatically deploy some SQLs by goose.

## Examples

Following code launch MySQL instance with default settings,
and deploy SQL migrations on `./db` directory.

```go
import (
    "github.com/typester/go-test-mysqld-goose/mysqltest"
    "path/filepath"
)

mysql, err := mysqltest.New(filepath.Join("db"))
if err != nil {
    panic(err)
}
defer mysql.Server.Stop()
```

This package also provides connection helper function `Dns()` for [mymysql](https://github.com/ziutek/mymysql) driver.
To use this function, you can connect test mysql server very easily.

```go
import (
    "database/sql"
    _"github.com/ziutek/mymysql/godrv"
)

db, err := db.Open("mymysql", mysql.Dns())
if err != nil {
    panic(err)
}
defer db.Close()

// use db
```

## Notice

This package is made for my personal usage so that many features may lacks for others.
Patches are very appreciated.
