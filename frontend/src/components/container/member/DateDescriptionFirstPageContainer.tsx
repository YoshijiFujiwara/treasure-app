import React, { Component } from 'react';
import DescriptionFirstPage from '../../presentation/page/member/DescriptionFirstPage';

export default class DateDescriptionFirstPageContainer extends Component {
  render() {
    return (
      <DescriptionFirstPage
        header={{ title: '説明' }}
        bigCard={{ title: '日付の選択', body: '空いてる日を選択してください' }}
        nextButton={{ title: '進む', linkTo: 'hogehoge' }}
      />
    );
  }
}
