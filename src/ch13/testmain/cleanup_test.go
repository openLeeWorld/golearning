package testmain

// 이 파일은 동작하지 않는 예시 파일임
import (
	"os"
	"testing"
)

// CreateFile은 여러 테스트에서 호출되는 헬퍼 함수
func CreateFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	// f에 데이터를 쓴다.
	t.Cleanup(func() { // 단일 테스트를 위해 생성된 임시 자원을 정리
		os.Remove(f.Name())
	}) // 샘플데이터를 설정하기 위한 헬퍼 함수에 의존적일 때 유용
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	//fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// 테스트 수행, 정리는 위에서 함
}
