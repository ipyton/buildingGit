package diff

import "strings"

type Myers struct {
	a []string
	b []string
}

type Edit struct {
	operation string
	text string
}
func newEdit(text, operation string) * Edit{
	return & Edit{text: text, operation: operation}
}


func NewMyers(a []string, b []string) * Myers{
	return & Myers{a: a,b:b}
}


func (this * Myers) diff(a string, b string) {

}

func (this * Myers) shortestEdit() []string {
	dp := make([][]int, len(this.a) + 10, len(this.a) + 10)
	for i:=0; i < len(dp); i ++ {
		dp[i] = make([]int, len(this.b) + 10, len(this.b) + 10)
	}
	ops := make([]string, len(this.b) + 10, len(this.b) + 10)

	for i := 1; i < len(this.a); i ++ {
		for j := 1; j < len(this.b); j ++ {
			if this.a[i - 1] == this.b[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
				ops[j] = "eq"
			} else {
				if dp[i][j - 1] <= dp[i - 1][j] {
					ops[j] = "del"
				} else {
					ops[j] = "add"
				}
				dp[i][j] = min(dp[i][j - 1] + 1 , dp[i - 1][j] + 1)
			}
		}
	}
	return ops
}


func (this * Myers) getResults(ops, text []string) []* Edit{
	symbol := map[string]string{"del": "-", "add" : "+", "eq" : ""}
	edits := make([]*Edit, len(text), len(text))
	for idx, value := range ops {
		edits[idx] = newEdit(text[idx], symbol[value])
	}
	return edits
}

func lines(a string) []string {
	return strings.Split(a, "\n\r")
}