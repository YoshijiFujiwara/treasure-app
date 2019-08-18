import React from 'react';
import {
  createStyles,
  Grid,
  makeStyles,
  Theme,
  Typography
} from '@material-ui/core';
import DateDisplay, { DateDisplayProps } from '../../DateDisplay';
import Header, { HeaderProps } from '../../Header';
import ShopDisplay, {ShopDisplayProps} from '../../ShopDisplay';
import CenterButton, { CenterButtonProps } from '../../CenterButton';
import Fab from '@material-ui/core/Fab';
import { Cancel } from '@material-ui/icons';

interface EventDetailPageProps {
  header: HeaderProps;
  dateDisplay: DateDisplayProps;
  shopDisplay: ShopDisplayProps;
  questionnaireCreateButton: CenterButtonProps;
  questionnaireResultButton: CenterButtonProps;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    fab: {
      margin: theme.spacing(1),
      marginLeft: theme.spacing(1)
    }
  })
);

export default function EventDetailPage(props: EventDetailPageProps) {
  const classes = useStyles();

  return (
    <div>
      <Header title={props.header.title} />
      <div style={{ paddingTop: 30 }}>
        <DateDisplay {...props.dateDisplay} />
        <ShopDisplay {...props.shopDisplay} />
        <CenterButton {...props.questionnaireCreateButton} />
        <CenterButton {...props.questionnaireResultButton} />
        <Grid container direction="row" justify="center" alignItems="center">
          <Fab aria-label="add" className={classes.fab}>
            ×
          </Fab>
        </Grid>
      </div>
    </div>
  );
}
