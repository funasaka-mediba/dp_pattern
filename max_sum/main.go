package main

import "fmt"

/*
最大和問題

n個の整数が与えられる。
これらの整数から何個かの整数を選んで総和をとったときの、総和の最大値を求めよ。また、何も選ばない場合の総和は 0 であるものとする。
*/

func main() {
	// 標準入力はいつでもできるから一旦スキップ
	n := 6
	input := []int{1, 4, -1, -6, 9, 2}

	// これは正の整数を全て選んで足していけば総和の最大値になるが、DP法に慣れるためにやってみよう。
	// DP法とは、全探索で行なっていくと計算量がとても大きくなってしまう場合に、一度行った計算結果をメモしておき、それを利用すれば計算量が減らせるんじゃないのという考え方。
	// 数学の漸化式を作っていくことにも似ている。
	// 漸化式を考えるときに何を考えるかというと、数列のi番目とi+1番目の関係を式にして、初期値を与えてあげることで式が作れる。

	/*
		inputをa[i]とすると、
		a[0] = 1
		a[1] = 4
		a[2] = -1
		a[3] = -6
		a[4] = 9
		a[5] = 2

		ここで　dp []int というメモ用のスライスを考えていく。ここでいうメモの内容はそれまでの総和の最大値のこと。
		求めたい答えのメモを作っていくと考える。
		dp[0] := なにも選ばない状態(つまり0)
		dp[1] := a[0]までの中で総和の最大値を作るとき・・・a[0] = 1は正の整数だから 1が最大
		dp[2] := a[0]からa[1]までの間で総和の最大値を作るとき・・・a[1]は正の整数だから、a[0]+a[1] = 1+4 = 5
		dp[3] := a[0]からa[2]までの間で総和の最大値を作るとき・・・a[2]は負の整数だから足さない。a[0]+a[1] = 1+4 = 5
		~~
		dp[i+1] := a[i]が正だったら dp[i]+a[i]、a[i]が負だったらdp[i]

		この関係式を漸化式で表すなら、d[i+1] = max(dp[i], dp[i]+a[i])
	*/

	dp := make([]int, 7)
	dp[0] = 0
	for i := 0; i < n; i++ {
		dp[i+1] = Max(dp[i], dp[i]+input[i])
		fmt.Printf("dp[%v]: %+v\n", i+1, dp[i+1])
	}
	fmt.Printf("max sum is: %+v\n", dp[n])
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
