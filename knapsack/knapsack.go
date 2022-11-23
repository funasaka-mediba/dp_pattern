package knapsack

import (
	"dp_pattern/stdin"
	"fmt"
	"log"
)

/*
ナップサック問題

n 個の品物があり、i 番目の品物のそれぞれ重さと価値が weight[i],value[i] となっている (i=0,1,...,n−1)。
これらの品物から重さの総和が W を超えないように選んだときの、価値の総和の最大値を求めよ。

【制約】
・1≤n≤100
・weight[i],value[i] は整数
・1≤weight[i],value[i]≤1000
・1≤W≤10000

【数値例】
1)
　n=6
　(weight,value)=(2,3),(1,2),(3,6),(2,1),(1,3),(5,85)
　W=9
　答え: 94 ((3,6), (1,3), (5,85) を選んだときが最大です)
*/

func Knapsack() {
	/*
		入力例
		knapsack
		6
		2 1 3 2 1 5
		3 2 6 1 3 85
		9

		答え
		94
	*/

	n, err := stdin.IntStdin()
	if err != nil {
		log.Fatalf("failed: %s", err)
	}
	weight, err := stdin.SplitIntStdin(" ")
	if err != nil {
		log.Fatalf("failed: %s", err)
	}
	value, err := stdin.SplitIntStdin(" ")
	if err != nil {
		log.Fatalf("failed: %s", err)
	}
	W, err := stdin.IntStdin()
	if err != nil {
		log.Fatalf("failed: %s", err)
	}

	/*
		max_sumと同じように、dp[i+1]を考えるときにに、dp[i]とvalue[i]を足すかどうかを判断するのだが、それにはweight[i]とWの情報が制約として必要になってくる。
		dp[i]
		value[i]
		weight[i]
		W
		をどのように使用して漸化式を作っていくかが問題。

		漸化式、数列を考えるときには、やはり順番に書き出していって関係性を見るのがわかりやすい。
		dp[0] = 0
		dp[1] = (weight[0]はw以下なので) value[0] = 3
		dp[2] = (weight[0]+weight[1]はw以下なので) dp[1]+value[1] = 3+2 = 5
		dp[3] = (weight[0]+weight[1]+weight[2]はw以下なので) dp[2]+value[2] = 5+6 = 11

		まとめると

		dp[i+1][w] := { max( dp[i][w - weight[i]] + value[i], dp[i][w])    (w >= weight[i] の時) }
					  { dp[i][w] 										   (w <  weight[i] の時) }

		dp[0][w] := 0 (w = 0, 1, 2, 3, 4,,,,W)
	*/

	var dp [110][10010]int
	for i := 0; i <= W; i++ {
		dp[0][i] = 0
	}

	for i := 0; i < n; i++ {
		for w := 0; w <= W; w++ {
			if w >= weight[i] {
				dp[i+1][w] = Max(dp[i][w-weight[i]]+value[i], dp[i][w])
			} else {
				dp[i+1][w] = dp[i][w]
			}
			fmt.Printf("dp[%v+1][%v]: %+v\n", i, w, dp[i+1][w])
		}
	}

	fmt.Printf("dp[n][W]: %+v\n", dp[n][W])
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
