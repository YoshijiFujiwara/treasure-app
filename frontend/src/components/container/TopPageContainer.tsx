import React, { Component } from 'react';
import TopPage from '../presentation/page/TopPage';
import { ROUTE_HOUSE_CREATE, ROUTE_HOUSE_LOGIN } from '../../Routes';
import { Add, Home } from '@material-ui/icons';
import BigOk from "../presentation/BigOk";
import BigNg from "../presentation/BigNg";

export default class TopPageContainer extends Component {
  render() {
    return (
      <div>
        <TopPage
          appName="ホイッスル"
          joinButton={{
            title: 'ハウスに参加する',
            icon: <Home />,
            linkTo: ROUTE_HOUSE_LOGIN
          }}
          createButton={{
            title: '新しいハウスを作成する',
            icon: <Add />,
            linkTo: ROUTE_HOUSE_CREATE
          }}
        />
        <BigNg/>
      </div>
    );
  }
}
