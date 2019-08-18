import React from 'react';
import Header, { HeaderProps } from '../../Header';
import EmailCategoryList, {
  EmailCategoryListItem
} from '../../EmailCategoryList';

interface HouseTopPageProps {
  header: HeaderProps;
  items: EmailCategoryListItem[];
}

export default function HouseTopPage(props: HouseTopPageProps) {
  return (
    <div>
      <Header {...props.header} />
      <EmailCategoryList items={props.items} />
    </div>
  );
}
