import React from 'react';
import ReactDOM from 'react-dom';
import {Router} from 'react-router';
import store from './store'
import {Provider} from 'react-redux';
import {createBrowserHistory} from 'history';
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import {Routes} from './Routes'
import {red} from "@material-ui/core/colors";

const history = createBrowserHistory();

// カラースキーマはここで設定しようか
const theme = createMuiTheme({
        palette: {
            primary: {
                main: '#F1A747'
            },
            secondary: {
                main: '#FFD05B'
            },
            error: red,
        }
    },
);

ReactDOM.render(
    <MuiThemeProvider theme={theme}>
        <Provider store={store}>
            <Router history={history}>
                <Routes/>
            </Router>
        </Provider>
    </MuiThemeProvider>
    , document.getElementById('root')
);
