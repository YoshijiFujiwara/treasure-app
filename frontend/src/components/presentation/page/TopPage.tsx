import React, { Component } from 'react';
import TopLogo from '../TopLogo';
import CenterButton, { CenterButtonProps } from '../CenterButton';
import { Grid, Typography } from '@material-ui/core';
import logo from '../../../img/logo.png';

interface TopPageProps {
  appName: string;
  joinButton: CenterButtonProps;
  createButton: CenterButtonProps;
}

export default function TopPage(props: TopPageProps) {
  return (
    <Grid container alignItems="center" justify="center">
      <Typography align="center" variant="h4">
        {props.appName}
      </Typography>
      <TopLogo logo={logo} />
      <CenterButton {...props.joinButton} />
      <CenterButton {...props.createButton} />
    </Grid>
  );
}
