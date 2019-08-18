import React from 'react';
import Header, { HeaderProps } from '../../Header';
import { Grid } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';
import CenterButton, { CenterButtonProps } from '../../CenterButton';

interface EventCreatePageProps {
  header: HeaderProps;
  centerButton: CenterButtonProps;
  textLabel: string;
}

export default function EventCreatePage(props: EventCreatePageProps) {
  return (
    <div>
      <Header {...props.header} />
      <Grid container={true} alignItems="center" justify="center">
        <form style={{ width: '90%', marginTop: '20px' }}>
          <TextField fullWidth={true} label={props.textLabel} />
        </form>
        <CenterButton {...props.centerButton} />
      </Grid>
    </div>
  );
}
