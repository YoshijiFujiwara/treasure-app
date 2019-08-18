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

export interface EventCardProps {
    eventName: string,
    place: string,
    time: string
}

export default function EventCard(props: EventCardProps) {
    const classes = useStyles();
    const theme = useTheme();

    return (
        <Card className={classes.card}>
            <CardMedia
                className={classes.cover}
                image="https://via.placeholder.com/150"
                title="Live from space album cover"
            />
            <div className={classes.details}>
                <CardContent className={classes.content}>
                    <Typography component="h5" variant="h5">
                        {props.eventName}
                    </Typography>
                    <Typography variant="subtitle1" color="textSecondary">
                        {props.place}
                    </Typography>
                    <Typography variant="subtitle1" color="textSecondary">
                        {props.time}
                    </Typography>
                </CardContent>
            </div>
        </Card>
    );
}
