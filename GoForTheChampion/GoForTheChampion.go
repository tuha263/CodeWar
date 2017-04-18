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
	team   [4]Team
	result int
)

func main() {
	ReadData()
	var (
		indexOfVN int
	)
	team[FindByName(VN, &indexOfVN)].score += 3
	ShortTeams()

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

	for FindByName(VN, &indexOfVN) > 1 {
		AddGoal()
	}

	fmt.Println(strconv.Itoa(result) + ":0")
}

// AddGoal add gold to teams
func AddGoal() {
	var indexOfVN int

	k := GetVNGoal()
	result += k
	finalCompetitor := FindCompetitor()
	team[FindByName(VN, &indexOfVN)].goal += k
	team[finalCompetitor].lose += k
	ShortTeams()
}

//GetVNGoal return VN goal to get 2nd
func GetVNGoal() int {
	var (
		indexOfVN, finalCompetitor int
	)

	ShortTeams()
	indexOfVN = FindByName(VN, &indexOfVN)
	finalCompetitor = FindCompetitor()

	vnTeam := &team[indexOfVN]
	opTeam := &team[finalCompetitor]
	team2dn := &team[1]

	space := (team2dn.goal - team2dn.lose) - (vnTeam.goal - vnTeam.lose)
	if space == 0 {
		return 1
	}

	//fmt.Println(space)

	if finalCompetitor == 1 {
		var k int
		k = (space + 1) / 2

		if VN > opTeam.name && space%2 == 0 {
			k++
		}
		return k
	} else {
		if team[1].goal-vnTeam.goal > space || (team[1].goal-vnTeam.goal == space && team[1].name < vnTeam.name) {
			return space + 1
		} else {
			return space
		}
	}
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

//ShortTeams sort team from high to low
func ShortTeams() {
	for i := 0; i < quantityOfTeam-1; i++ {
		for j := i + 1; j < quantityOfTeam; j++ {
			if team[j].IsHighPosition(team[i]) {
				temp := team[i]
				team[i] = team[j]
				team[j] = temp
			}
		}
	}

	/*for i := 0; i < quantityOfTeam; i++ {
		fmt.Println(team[i])
	}*/
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
