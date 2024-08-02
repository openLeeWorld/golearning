.DEFAULT_GOAL := build # 기본적으로 수행하는 타깃 정의

fmt: # 타깃 이름
	go fmt ./...  
.PHONY:fmt # 타깃과 동일한 이름으로 디렉토리를 생성한 경우, MAKE가 혼동X
# GO 소스코드의 형식을 맞춰줌

lint: fmt # 타깃 : 선행 타깃 조건 
	golint ./...go
.PHONY:lint # phony: 사전적으로 "가짜의"
# GO 소스코드의 코딩 스타일을 통일해줌

vet: lint
	go vet ./...
	shadow ./... 
.PHONY:vet # vet: 정적 분석도구, shadow: 섀도우 변수 감지 
# 도달 불가능 코드, 오류 구문, mutext 락, 불안정 포인터 등 분석

build: vet
	go build .
.PHONY:build
# 실제 바이너리로 컴파일