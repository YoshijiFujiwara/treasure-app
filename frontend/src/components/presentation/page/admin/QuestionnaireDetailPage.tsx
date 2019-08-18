import React from 'react';
import Header, { HeaderProps } from '../../Header';
import Typography from '@material-ui/core/Typography';
import CenterButton, { CenterButtonProps } from '../../CenterButton';
import DateDisplay, { DateDisplayProps } from '../../DateDisplay';
import ShopDisplay, { ShopDisplayProps } from '../../ShopDisplay';

interface QuestionnaireDetailPageProps {
  header: HeaderProps;
  dates: DateDisplayProps[];
  shops: ShopDisplayProps[];
  sendButton: CenterButtonProps;
}

export default function QuestionnaireDetailPage(
  props: QuestionnaireDetailPageProps
) {
  return (
    <div>
      <Header {...props.header} />
      <Typography component="h5" variant="h5">
        日付リスト
      </Typography>
      {renderDates(props.dates)}
      <Typography component="h5" variant="h5">
        お店リスト
      </Typography>
      {renderShops(props.shops)}
      <CenterButton {...props.sendButton} />
    </div>
  );
}

const renderDates = (dates: DateDisplayProps[]) => {
  return dates.map((date, index) => (
    <DateDisplay key={index} label={date.label} />
  ));
};

const renderShops = (shops: ShopDisplayProps[]) => {
  return shops.map((shop, index) => (
    <ShopDisplay key={index} name={shop.name} />
  ));
};
