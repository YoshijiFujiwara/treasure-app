import React from 'react';
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import ImageIcon from '@material-ui/icons/Image';
import WorkIcon from '@material-ui/icons/Work';
import BeachAccessIcon from '@material-ui/icons/BeachAccess';
import { JSXElement } from '@babel/types';
import { render } from 'react-dom';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
      backgroundColor: theme.palette.background.paper
    }
  })
);

export interface EmailListProps {
  items: ListItem[];
}

export interface ListItem {
  icon: JSX.Element;
  primaryText: string;
  secondaryText: string;
}

export default function EmailList(props: EmailListProps) {
  const classes = useStyles();

  return <List className={classes.root}>{renderListItems(props.items)}</List>;
}

const renderListItems = (items: ListItem[]) => {
  return items.map((item, index) => (
    <ListItem key={index}>
      <ListItemAvatar>
        <Avatar>{item.icon}</Avatar>
      </ListItemAvatar>
      <ListItemText primary={item.primaryText} secondary={item.secondaryText} />
    </ListItem>
  ));
};
