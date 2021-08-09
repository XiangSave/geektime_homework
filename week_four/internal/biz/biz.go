package biz

import (
	"fmt"
	"log"
	"time"
	"week_four/internal/data"
	"week_four/internal/pkg/mysql"
)

func checkInSlice(s string, sli []*data.DepPeople) bool {
	for i := 0; i < len(sli); i++ {
		if sli[i].Name == s {
			return true
		}

	}
	return false
}

func UnOrderedString(m *mysql.DBModel, firstDeptId int) string {

	now := time.Now()
	tStr := now.Format("2006-01-02") + " 12:00:00"
	log.Printf("统计时间 %s", tStr)

	allPeople := data.GetAllDepPeople(m, firstDeptId)
	orderedPeoples := data.GetOrderedDepPeople(m, firstDeptId, tStr)

	var unOrderedPeoples = []string{}
	var orderedPeoplesList = []string{}
	for i := 0; i < len(allPeople); i++ {
		if !checkInSlice(allPeople[i].Name, orderedPeoples) {
			unOrderedPeoples = append(unOrderedPeoples, allPeople[i].Name)
		} else {
			orderedPeoplesList = append(orderedPeoplesList, allPeople[i].Name)
		}
	}

	log.Println(len(unOrderedPeoples))
	log.Printf("未订餐人数：%d", len(unOrderedPeoples))
	log.Println(orderedPeoplesList)
	log.Printf("订餐人数：%d", len(orderedPeoples))

	// log.Println(unOrderedPeoples)

	sendMsg := fmt.Sprintf("订餐人数：%d\n未订餐人数：%d\n未订餐人员：\n",
		len(orderedPeoples), len(unOrderedPeoples))

	sendMsg = sendMsg + fmt.Sprintf("%s", unOrderedPeoples)
	return sendMsg

}
