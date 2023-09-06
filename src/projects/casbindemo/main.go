package main

import (
	"fmt"
	"gopractice/projects/casbindemo/server"
	"log"
)

func main() {
	s := server.InitServer()

	ok, err := s.Enforcer.AddRoleForUser("alice", "admin")
	fmt.Printf("%s, %s\n", ok, err)
	ok, err = s.Enforcer.AddPermissionForUser("admin", "data1", "read")
	fmt.Printf("%s, %s\n", ok, err)
	ok, err = s.Enforcer.AddPermissionForUser("alice", "data2", "read")
	fmt.Printf("%s, %s\n", ok, err)

	ok, err = s.Enforcer.Enforce("admin", "data2", "read")
	if err != nil {
		log.Fatal("check fail.")
	}

	ok, err = s.Enforcer.Enforce("admin", "data1", "read")
	if err != nil {
		log.Fatal("check fail.")
	}

	ok, err = s.Enforcer.Enforce("alice", "data1", "read")
	ok, err = s.Enforcer.Enforce("alice", "data2", "read")

	return
}
