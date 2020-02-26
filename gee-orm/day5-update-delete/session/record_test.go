package session

import "testing"

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession()
	err1 := s.DropTable(&User{})
	err2 := s.CreateTable(&User{})
	_, err3 := s.Create(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestSession_Create(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Create(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	users := []User{}
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}

func TestSession_First(t *testing.T) {
	s := testRecordInit(t)
	u := &User{}
	err := s.First(u)
	if err != nil || u.Name != "Tom" || u.Age != 18 {
		t.Fatal("failed to query first")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestSession_Where(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	_, err1 := s.Create(user3)
	err2 := s.Where("Age = ?", 25).Find(&users)

	if err1 != nil || err2 != nil || len(users) != 2 {
		t.Fatal("failed to query with where condition")
	}
}

func TestSession_OrderBy(t *testing.T) {
	s := testRecordInit(t)
	u := &User{}
	err := s.OrderBy("Age DESC").First(u)

	if err != nil || u.Age != 25 {
		t.Fatal("failed to query with order by condition")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "Tom").Set("Age", 30).Update(&User{})
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)

	if affected != 1 || u.Age != 30 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "Tom").Delete("User")
	count, _ := s.Count("User")

	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}
