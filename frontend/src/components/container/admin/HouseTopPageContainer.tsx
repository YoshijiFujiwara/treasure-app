import React, { Component } from 'react';
import HouseTopPage from '../../presentation/page/admin/HouseTopPage';
import { EventCardProps } from '../../presentation/EventCard';

export default class HouseTopPageContainer extends Component {
  render() {
    const dummyEventData: EventCardProps[] = [
      {
        eventName: 'イベント1',
        place: '渋谷',
        time: '2019-08-22'
      },
      {
        eventName: 'イベント2',
        place: '浅草',
        time: '2019-08-23'
      },
      {
        eventName: 'イベント3',
        place: 'お台場',
        time: '2019-08-24'
      },
      {
        eventName: 'イベント4',
        place: '大阪',
        time: '2019-08-25'
      },
      {
        eventName: 'イベント5',
        place: '名古屋',
        time: '2019-08-27'
      }
    ];

    return (
      <HouseTopPage header={{ title: 'トップ' }} events={dummyEventData} />
    );
  }
}
