import React from 'react';
import Header, { HeaderProps } from '../../Header';
import BigCard, { BigCardProps } from '../../BigCard';
import BigCenterButton, { BigCenterButtonProps } from '../../BigCenterButton';

interface DateDescriptionFirstPageProps {
  header: HeaderProps;
  bigCard: BigCardProps;
  nextButton: BigCenterButtonProps;
}

export default function DescriptionFirstPage(
  props: DateDescriptionFirstPageProps
) {
  return (
    <div>
      <Header {...props.header} />
      <BigCard {...props.bigCard} />
      <BigCenterButton {...props.nextButton} />
    </div>
  );
}
