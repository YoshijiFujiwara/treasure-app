import * as React from 'react';
import EventCard, { EventCardProps } from '../../EventCard';
import Header, { HeaderProps } from '../../Header';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Fab from '@material-ui/core/Fab';
import AddIcon from '@material-ui/icons/Add';
import { Grid } from '@material-ui/core';

interface HouseTopPageProps {
  header: HeaderProps;
  events: EventCardProps[];
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    fab: {
      margin: theme.spacing(1),
      marginLeft: theme.spacing(1)
    }
  })
);

export default function HouseTopPage(props: HouseTopPageProps) {
  const classes = useStyles();

  return (
    <div>
      <Header {...props.header} />
      {renderEvents(props.events)}
      <Grid container direction="row" justify="flex-end" alignItems="center">
        <Fab color="primary" aria-label="add" className={classes.fab}>
          <AddIcon />
        </Fab>
      </Grid>
    </div>
  );
}

const renderEvents = (eventData: EventCardProps[]) => {
  return eventData.map((event, index) => (
    <EventCard
      key={index}
      {...event}
    />
  ));
};
