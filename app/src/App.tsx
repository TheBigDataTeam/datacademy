import React from 'react'
import './App.module.css'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { LandingPage, SignUpPage, LoginPage, 
		ForgetPassPage, CatalogPage, AuthorsPage, 
		ProjectInfoPage, PricingPage, PaymentPage,
		CoursePage, AuthorPage, DashboardPage
} from 'components/pages'

export const App: React.FunctionComponent = (): JSX.Element => {
	return (
		<Router>
			<Route exact path="/" component={LandingPage} />
			<Route path="/auth/signup" component={SignUpPage} />
			<Route path="/auth/login" component={LoginPage} />
			<Route path="/auth/forget" component={ForgetPassPage} />
			<Route exact path="/courses" component={CatalogPage} />
			<Route path="/courses/:id" component={CoursePage}/>
			<Route exact path="/authors" component={AuthorsPage}/>
			<Route path="/authors/:id" component={AuthorPage}/>
			<Route path="/pricing" component={PricingPage}/>
			<Route path="/project" component={ProjectInfoPage}/>
			<Route path="/payment" component={PaymentPage}/>
			<Route path="/dashboard" component={DashboardPage}/>
		</Router>
	)
}
