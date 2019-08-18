import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { Button } from '@material-ui/core';
import { Clear } from '@material-ui/icons';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    fab: {
      margin: theme.spacing(1),
      height: 400,
      width: '45%'
    },
    extendedIcon: {
      marginRight: theme.spacing(1)
    }
  })
);

export default function BigNg() {
  const classes = useStyles();

  return (
    <div>
      <Button
        style={{ backgroundColor: '#ff6666' }}
        aria-label="add"
        className={classes.fab}
      >
        <Clear style={{ color: 'white', fontSize: 300 }} />
      </Button>
    </div>
  );
}
