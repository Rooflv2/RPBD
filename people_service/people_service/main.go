package main

import (
	"fmt"

	"github.com/RyabovNick/databasecourse_p2/golang/tasks/people_service/service/store"
)

func main() {
	conn := "postgresql://tselikova:284463@95.217.232.188:7777/tselikova"
	s := store.NewStore(conn)
	fmt.Println(s.ListPeople())
	fmt.Println(s.GetPeopleByID(3))
}
