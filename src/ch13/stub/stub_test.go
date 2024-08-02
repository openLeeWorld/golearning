package stub

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// 해당 메서드를 테스트하는데 필요한 메서드만 구현하는 스텁 구조체
type GetPetNamesStub struct {
	Entities
}

// 스텁 구조체에 대한 테스트 메서드 구현
func (ps GetPetNamesStub) GetPets(userId string) ([]Pet, error) {
	switch userId {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userId)
	}
}

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Snowball II"}},
		//{"case3", "3", nil},
	}
	l := Logic{GetPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
}

// 아래는 함수 항목에 메서드 호출을 연결하는 스텁 구조체를 정의해서 사용
// Entities에 정의된 각 메서드를 위해 스텁 구조체에 동일한 시그리너의 함수 항목을 정의
type EntitiesStub struct {
	getUser     func(id string) (User, error)
	getPets     func(userId string) ([]Pet, error)
	getChildren func(userId string) ([]Person, error)
	getFriends  func(userId string) ([]Person, error)
	saveUser    func(user User) error
}

// 메서드르 정의하여 EntitiesStub이 Entites 인터페이스를 충족하도록 한다.
func (es EntitiesStub) GetUser(id string) (User, error) {
	return es.getUser(id)
}

func (es EntitiesStub) GetPets(userId string) ([]Pet, error) {
	return es.getPets(userId)
}

func (es EntitiesStub) GetChildren(userId string) ([]Person, error) {
	return es.getChildren(userId)
}

func (es EntitiesStub) GetFriends(userId string) ([]Person, error) {
	return es.getFriends(userId)
}

func (es EntitiesStub) SaveUser(user User) error {
	return es.saveUser(user)
}

// 테이블 테스트를 위한 자료 구조의 항목을 통해 다른 테스트에서 다른 메서드의 다른 구현 제공
func TestLogicGetPetNames2(t *testing.T) {
	data := []struct { // 함수 타입의 항목 및 데이터
		name     string
		getPets  func(userID string) ([]Pet, error)
		userID   string
		petNames []string
		errMsg   string
	}{
		{"case1", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, ""},
		{"case2", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
		}, "2", []string{"Stampy", "Snowball II"}, ""},
		{"case3", func(userID string) ([]Pet, error) {
			return nil, errors.New("invalid id: 3")
		}, "3", nil, "invalid id: 3"},
	}
	l := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = EntitiesStub{getPets: d.getPets} // 새로운 EntitesStub이 초기화
			petNames, err := l.GetPetNames(d.userID)
			if diff := cmp.Diff(petNames, d.petNames); diff != "" {
				t.Error(diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
