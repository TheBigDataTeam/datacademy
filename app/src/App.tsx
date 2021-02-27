import React from 'react';
import './App.module.css'
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { LandingPage, SignUpPage, LoginPage, ForgetPassPage, Catalog, Authors } from './components/pages';

export const App: React.FunctionComponent = () => {
	return (
		<Router>
			<Route exact path="/" component={LandingPage} />
			<Route path="/auth/signup" component={SignUpPage} />
			<Route path="/auth/login" component={LoginPage} />
			<Route path="/auth/forget" component={ForgetPassPage} />
			<Route path="/catalog" component={Catalog} />
			<Route path="/authors" component={Authors}/>
		</Router>
	);
};
