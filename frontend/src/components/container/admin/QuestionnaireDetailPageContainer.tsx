import React, { Component } from 'react';
import QuestionnaireDetailPage from '../../presentation/page/admin/QuestionnaireDetailPage';
import { ROUTE_ADMIN_HOUSE_EVENT_QUES } from '../../../Routes';
import { DateDisplayProps } from '../../presentation/DateDisplay';
import { ShopDisplayProps } from '../../presentation/ShopDisplay';

export default class QuestionnaireDetailPageContainer extends Component {
  render() {
    return (
      <QuestionnaireDetailPage
        header={{ title: 'アンケート詳細' }}
        sendButton={{
          title: 'このアンケートを送信',
          linkTo: ROUTE_ADMIN_HOUSE_EVENT_QUES
        }}
        dates={dummyDates}
        shops={dummyShops}
      />
    );
  }
}

const dummyDates: DateDisplayProps[] = [
  {
    label: '2019-08-23 12:00:00'
  },
  {
    label: '2019-08-23 12:00:00'
  },
  {
    label: '2019-08-13 12:00:00'
  },
  {
    label: '2019-08-30 12:00:00'
  },
  {
    label: '2019-08-10 12:00:00'
  }
];

const dummyShops: ShopDisplayProps[] = [
  {
    name: 'ホゲホゲラーメン'
  },
  {
    name: 'ふがふがそば'
  },
  {
    name: 'ホゲホゲマック'
  }
];
