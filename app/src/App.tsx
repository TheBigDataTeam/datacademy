import React, { useEffect } from 'react'
import './App.module.css'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import { LandingPage, SignUpPage, LoginPage, LogoutPage,
		ForgetPassPage, CatalogPage, AuthorsPage, 
		ProjectInfoPage, PricingPage, PaymentPage,
		CoursePage, AuthorPage, DashboardPage, Page404 } from 'components/pages'
import { AddAuthorPage, AdminDashboardPage } from 'components/admin'
import { useDispatch } from 'react-redux'
import { fetchUserLogin } from 'redux/user_auth/userAuthActions'

export const App: React.FunctionComponent = (): JSX.Element => {

	const dispatch = useDispatch()

	useEffect(() => {
		try {
			dispatch(fetchUserLogin())
		} catch (error) {
			console.log(error) /* TODO handle errors properly */
		}
	}, [dispatch])
	
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
				<Route component={Page404}/>
				{/* Admin Section */}
				<Route path="/admin/add/course" />
				<Route path="/admin/add/author" component={AddAuthorPage}/>
				<Route path="/admin/dashboard" component={AdminDashboardPage}/>
			</Switch>
		</Router>
	)
}
