import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Input from '@material-ui/core/Input';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import FilledInput from '@material-ui/core/FilledInput';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
        },
        formControl: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
            minWidth: 300
        },
        selectEmpty: {
            marginTop: theme.spacing(2),
        },
    }),
);

export interface SimpleSelectProps {
    label: string,

}

export default function SimpleSelect(props: SimpleSelectProps) {
    const classes = useStyles();

    return (
        <FormControl className={classes.formControl}>
            <InputLabel htmlFor="age-simple">{props.label}</InputLabel>
            <Select>
                <MenuItem value={10}>Ten</MenuItem>
                <MenuItem value={20}>Twenty</MenuItem>
                <MenuItem value={30}>Thirty</MenuItem>
            </Select>
        </FormControl>
    );
}
