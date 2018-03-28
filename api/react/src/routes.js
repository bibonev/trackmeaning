import React from 'react';
import {Router, Route, IndexRoute} from 'react-router';
import App from './components/App.js';

export default(
    <Router>
        <Route path="/" component={App}>
        </Route>
    </Router>
);