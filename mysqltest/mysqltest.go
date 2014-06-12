package mysqltest

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	mt "github.com/lestrrat/go-test-mysqld"
    "fmt"
)

type TestMysqld struct {
    Server *mt.TestMysqld
}

// New function launch new MySQL instance with default settings,
// and deploy SQLs in `dir` by `goose` automatically.
func New(dir string) (*TestMysqld, error) {
    mysql, err := mt.NewMysqld(nil)
	if err != nil {
		return nil, err
	}

    t := &TestMysqld{mysql}

    if err := t.deployDB(dir); err != nil {
        mysql.Stop()
        return nil, err
    }

	return &TestMysqld{mysql}, nil
}

func (t *TestMysqld) deployDB(dir string) error {
    conf, err := goose.NewDBConf(dir, "development")
	if err != nil {
        return err
	}
    conf.Driver.OpenStr = t.Dns()

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
        return err
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		return err
	}

    return nil
}

// Dns returns connection info for mymysql driver
func (t *TestMysqld) Dns() string {
    return fmt.Sprintf("unix:%s*test/root/", t.Server.Socket())
}
