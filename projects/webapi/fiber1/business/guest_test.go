package business

import (
	"errors"
	"fiber1/config"
	"fiber1/model"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hexops/autogold"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

var testDb *sqlx.DB
func TestMain(m *testing.M)()  {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Couldn't connect to docker: %s", err)
	}

	resource, err := pool.Run("mariadb", "latest", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Couldn't start resource: %s", err)
	}

	var connStr string
	if err := pool.Retry(func() error {
		connStr = fmt.Sprintf("root:secret@(127.0.0.1:%s)/mysql", resource.GetPort("3306/tcp"))
		testDb = config.ConnectMySql(connStr)
		if  testDb != nil{
			return nil
		}
		return errors.New(`testDb is nil`)
		}); err != nil {
		log.Fatalf("couldn't connect to database: %s", err)
	}

	log.Println(connStr)
	code := m.Run()
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("couldn't purge resource: %s", err)
	}

	os.Exit(code)
}

func TestGuest(t *testing.T) {
	user := model.NewUser(testDb)
	err := user.Migrate()
	if err != nil {
		t.Errorf("error migrating: %v", err)
	}

	guest := Guest{
		Db *sqlx.DB,
	}

	t.Run(`registerWithEmptyEmailMustFail`, func(t *testing.T) {
		in := Guest_RegisterIn{}
		out := Guest.RegisterOut{&in}

		want := autogold.Want(`registerWithEmptyEmailMustFail1`, Guest_RegisterOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg: "invalid email",
		}})
		want.Equal(t, out)
	})
}