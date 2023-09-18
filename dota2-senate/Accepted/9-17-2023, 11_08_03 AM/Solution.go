// https://leetcode.com/problems/dota2-senate

type Queue []byte

func (q *Queue) push(s byte) {
    *q = append(*q, s)
}

func predictPartyVictory(senate string) string {
    var d, r int
    arg := Queue(senate)
    for true{
        q := make(Queue, 0)
        for i := 0; i<len(arg); i++ {
            if arg[i] == 'R' {
                if d == 0 {
                    r++
                    q.push('R')
                } else {
                    d--
                }
            } else {
                if r == 0 {
                    d++
                    q.push('D')
                } else {
                    r--
                }
            }
        }
        if len(q) == len(arg) {
            if d > 0 {
                return "Dire"
            } else {
                return "Radiant"
            }
        }
        arg = q
    }
    return "Radiant"
}