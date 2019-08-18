import Header, { HeaderProps } from '../Header';
import React, { Component } from 'react';
import { Grid } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';
import CenterButton, { CenterButtonProps } from '../CenterButton';

interface CreateHousePageProps {
  headerTitle: HeaderProps;
  centerButton: CenterButtonProps;
  textLabel: string;
}

export default function CreateHousePage(props: CreateHousePageProps) {
  return (
    <div>
      <Header title={props.headerTitle.title} />
      <Grid container alignItems="center" justify="center">
        <form style={{ width: '90%', marginTop: '20px' }}>
          <TextField fullWidth={true} label={props.textLabel} />
        </form>
        <CenterButton {...props.centerButton} />
      </Grid>
    </div>
  );
}
