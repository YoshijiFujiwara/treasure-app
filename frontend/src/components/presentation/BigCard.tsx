import React from 'react';
import { Theme, createStyles, makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing(3, 2)
    }
  })
);

export interface BigCardProps {
  title: string;
  body: string;
}

export default function BigCard(props: BigCardProps) {
  const classes = useStyles();

  return (
    <div>
      <Paper className={classes.root}>
        <Typography variant="h3" component="h4">
          {props.title}
        </Typography>
        <Typography variant="h4" component="h4">
          {props.body}
        </Typography>
      </Paper>
    </div>
  );
}
