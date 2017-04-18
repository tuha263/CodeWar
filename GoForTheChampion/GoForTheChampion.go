package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Team save info of teams
type Team struct {
	name                          string
	goal, lose, score, totalMatch int
}

const (
	//quantityOfTeam is total team
	quantityOfTeam = 4
	//VN is name of viet nam team
	VN         = "VIETNAM"
	impossible = "IMPOSSIBLE"
)

var (
	team       [4]Team
	goalResult int
	loseResult int
)

func main() {
	ReadData()
	var (
		indexOfVN int
	)
	team[FindByName(VN, &indexOfVN)].score += 3
	SortTeams()

	indexOfVN = FindByName(VN, &indexOfVN)
	team2dn := &team[1]

	if indexOfVN < 2 {
		fmt.Println("1:0")
		os.Exit(0)
	}
	if team[indexOfVN].score < team2dn.score {
		fmt.Println(impossible)
		os.Exit(0)
	}

	fmt.Println(GetResult())
}

//GetResult return result
func GetResult() string {
	for i := 1; i < 10; i++ {
		if Try(i, 0) {
			for j := i; j < 10; j++ {
				if Try(j, j-(i-1)) {
					return strconv.Itoa(j) + ":" + strconv.Itoa(j-(i-1))
				}
			}
			return strconv.Itoa(i) + ":" + strconv.Itoa(0)
		}
	}
	return impossible
}

//ConvertResult convert result to string
func ConvertResult(goal, result int) {

}

//Try is return position of vietnam after add goal
func Try(goal, lose int) bool {
	var indexOfVN int
	indexOfVN = FindByName(VN, &indexOfVN)
	indexOfOP := FindCompetitor()

	team[indexOfVN].goal += goal
	team[indexOfVN].lose += lose

	team[indexOfOP].goal += lose
	team[indexOfOP].lose += goal

	SortTeams()
	/*fmt.Println("-------------")
	fmt.Println(strconv.Itoa(goal) + "-" + strconv.Itoa(lose))
	for i := 0; i < quantityOfTeam; i++ {
		fmt.Println(team[i])
	}*/

	indexOfVN = FindByName(VN, &indexOfVN)
	indexOfOP = FindCompetitor()
	team[indexOfVN].goal -= goal
	team[indexOfVN].lose -= lose

	team[indexOfOP].goal -= lose
	team[indexOfOP].lose -= goal
	if indexOfVN < 2 {
		return true
	}
	return false
}

//IsHighPosition return true if > and false if <
func (own Team) IsHighPosition(team Team) bool {
	//compare score
	if own.score > team.score {
		return true
	}
	if own.score < team.score {
		return false
	}

	//compare goal - lose
	if own.goal-own.lose > team.goal-team.lose {
		return true
	}
	if own.goal-own.lose < team.goal-team.lose {
		return false
	}

	//cmpare goal
	if own.goal > team.goal {
		return true
	}
	if own.goal < team.goal {
		return false
	}

	return own.name < team.name
}

//SortTeams sort team from high to low
func SortTeams() {
	for i := 0; i < quantityOfTeam-1; i++ {
		for j := i + 1; j < quantityOfTeam; j++ {
			if team[j].IsHighPosition(team[i]) {
				temp := team[i]
				team[i] = team[j]
				team[j] = temp
			}
		}
	}

}

//ReadData read data from console
func ReadData() {
	curTeam := -1
	for i := 0; i < 5; i++ {
		var t1, t2, match string
		fmt.Scan(&t1)
		fmt.Scan(&t2)
		fmt.Scan(&match)
		goals := strings.Split(match, ":")
		t1Goal, _ := strconv.Atoi(goals[0])
		t2Goal, _ := strconv.Atoi(goals[1])
		t1Index := FindByName(t1, &curTeam)
		t2Index := FindByName(t2, &curTeam)
		team[t1Index].goal += t1Goal
		team[t1Index].lose += t2Goal
		team[t2Index].goal += t2Goal
		team[t2Index].lose += t1Goal
		team[t1Index].totalMatch++
		team[t2Index].totalMatch++
		result := t1Goal - t2Goal
		if result < 0 {
			team[t1Index].score += 0
			team[t2Index].score += 3
		} else if result == 0 {
			team[t1Index].score++
			team[t2Index].score++
		} else if result > 0 {
			team[t1Index].score += 3
			team[t2Index].score += 0
		}
	}
}

//FindCompetitor find index of final competitor
func FindCompetitor() int {
	for i := 0; i < quantityOfTeam; i++ {
		if team[i].name != VN && team[i].totalMatch == 2 {
			return i
		}
	}
	return -1
}

//FindByName return team with same name
func FindByName(name string, cur *int) int {
	for i := 0; i < quantityOfTeam; i++ {
		if name == team[i].name {
			return i
		}
	}
	*cur++
	team[*cur].name = name
	return *cur
}
