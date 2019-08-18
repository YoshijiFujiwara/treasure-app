import React from 'react';
import { Theme, createStyles, makeStyles, useTheme } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        card: {
            display: 'flex',
            margin: 10
        },
        details: {
            display: 'flex',
            flexDirection: 'column',
        },
        content: {
            flex: '1 0 auto',
        },
        cover: {
            width: 151,
        },
    }),
);

export interface ShopDisplayProps {
    name: string
}

export default function ShopDisplay(props: ShopDisplayProps) {
    const classes = useStyles();

    return (
        <Card className={classes.card}>
            <CardMedia
                className={classes.cover}
                image="https://via.placeholder.com/150"
                title="Live from space album cover"
            />
            <div className={classes.details}>
                <CardContent className={classes.content}>
                    <Typography variant="subtitle1" color="textSecondary">
                        {props.name}
                    </Typography>
                </CardContent>
            </div>
        </Card>
    );
}
