import React, {Component} from 'react';
import QuestionnaireResultPage from '../../presentation/page/admin/QuestionnaireResultPage';
import {QuestionnaireCardProps} from "../../presentation/QuestionnaireCard";

export default class QuestionnaireResultPageContainer extends Component {
    render(){
        return (
            <QuestionnaireResultPage header={{title: 'アンケート結果'}} questionnires={questionnaireDummys}/>
        )
    }
}

const questionnaireDummys: QuestionnaireCardProps[] = [
    {
        name: 'アンケートA'
    },
    {
        name: 'アンケートB'
    },
    {
        name: 'アンケートC'
    },
    {
        name: 'アンケートD'
    },
    {
        name: 'アンケートE'
    },
];
