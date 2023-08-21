import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import { LandingPage, CatalogPage, AuthorsPage, 
		ProjectInfoPage, PricingPage, PaymentPage,
		CoursePage, AuthorPage, DashboardPage, Page404, SignUpPage, 
		LoginPage, LogoutPage, ForgetPassPage } from 'components/pages'
import { AddAuthorPage, AddCoursePage, AddModulePage, AdminDashboardPage } from 'components/admin'

export const RoutesComponent: React.FunctionComponent = (): JSX.Element => {
	
	return (
		<Router>
			<Routes>
				<Route path="/" element={<LandingPage />}/>
				<Route path="/auth/signup" element={<SignUpPage />}/>
				<Route path="/auth/login" element={<LoginPage />}/>
				<Route path="/auth/logout" element={<LogoutPage />}/>
				<Route path="/auth/forget" element={<ForgetPassPage />}/>
				<Route path="/courses" element={<CatalogPage />}/>
				<Route path="/courses/:id" element={<CoursePage />}/>
				<Route path="/authors" element={<AuthorsPage />}/>
				<Route path="/authors/:id" element={<AuthorPage />}/>
				<Route path="/pricing" element={<PricingPage />}/>
				<Route path="/project" element={<ProjectInfoPage />}/>
				<Route path="/payment" element={<PaymentPage />}/>
				<Route path="/dashboard" element={<DashboardPage />}/>
				{/* Admin Section */}
				<Route path="/admin/add/course" element={<AddCoursePage />}/>
				<Route path="/admin/add/module" element={<AddModulePage />}/>
				<Route path="/admin/add/author" element={<AddAuthorPage />}/>
				<Route path="/admin/dashboard" element={<AdminDashboardPage />}/>
				{/* Page not found section */}
				<Route path="*" element={Page404}/>
			</Routes>
		</Router>
	)
}