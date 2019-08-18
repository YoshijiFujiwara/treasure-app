import React, { Component, ComponentType } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Fab from '@material-ui/core/Fab';
import { Grid } from '@material-ui/core';
import { Link } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    fab: {
      margin: theme.spacing(1)
    },
    extendedIcon: {
      marginRight: theme.spacing(1)
    }
  })
);

export interface BigCenterButtonProps {
  title: string;
  icon?: JSX.Element;
  linkTo: string;
}

export default function BigCenterButton(props: BigCenterButtonProps) {
  const classes = useStyles();

  return (
    <Grid container alignItems="center" justify="center">
      <Fab variant="extended" aria-label="delete" className={classes.fab}>
        <Link to={props.linkTo}>
          {props.icon}
          <Typography variant="h2" component="h3">
            {props.title}
          </Typography>
        </Link>
      </Fab>
    </Grid>
  );
}
