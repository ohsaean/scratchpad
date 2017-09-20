package main

import (
	"fmt"
	"sort"
)

/*
[매일매일 알고리즘 트레이닝]  lv.3

밤늦게까지 놀다온 학생들이 선생님들을 피해 학교 정문에서 기숙사까지 들키지 않고 가려고한다.

학생들은 투명 망토가 있는데 망토는 한번에 최대 두사람씩 이용할 수 있다.

각 학생들은 학교 정문에서 기숙사까지 가는데 걸리는 시간이 주어지는데 만약 속도가 다른 두 사람이 망토를 같이 쓰고 정문에서

기숙사까지 간다면 더 빨리 갈 수 있는 사람이 더 느린사람의 페이스에 맞춰서 가야한다.

이런식으로 최대 두명씩 망토를 쓰고 이동가능하다고할때 최대한 빨리 기숙사에 도착하려고한다. 예를 한번 들어보자.

학생 4명 (A,B,C,D)가 있다고하고 각각 정문에서 기숙사까지 1분,2분,7분,10분이 걸린다고하면, 4명모두 기숙사까지 최대한 빨리 가는 방법은 다음과 같다:

A,B가 망토를 쓰고 정문에서 기숙사로 간다 (2분걸림)

A가 망토를 쓰고 기숙사에서 정문으로 돌아온다 (1분걸림)

C,D가 망토를 쓰고 정문에서 기숙사로 간다 (10분걸림)

B가 망토를 쓰고 기숙사에서 정문으로 돌아온다 (2분걸림)

A,B가 망토를 쓰고 정문에서 기숙사로 간다 (2분걸림)

따라서 총 17분이 걸리고 이것이 4사람이 정문에서 기숙사까지 들키지 않고 갈 수 있는 최소시간이다.

입력 : 학생들의 수 N (1 <= N <= 15)가 주어지고 한줄에 N개의 숫자들이 주어진다. 이 숫자는 각 학생이 정문에서 기숙사로 가는데 걸리는 시간을 의미한다.

출력 : N명의 학생들이 정문에서 기숙사까지 가는데 걸리는 최소시간을 출력한다.

예제 입력 2 15 5
예제 출력 15
예제 입력 4 1 2 7 10
예제 출력 17
예제 입력 5 12 1 3 8 6
예제 출력 29
*/

func main() {

	g := []int{12, 1, 3, 8, 6}
	//g := []int{1, 2, 7, 10}
	d := make([]int, 0)

	time := 0
	t := 0
	for {
		// 1. 무리중 제일 빠른 학생 2명 (or 1명) 선택 후 먼저 보냄
		g, d, t = SendFast2p(g, d)
		time += t
		if len(g) == 0 {
			break
		}

		// 2. 기숙사에 있는 제일빠른 학생 1명 망토들고 다시 돌아옴
		d, g, t = BackFast1p(d, g)
		time += t

		// 3. 제일 느린 학생들 2명 보냄
		g, d, t = SendSlow2p(g, d)
		time += t
		if len(g) == 0 {
			break
		}

		// 4. 기숙사에 있는 제일빠른 학생 1명 망토들고 다시 돌아옴
		d, g, t = BackFast1p(d, g)
		time += t
	}
	fmt.Println(time)
}

func SendFast2p(src []int, dst []int) ([]int, []int, int) {
	elapsed := 0
	sort.Ints(src)
	sort.Ints(dst)
	if len(src) < 2 {
		dst = append(dst, src[0])
		elapsed = src[0]
		src = nil
	} else if len(src) == 2 {
		dst = append(dst, src[0], src[1]) // 0 ~ 1
		elapsed = src[1]
		src = nil

	} else {
		dst = append(dst, src[0], src[1]) // 0 ~ 1
		elapsed = src[1]
		src = src[2:]
	}
	fmt.Println("SendFast2p g", src)
	fmt.Println("SendFast2p d ", dst)
	return src, dst, elapsed
}

func SendSlow2p(src []int, dst []int) ([]int, []int, int) {
	elapsed := 0
	sort.Sort(sort.Reverse(sort.IntSlice(src)))
	sort.Ints(dst)
	if len(src) < 2 {
		dst = append(dst, src[0])
		elapsed = src[0]
		src = nil
	} else if len(src) == 2 {
		dst = append(dst, src[0], src[1]) // 0 ~ 1
		elapsed = src[0]
		src = nil
	} else {
		dst = append(dst, src[0], src[1]) // 0 ~ 1
		elapsed = src[0]
		src = src[2:]
	}
	fmt.Println("SendSlow2p g", src)
	fmt.Println("SendSlow2p d", dst)
	return src, dst, elapsed
}

func BackFast1p(src []int, dst []int) ([]int, []int, int) {
	elapsed := 0
	sort.Ints(src)
	sort.Ints(dst)
	dst = append(dst, src[0])
	elapsed = src[0]
	src = src[1:]
	fmt.Println("BackFast1p g", dst)
	fmt.Println("BackFast1p d", src)
	return src, dst, elapsed
}
