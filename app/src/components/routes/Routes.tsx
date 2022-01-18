import React from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import { LandingPage, CatalogPage, AuthorsPage, 
		ProjectInfoPage, PricingPage, PaymentPage,
		CoursePage, AuthorPage, DashboardPage, Page404, SignUpPage, 
		LoginPage, LogoutPage, ForgetPassPage } from 'components/pages'
import { AddAuthorPage, AddCoursePage, AddModulePage, AdminDashboardPage } from 'components/admin'

export const Routes: React.FunctionComponent = (): JSX.Element => {
	
	return (
		<Router>
			<Switch>
				<Route exact path="/" component={LandingPage}/>
				<Route path="/auth/signup" component={SignUpPage}/>
				<Route path="/auth/login" component={LoginPage}/>
				<Route path="/auth/logout" component={LogoutPage}/>
				<Route path="/auth/forget" component={ForgetPassPage}/>
				<Route exact path="/courses" component={CatalogPage}/>
				<Route path="/courses/:id" component={CoursePage}/>
				<Route exact path="/authors" component={AuthorsPage}/>
				<Route path="/authors/:id" component={AuthorPage}/>
				<Route path="/pricing" component={PricingPage}/>
				<Route path="/project" component={ProjectInfoPage}/>
				<Route path="/payment" component={PaymentPage}/>
				<Route path="/dashboard" component={DashboardPage}/>
				{/* Admin Section */}
				<Route path="/admin/add/course" component={AddCoursePage}/>
				<Route path="/admin/add/module" component={AddModulePage}/>
				<Route path="/admin/add/author" component={AddAuthorPage}/>
				<Route path="/admin/dashboard" component={AdminDashboardPage}/>
				{/* Page not found section */}
				<Route component={Page404}/>
			</Switch>
		</Router>
	)
}