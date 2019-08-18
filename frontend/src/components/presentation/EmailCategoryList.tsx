import React, { ComponentType, ReactComponentElement, ReactNode } from 'react';
import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem, { ListItemProps } from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
      backgroundColor: theme.palette.background.paper
    }
  })
);

function ListItemLink(props: ListItemProps<'a', { button?: true }>) {
  return <ListItem button component="a" {...props} />;
}

interface EmailCategoryListProps {
  items: EmailCategoryListItem[];
}

export interface EmailCategoryListItem {
  icon: JSX.Element; // ガバガバな型だけどアイコンだからいいか
  text: string;
}

export default function EmailCategoryList(props: EmailCategoryListProps) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <List component="nav" aria-label="main mailbox folders">
        {renderListItems(props.items)}
      </List>
    </div>
  );
}

const renderListItems = (items: EmailCategoryListItem[]) => {
  return items.map((item, index) => (
    <ListItem key={index} button>
      <ListItemIcon>{item.icon}</ListItemIcon>
      <ListItemText primary={item.text} />
    </ListItem>
  ));
};
