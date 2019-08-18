import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        button: {
            margin: theme.spacing(1),
        },
        input: {
            display: 'none',
        },
    }),
);

export interface AddButtonProps {
    label: string
}

export default function AddButton(props: AddButtonProps) {
    const classes = useStyles();

    return (
        <div>
            <Button color="primary" className={classes.button}>
                {props.label}
            </Button>
        </div>
    );
}
