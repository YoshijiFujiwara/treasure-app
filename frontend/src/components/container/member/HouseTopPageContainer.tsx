import React, { Component } from 'react';
import HouseTopPage from '../../presentation/page/member/HouseTopPage';
import { EmailCategoryListItem } from '../../presentation/EmailCategoryList';
import InboxIcon from '@material-ui/core/SvgIcon/SvgIcon';

export default class HouseTopPageContainer extends Component {
  render() {
    return (
      <HouseTopPage header={{ title: 'ホゲホゲ' }} items={dummyListItems} />
    );
  }
}

const dummyListItems: EmailCategoryListItem[] = [
  {
    icon: <InboxIcon />,
    text: '受信アンケート'
  },
  {
    icon: <InboxIcon />,
    text: 'アンケート結果'
  }
];
