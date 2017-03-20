package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/boolow5/BolWeydi/g"
	/* _ "github.com/go-sql-driver/mysql"
	   _ "github.com/lib/pq" */
	_ "github.com/mattn/go-sqlite3"
)

var o orm.Ormer

type MyModel interface {
	Valid() bool
	SetId(id int)
	String() string
}

func init() {
	fmt.Println("initializing models...")

	// REGISTER DRIVERS & DATABASES:
	/*// MYSQL
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	*/
	// SQLITE
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	orm.RegisterModel(new(User), new(Profile), new(Question), new(Answer), new(Discussion))

	if g.AUTO_MIGRATE {
		name := "default"                          // Database alias.
		force := true                              // Drop table and re-create.
		verbose := true                            // Print log
		err := orm.RunSyncdb(name, force, verbose) // Sync with database
		if err != nil {
			fmt.Println(err)
		}
	}
	o = orm.NewOrm()
	o.Raw("PRAGMA foreign_keys = ON")
}

// Q&A related:
// 1. Questions, 2. Answers, 3. Discussions
// OVERVIEW: Discussion is held by a writer, and readers ask a question related to the topics,
// When the writer answers all the Questions in the Discussion, it's released to the readers.
// If user ask question, s/he gets notified whenever followed, answered or edited.
// users have notification collection, questions collection, answers collection.
// to keep the database clean, there is collection and reference collections.
// for example there is notification collection containing all notifications
// reference to notification in every user it is targetting.

// Question may belong to more than one Topic, Topic can have child Topics
// Anwer belongs only to one question
