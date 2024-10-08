/*
	You are given an m x n integer grid accounts where accounts[i][j] is the amount

of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest
customer has.

A customer's wealth is the amount of money they have in all their bank accounts.
The richest customer is the customer that has the maximum wealth.
*/
package richestcustomerwealth

func MaximumWealth(accounts [][]int) int {
	var res int
	for _, linha := range accounts {
		sum := 0
		for _, coluna := range linha {
			sum += coluna
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
