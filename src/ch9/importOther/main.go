package main

import (
	"fmt" // 내장 모듈

	"github.com/learning-go-book/package_example/formatter" // 깃허브 소스 모듈
	"github.com/learning-go-book/package_example/math"      // go.mod 있음
) // 깃허브/사용자/리포지토리이름/폴더
// 임포트 경로는 모듈에 있는 패키지 경로를  모듈 경로에 추가하여 만듬
// 절대경로를 왠만하면 사용 (pkg/mod에 저장됨)

func main() {
	num := math.Double(2) //패키지명.함수 로 호출
	output := print.Format(num)
	fmt.Println(output)
}
