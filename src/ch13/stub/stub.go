package stub

type User struct{}
type Pet struct {
	Name string
}
type Person struct{}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userId string) ([]Pet, error)
	GetChildren(userId string) ([]Person, error)
	GetFriends(userId string) ([]Person, error)
	SaveUser(user User) error
}

// 대형 인터페이스에 의존하는 코드를 테스트하기 위해 구조체에 인터페이스를 넣음
type Logic struct {
	Entities Entities // 구조체에 인터페이스 메서드 모두를 자동적으로 정의
}

// 테스트 대상 메서드
func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(pets)) // 슬라이스 길이 0 초기화, 용량 pets 길이
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
}
