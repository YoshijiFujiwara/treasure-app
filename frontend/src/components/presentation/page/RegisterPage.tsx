import React from 'react';
import TextInputField, { TextInputFieldProps } from '../TextInputField';
import SimpleSelect, { SimpleSelectProps } from '../SimpleSelect';
import CenterButton, { CenterButtonProps } from '../CenterButton';

interface RegisterPageProps {
  nameText: TextInputFieldProps;
  emailText: TextInputFieldProps;
  passwordText: TextInputFieldProps;
  genderSelect: SimpleSelectProps;
  birthdaySelect: SimpleSelectProps;
  registerButton: CenterButtonProps;
}

export default function RegisterPage(props: RegisterPageProps) {
  return (
    <div>
      <TextInputField label={props.nameText.label} />
      <TextInputField label={props.emailText.label} />
      <TextInputField label={props.passwordText.label} />
      <SimpleSelect label={props.genderSelect.label} />
      <SimpleSelect label={props.birthdaySelect.label} />
      <CenterButton
        title={props.registerButton.title}
        linkTo={props.registerButton.linkTo}
      />
    </div>
  );
}
