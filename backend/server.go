package backend

import (
	"fmt"
	"github.com/YoshijiFujiwara/u22/backend/controller"
	"github.com/YoshijiFujiwara/u22/backend/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	db     *gorm.DB
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(dbDriver string, dataSource string) {
	// db
	dbcon, err := gorm.Open(dbDriver, dataSource)
	if err != nil {
		log.Fatal("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.db)
	houseMemberMiddleware := middleware.NewHouseMember(s.db)
	houseAdminMiddleware := middleware.NewHouseAdmin(s.db)
	myDateAnswerMiddleware := middleware.NewMyDateAnswer(s.db)
	myShopAnswerMiddleware := middleware.NewMyShopAnswer(s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
	})

	__0__commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)
	// 正常なトークンを持っていないとダメだよ
	__1__authChain := __0__commonChain.Append(
		authMiddleware.Handler,
	)
	// そのハウスに所属していないとダメだよ
	__2__isMemberChain := __1__authChain.Append(
		houseMemberMiddleware.Handler,
	)
	// そのハウスのadminじゃないとダメだよ
	__3__isAdminChain := __2__isMemberChain.Append(
		houseAdminMiddleware.Handler,
	)
	// その日付アンケートの回答者本人じゃないとダメだよ
	__2_1__isMyDateAnswerChain := __2__isMemberChain.Append(
		myDateAnswerMiddleware.Handler,
	)
	// その店アンケートの回答者本人じゃないとダメだよ
	__2_1__isMyShopAnswerChain := __2__isMemberChain.Append(
		myShopAnswerMiddleware.Handler,
	)

	r := mux.NewRouter()

	apiPrefix := os.Getenv("API_PREFIX")

	userController := controller.NewUser(s.db)
	houseController := controller.NewHouse(s.db)
	houseUserController := controller.NewHouseUser(s.db)
	houseEventController := controller.NewHouseEvent(s.db)
	houseEventQuestionnaireController := controller.NewHouseEventQuestionnaire(s.db)
	houseEventQuestionnaireDateController := controller.NewHEQDate(s.db)
	houseEventQuestionnaireDateAnswerController := controller.NewHEQDAnswer(s.db)
	houseEventQuestionnaireShopController := controller.NewHEQShop(s.db)
	houseEventQuestionnaireShopAnswerController := controller.NewHEQSAnswer(s.db)

	// ユーザーの登録とか
	r.Methods(http.MethodPost).Path(apiPrefix + "/users").Handler(__0__commonChain.Then(AppHandler{userController.Register}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/users/login").Handler(__0__commonChain.Then(AppHandler{userController.Login}))
	r.Methods(http.MethodPut).Path(apiPrefix + "/users/update_me").Handler(__1__authChain.Then(AppHandler{userController.Update}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/users/delete_me").Handler(__1__authChain.Then(AppHandler{userController.Destroy}))

	// ハウスの登録とか
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses").Handler(__1__authChain.Then(AppHandler{houseController.Create}))
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}").Handler(__2__isMemberChain.Then(AppHandler{houseController.Show}))

	// ハウスのユーザーの管理
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/users/{user_id}").Handler(__3__isAdminChain.Then(AppHandler{houseUserController.Add}))
	r.Methods(http.MethodPut).Path(apiPrefix + "/houses/{house_id}/users/{user_id}").Handler(__3__isAdminChain.Then(AppHandler{houseUserController.Update}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/users/{user_id}").Handler(__3__isAdminChain.Then(AppHandler{houseUserController.Destroy}))

	// イベント
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events").Handler(__2__isMemberChain.Then(AppHandler{houseEventController.Index}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/events").Handler(__2__isMemberChain.Then(AppHandler{houseEventController.Create}))
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}").Handler(__2__isMemberChain.Then(AppHandler{houseEventController.Show}))
	r.Methods(http.MethodPut).Path(apiPrefix + "/houses/{house_id}/events/{event_id}").Handler(__3__isAdminChain.Then(AppHandler{houseEventController.Update}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/events/{event_id}").Handler(__3__isAdminChain.Then(AppHandler{houseEventController.Delete}))

	// アンケート
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireController.Index}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires").Handler(__3__isAdminChain.Then(AppHandler{houseEventQuestionnaireController.Create}))
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireController.Show}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}").Handler(__3__isAdminChain.Then(AppHandler{houseEventQuestionnaireController.Delete}))

	// アンケート項目　日
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireDateController.Index}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates").Handler(__3__isAdminChain.Then(AppHandler{houseEventQuestionnaireDateController.Create}))
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates/{questionnaire_date_id}").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireDateController.Show}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates/{questionnaire_date_id}").Handler(__3__isAdminChain.Then(AppHandler{houseEventQuestionnaireDateController.Delete}))

	// アンケート項目　店
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireShopController.Index}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops").Handler(__3__isAdminChain.Then(AppHandler{houseEventQuestionnaireShopController.Create}))
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops/{questionnaire_shop_id}").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireShopController.Show}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops/{questionnaire_shop_id}").Handler(__3__isAdminChain.Then(AppHandler{houseEventQuestionnaireShopController.Delete}))

	// アンケートの答え　日
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates/{questionnaire_date_id}/answers").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireDateAnswerController.Index}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates/{questionnaire_date_id}/answers").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireDateAnswerController.Create}))
	r.Methods(http.MethodPut).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates/{questionnaire_date_id}/answers/{questionnaire_date_answer_id}").Handler(__2_1__isMyDateAnswerChain.Then(AppHandler{houseEventQuestionnaireDateAnswerController.Update}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/dates/{questionnaire_date_id}/answers/{questionnaire_date_answer_id}").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireDateAnswerController.Delete}))

	// アンケート項目　店
	r.Methods(http.MethodGet).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops/{questionnaire_shop_id}/answers").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireShopAnswerController.Index}))
	r.Methods(http.MethodPost).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops/{questionnaire_shop_id}/answers").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireShopAnswerController.Create}))
	r.Methods(http.MethodPut).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops/{questionnaire_shop_id}/answers/{questionnaire_shop_answer_id}").Handler(__2_1__isMyShopAnswerChain.Then(AppHandler{houseEventQuestionnaireShopAnswerController.Update}))
	r.Methods(http.MethodDelete).Path(apiPrefix + "/houses/{house_id}/events/{event_id}/questionnaires/{questionnaire_id}/shops/{questionnaire_shop_id}/answers/{questionnaire_shop_answer_id}").Handler(__2__isMemberChain.Then(AppHandler{houseEventQuestionnaireShopAnswerController.Delete}))

	return r
}
