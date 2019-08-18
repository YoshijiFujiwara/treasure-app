import React, { Component } from 'react';
import ReceiveQuestionnairesPage from '../../presentation/page/member/ReceiveQuestionnairesPage';
import { EmailListProps, ListItem } from '../../presentation/EmailList';
import InboxIcon from '@material-ui/core/SvgIcon/SvgIcon';

export default class ReceiveQuestionnairesPageContainer extends Component {
  render() {
    return (
      <ReceiveQuestionnairesPage
        header={{ title: '受信アンケート一覧' }}
        listItem={dummyItem}
      />
    );
  }
}

const dummyItem: EmailListProps = {
  items: [
    {
      icon: <InboxIcon />,
      primaryText: 'アンケートA',
      secondaryText: 'その日時'
    },
    {
      icon: <InboxIcon />,
      primaryText: 'アンケートB',
      secondaryText: 'その日時'
    },
    {
      icon: <InboxIcon />,
      primaryText: 'アンケートC',
      secondaryText: 'その日時'
    },
    {
      icon: <InboxIcon />,
      primaryText: 'アンケートD',
      secondaryText: 'その日時'
    }
  ]
};
