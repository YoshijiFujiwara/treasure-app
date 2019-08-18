import React from 'react';
import {Grid} from "@material-ui/core";

interface TopLogoProps {
    logo: string
}

export default (props: TopLogoProps) => {
    return (
        <Grid
            container
            spacing={0}
            direction="column"
            alignItems="center"
            justify="center"
            style={{minHeight: '30vh'}}
        >
            <img src={props.logo} alt="ロゴ"/>
        </Grid>
    );
}
