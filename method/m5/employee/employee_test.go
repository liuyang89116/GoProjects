package employee

import "testing"

// 不需要把所有的 stmt 接口都实现
type fakeStmtForMaleCount struct {
	Stmt
}

func (fakeStmtForMaleCount) Exec(stmt string, args ...string) (Result, error) {
	return Result{Count: 5}, nil
}

func TestEmployeeMaleCount(t *testing.T) {
	f := fakeStmtForMaleCount{}
	count, _ := MaleCount(f)
	if count != 5 {
		t.Errorf("Expected: %d, actual: %d", 5, count)
		return
	}
}
