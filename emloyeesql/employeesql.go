package employeesql

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	emp_id         int
	emp_name       string
	emp_position   string
	emp_experience float64
	emp_salary     float64
}

func insert(e employee) (employee, error) {
	db, err := sql.Open("mysql", "root:tzopsmart@tcp(127.0.0.1:3306)/test")

	defer db.Close()

	if err != nil {
		return employee{}, err
	}
	if e.emp_id <= 0 {
		err = errors.New("id should be 1 and above")
		return employee{}, err
	}

	if e.emp_name == "" {
		err = errors.New("name can't be empty")
		return employee{}, err
	}

	res, err := db.Exec("INSERT INTO Employee (emp_id, emp_name,emp_position,emp_exp,emp_salary) VALUES (?,? ,?, ?, ?)", e.emp_id, e.emp_name, e.emp_position, e.emp_experience, e.emp_salary)

	if err != nil {
		return employee{}, err
	}

	lastId, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if lastId == 1 {

		return e, nil
	}
	return e, nil
}

func fetch(id int) (employee, error) {

	db, err := sql.Open("mysql", "root:tzopsmart@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		return employee{}, err
	}

	if id <= 0 {
		err = errors.New("id should be 1 and above")
		return employee{}, err
	}
	var e employee
	res, err := db.Query("SELECT * FROM Employee WHERE emp_id=?", id)

	defer res.Close()

	if err != nil {
		return employee{}, err
	}

	if res.Next() {
		err = res.Scan(&e.emp_id, &e.emp_name, &e.emp_position, &e.emp_experience, &e.emp_salary)

		if err != nil {
			return employee{}, err
		}
	}
	return e, nil
}
