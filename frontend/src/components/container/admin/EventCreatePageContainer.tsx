import React, { Component } from 'react';
import EventCreatePage from '../../presentation/page/admin/EventCreatePage';
import { ROUTE_ROOT } from '../../../Routes';

export default class EventCreatePageContainer extends Component {
  render() {
    return (
      <EventCreatePage
        centerButton={{ title: 'イベントを作成', linkTo: ROUTE_ROOT }}
        header={{ title: '新規イベント作成' }}
        textLabel="イベント名"
      />
    );
  }
}
