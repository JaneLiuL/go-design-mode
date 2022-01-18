package withoutorm

import "testing"

func TestInsertDB(t *testing.T) {
	db, err := ConnMysql()
	if err != nil {
		t.Fatal("connect mysql fail")
	}
	result, err := InsertDB(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestQueryDB(t *testing.T) {
	db, err := ConnMysql()
	if err != nil {
		t.Fatal("connect mysql fail")
	}

	result, err := QueryDB(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
