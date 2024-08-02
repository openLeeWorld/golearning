package main

import (
	"context"
)

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	ch := make(chan wrapper, 1) // 버퍼 크기가 1인 wrapper인 채널
	// 채널을 버퍼링하면 버퍼링된 값이 취소로 인해 읽히지 않더라도 고루틴이 종료됨
	go func() { // 오래 수행하는 함수로부터 결과를 받아 버퍼가 있는 채널에 넣는다.
		// do the long running thing
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()
	select {
	case data := <-ch: // 취소 함수의 실행이나 타임아웃으로 인해 컨텍스트 취소 시
		return data.result, data.err
	case <-ctx.Done(): // 컨텍스트가 취소되면 실행
		return "", ctx.Err() // 왜 취소되었는지 알려주는 데이터를 위한 제로 값과 오류
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {
	//TODO
}
