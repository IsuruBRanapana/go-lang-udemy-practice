package model

func CreateTodo() error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", "isuru", "practice")
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
