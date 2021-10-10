import React, { useEffect } from 'react'
import './App.module.css'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { LandingPage, SignUpPage, LoginPage, LogoutPage,
		ForgetPassPage, CatalogPage, AuthorsPage, 
		ProjectInfoPage, PricingPage, PaymentPage,
		CoursePage, AuthorPage, DashboardPage
} from 'components/pages'
import { AddAuthorPage } from 'components/admin'
import { useDispatch } from 'react-redux'
import { fetchUserLogin } from 'redux/user_auth/userAuthActions'

export const App: React.FunctionComponent = (): JSX.Element => {

	const dispatch = useDispatch()
	//const sessionID = document.cookie.split(";").find(cookie => cookie.startsWith(" session_id=")).split("=")[1]

	useEffect(() => {
		try {
			dispatch(fetchUserLogin())
		} catch (error) {
			console.log(error) /* TODO handle errors properly */
		}
	}, [dispatch])
	
	return (
		<Router>
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
			<Route path="/admin/add/course" />
			<Route path="/admin/add/author" component={AddAuthorPage}/>
		</Router>
	)
}
