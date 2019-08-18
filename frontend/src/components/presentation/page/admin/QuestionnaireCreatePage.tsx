import React from 'react';
import Header, { HeaderProps } from '../../Header';
import AddButton, { AddButtonProps } from '../../AddButton';
import CenterButton, { CenterButtonProps } from '../../CenterButton';

interface QuestionnaireCreatePageProps {
  header: HeaderProps;
  dateAddButton: AddButtonProps;
  shopAddButton: AddButtonProps;
  okButton: CenterButtonProps;
}

export default function QuestionnaireCreatePage(
  props: QuestionnaireCreatePageProps
) {
  return (
    <div>
      <Header {...props.header} />
      <AddButton {...props.dateAddButton} />
      <AddButton {...props.shopAddButton} />
      <CenterButton {...props.okButton} />
    </div>
  );
}
