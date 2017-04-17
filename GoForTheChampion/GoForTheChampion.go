package main

import (
	"bufio"
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
	VN = "VIETNAM"
)

var (
	team            [4]Team
	indexOfVN       int
	finalCompetitor int
)

func main() {
	ReadData()
	indexOfVN = FindByName(VN, &indexOfVN)
	finalCompetitor = FindCompetitor()

	fmt.Println(indexOfVN)
	for i := 0; i < quantityOfTeam; i++ {
		fmt.Println(team[i])
	}
}

//ReadData read data from console
func ReadData() {
	in := bufio.NewReader(os.Stdin)

	ss, _ := in.ReadString('.')

	ls := strings.Split(ss, string('\n'))
	curTeam := -1
	for i := 0; i < len(ls); i++ {
		matchs := strings.Split(ls[i], " ")
		t1 := matchs[0]
		t2 := matchs[1]
		goals := strings.Split(matchs[2], ":")
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
