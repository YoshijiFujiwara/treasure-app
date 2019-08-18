import React from 'react';
import {makeStyles, createStyles, Theme} from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import {Grid} from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        container: {
            display: 'flex',
            flexWrap: 'wrap',
        },
        textField: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1)
        },
        dense: {
            marginTop: 19,
        },
        menu: {
            width: 200,
        },
    }),
);

export interface TextInputFieldProps {
    label: string
}

export default function TextInputField(props: TextInputFieldProps) {
    const classes = useStyles();

    return (
        <form noValidate style={{'width': '100%'}} autoComplete="off">
            <TextField
                label={props.label}
                // value={values.name}
                // onChange={handleChange('name')}
            />
        </form>
    );
}
