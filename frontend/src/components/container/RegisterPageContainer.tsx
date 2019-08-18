import React, { Component } from 'react';
import RegisterPage from '../presentation/page/RegisterPage';
import { ROUTE_ADMIN_HOUSE_TOP } from '../../Routes';

export default class RegisterPageContainer extends Component {
  render() {
    return (
      <RegisterPage
        nameText={{ label: 'ユーザーネーム' }}
        emailText={{ label: 'メールアドレス' }}
        passwordText={{ label: 'パスワード' }}
        genderSelect={{ label: '性別' }}
        birthdaySelect={{ label: '誕生日' }}
        registerButton={{
          title: 'この情報で作成する',
          linkTo: ROUTE_ADMIN_HOUSE_TOP
        }}
      />
    );
  }
}
