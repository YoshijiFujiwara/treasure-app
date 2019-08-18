import React from 'react';
import EmailList, { EmailListProps } from '../../EmailList';
import Header, { HeaderProps } from '../../Header';

interface ReceiveQuestionnairesPageProps {
  header: HeaderProps;
  listItem: EmailListProps;
}

export default function ReceiveQuestionnairesPage(
  props: ReceiveQuestionnairesPageProps
) {
  return (
    <div>
      <Header title={props.header.title} />
      <EmailList items={props.listItem.items} />
    </div>
  );
}
