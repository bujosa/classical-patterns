package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	counter2 := GetInstance()
	if counter2 != counter1 {
		t.Error("expected same instance after calling GetInstance() twice, got two different instances")
	}
}