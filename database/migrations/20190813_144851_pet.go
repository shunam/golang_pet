package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Pet_20190813_144851 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Pet_20190813_144851{}
	m.Created = "20190813_144851"

	migration.Register("Pet_20190813_144851", m)
}

// Run the migrations
func (m *Pet_20190813_144851) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE pet(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`age` int(11) DEFAULT NULL,`photo` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Pet_20190813_144851) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `pet`")
}
