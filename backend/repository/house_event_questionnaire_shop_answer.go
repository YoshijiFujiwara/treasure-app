package repository

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/model"
	"github.com/jinzhu/gorm"
)

// todo 多対多なので、おそらく手直しが必要になる
func HEQSAnswerAll(db *gorm.DB, questionnaireShopId uint) (*[]model.HouseEventQuestionnaireShopAnswer) {
	questionnaireShop := model.HouseEventQuestionnaireShop{}

	db.Find(&questionnaireShop, questionnaireShopId)
	db.Model(&questionnaireShop).Related(&questionnaireShop.Answers, "Answers")

	return &questionnaireShop.Answers
}

func HEQSAnswerCreate(db *gorm.DB, questionnaireShopId uint, createShopAnswer *model.HouseEventQuestionnaireShopAnswer) *model.HouseEventQuestionnaireShopAnswer {
	db.Debug().Exec(`
INSERT INTO house_event_questionnaire_shop_answers (user_id, questionnaire_shop_id, ok) VALUES (?, ?, ?)
`, createShopAnswer.UserID, createShopAnswer.QuestionnaireShopID, createShopAnswer.Ok)

	return createShopAnswer
}

func HEQSAnswerUpdate(db *gorm.DB, questionnaireShopAnswerId uint, updateShopAnswer *model.HouseEventQuestionnaireShopAnswer) *model.HouseEventQuestionnaireShopAnswer {
	db.Debug().Exec(`
UPDATE house_event_questionnaire_shop_answers SET ok = ? WHERE id = ? 
`, updateShopAnswer.Ok, questionnaireShopAnswerId)
	return updateShopAnswer
}

func HEQSAnswerDelete(db *gorm.DB, shopAnswerId uint) {
	shopAnswer := model.HouseEventQuestionnaireShopAnswer{}
	db.Delete(&shopAnswer, shopAnswerId)
}

func HEQSAnswerIsMine(db *gorm.DB, shopAnswerId uint, userId uint) bool {
	shopAnswer := model.HouseEventQuestionnaireShopAnswer{}
	shopAnswer.ID = shopAnswerId
	var count int
	db.Debug().Where("user_id = ?",userId).Find(&shopAnswer).Count(&count)

	fmt.Println("count")
	fmt.Println(count)
	if count > 0 {
		return true
	}
	return false
}
