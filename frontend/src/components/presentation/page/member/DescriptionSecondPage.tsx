import React from 'react';
import Header, { HeaderProps } from '../../Header';
import { Typography } from '@material-ui/core';
import BigOk from '../../BigOk';
import BigNg from '../../BigNg';
import BigCenterButton, { BigCenterButtonProps } from '../../BigCenterButton';

interface DateDescriptionSecondPageProps {
  header: HeaderProps;
  okText: string;
  ngText: string;
  nextButton: BigCenterButtonProps;
}

export default function DescriptionSecondPage(
  props: DateDescriptionSecondPageProps
) {
  return (
    <div>
      <Header {...props.header} />
      <Typography variant="h1" component="h2">
        {props.okText}
      </Typography>
      <BigOk />
      <Typography variant="h1" component="h2">
        {props.ngText}
      </Typography>
      <BigNg />
      <BigCenterButton {...props.nextButton} />
    </div>
  );
}
