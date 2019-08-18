import React, { Component } from 'react';
import DescriptionSecondPage from '../../presentation/page/member/DescriptionSecondPage';

export default class DateDescriptionSecondPageContainer extends Component {
  render() {
    return (
      <DescriptionSecondPage
        header={{ title: 'アンケート説明' }}
        okText="空いている日なら"
        ngText="空いてない日なら"
        nextButton={{ title: '進む', linkTo: 'hogehoge' }}
      />
    );
  }
}
