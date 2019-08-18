import React, {Component} from 'react';
import QuestionnaireCreatePage from "../../presentation/page/admin/QuestionnaireCreatePage";
import {ROUTE_ADMIN_HOUSE_EVENT_QUES} from "../../../Routes";

export default class QuestionnaireCreatePageContainer extends Component {
    render() {
        return (
            <QuestionnaireCreatePage
                header={{title: 'アンケート作成'}}
                dateAddButton={{label: '日程を追加'}}
                shopAddButton={{label: '店を追加'}}
                okButton={{title: 'このリストで決定', linkTo: ROUTE_ADMIN_HOUSE_EVENT_QUES }}
            />
        )
    }
}
