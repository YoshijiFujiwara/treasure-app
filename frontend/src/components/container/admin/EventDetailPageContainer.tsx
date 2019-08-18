import React, { Component } from 'react';
import EventDetailPage from '../../presentation/page/admin/EventDetailPage';
import {
  ROUTE_ADMIN_HOUSE_EVENT_QUES_CREATE,
  ROUTE_ADMIN_HOUSE_EVENT_QUES
} from '../../../Routes';

export default class EventDetailPageContainer extends Component {
  render() {
    return (
      <EventDetailPage
        header={{ title: 'イベント詳細' }}
        dateDisplay={{ label: '予定日時' }}
        shopDisplay={{ name: 'ホゲホゲラーメン' }}
        questionnaireCreateButton={{
          title: 'アンケートを作成',
          linkTo: ROUTE_ADMIN_HOUSE_EVENT_QUES_CREATE
        }}
        questionnaireResultButton={{
          title: 'アンケート結果',
          linkTo: ROUTE_ADMIN_HOUSE_EVENT_QUES
        }}
      />
    );
  }
}
