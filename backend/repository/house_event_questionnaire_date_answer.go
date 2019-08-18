package repository

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

// todo 多対多なので、おそらく手直しが必要になる
func HEQDAnswerAll(db *gorm.DB, questionnaireDateId uint) (*[]model.HouseEventQuestionnaireDateAnswer) {
	questionnaireDate := model.HouseEventQuestionnaireDate{}

	db.Find(&questionnaireDate, questionnaireDateId)
	db.Model(&questionnaireDate).Related(&questionnaireDate.Answers, "Answers")

	return &questionnaireDate.Answers
}

func HEQDAnswerCreate(db *gorm.DB, questionnaireDateId uint, createDateAnswer *model.HouseEventQuestionnaireDateAnswer) *model.HouseEventQuestionnaireDateAnswer {
	db.Debug().Exec(`
INSERT INTO house_event_questionnaire_date_answers (user_id, questionnaire_date_id, ok) VALUES (?, ?, ?)
`, createDateAnswer.UserID, createDateAnswer.QuestionnaireDateID, createDateAnswer.Ok)
	return createDateAnswer
}

func HEQDAnswerUpdate(db *gorm.DB, questionnaireDateAnswerId uint, updateDateAnswer *model.HouseEventQuestionnaireDateAnswer) *model.HouseEventQuestionnaireDateAnswer {
	db.Debug().Exec(`
UPDATE house_event_questionnaire_date_answers SET ok = ? WHERE id = ? 
`, updateDateAnswer.Ok, questionnaireDateAnswerId)
	return updateDateAnswer
}

func HEQDAnswerDelete(db *gorm.DB, dateAnswerId uint) {
	dateAnswer := model.HouseEventQuestionnaireDateAnswer{}
	db.Delete(&dateAnswer, dateAnswerId)
}

func HEQDAnswerIsMine(db *gorm.DB, dateAnswerId uint, userId uint) bool {
	dateAnswer := model.HouseEventQuestionnaireDateAnswer{}
	dateAnswer.ID = dateAnswerId
	var count int
	db.Debug().Where("user_id = ?",userId).Find(&dateAnswer).Count(&count)

	fmt.Println("count")
	fmt.Println(count)
	if count > 0 {
		return true
	}
	return false
}
