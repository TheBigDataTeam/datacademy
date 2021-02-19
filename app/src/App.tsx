import React from 'react';
import './App.css'
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { LandingPage, SignUpPage, LoginPage, Catalog, Authors } from './components/pages';

export const App: React.FunctionComponent = () => {
	return (
		<Router>
			<Route exact path="/" component={LandingPage} />
			<Route path="/signup" component={SignUpPage} />
			<Route path="/login" component={LoginPage} />
			<Route path="/catalog" component={Catalog} />
			<Route path="/authors" component={Authors}/>
		</Router>
	);
};
