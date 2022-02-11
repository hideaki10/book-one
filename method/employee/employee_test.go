package employee

import "testing"

type fakeStmtForMaleCount struct {
	Stmt
}

func (fakeStmtForMaleCount) Exec(stmt string, args ...string) (Result, error) {

	return Result{Count: 1}, nil
}

func TestEmployeeMaleCount(t *testing.T) {
	f := fakeStmtForMaleCount{}
	c, _ := MaleCount(f)
	if c != 1 {
		t.Errorf("want: %d, actual: %d", 1, c)
		return
	}
}
