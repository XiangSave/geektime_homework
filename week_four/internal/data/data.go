package data

import (
	"fmt"
	"log"
	"week_four/internal/pkg/mysql"
)

type DepPeople struct {
	Name string `json:"name"`
}

func queryDepPeople(m *mysql.DBModel, q string, d int) []*DepPeople {
	depPeoples := []*DepPeople{}
	query := fmt.Sprintf(q, d)

	rows, err := m.DBEngine.Query(query)
	if err != nil {
		log.Panicln(err)
	}
	defer rows.Close()

	for rows.Next() {
		depPeople := new(DepPeople)
		err := rows.Scan(&depPeople.Name)
		if err != nil {
			log.Panicln(err)
		}

		depPeoples = append(depPeoples, depPeople)
	}

	return depPeoples

}

func GetAllDepPeople(m *mysql.DBModel, d int) []*DepPeople {
	query := "SELECT user_name FROM tbl_user WHERE first_dept_id = %d and is_available = 0 and user_name NOT LIKE '%%离职%%';"
	return queryDepPeople(m, query, d)
}

func GetOrderedDepPeople(m *mysql.DBModel, d int, time string) []*DepPeople {
	query := fmt.Sprintf("SELECT u.user_name FROM tbl_order o,tbl_user u WHERE o.user_id=u.id and first_dept_id = %%d  and o.order_eat_date = \"%s\" and is_order_status = 1;", time)
	return queryDepPeople(m, query, d)
}
