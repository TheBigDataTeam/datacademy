import React from 'react';
import './App.module.css'
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { LandingPage, SignUpPage, LoginPage, 
		ForgetPassPage, CatalogPage, AuthorsPage, 
		ProjectInfoPage, PricingPage, PaymentPage,
		CoursePage 
} from 'components/pages';

export const App: React.FunctionComponent = () => {
	return (
		<Router>
			<Route exact path="/" component={LandingPage} />
			<Route path="/auth/signup" component={SignUpPage} />
			<Route path="/auth/login" component={LoginPage} />
			<Route path="/auth/forget" component={ForgetPassPage} />
			<Route exact path="/courses" component={CatalogPage} />
			<Route path="/courses/:id" component={CoursePage}/>
			<Route path="/authors" component={AuthorsPage}/>
			<Route path="/pricing" component={PricingPage}/>
			<Route path="/project" component={ProjectInfoPage}/>
			<Route path="/payment" component={PaymentPage}/>
		</Router>
	);
};
