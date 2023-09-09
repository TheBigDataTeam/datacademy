import { LandingPage, CatalogPage, AuthorsPage, ProjectInfoPage, PricingPage, PaymentPage,
    CoursePage, AuthorPage, DashboardPage, Page404, SignUpPage, LoginPage, LogoutPage, ForgetPassPage } from 'components/pages'
import { AddAuthorPage, AddCoursePage, AddModulePage, AdminDashboardPage } from 'components/admin'
import { HOMEPAGE_URL, LOGIN_URL, LOGOUT_URL, REGISTRATION_URL, FORGET_PASS_URL, 
    ADMIN_HOMEPAGE_URL, ADMIN_ADD_COURSE_URL, ADMIN_ADD_MODULE_URL, ADMIN_ADD_AUTHOR_URL,
    COURSES_URL, AUTHORS_URL, PRICING_URL, PROJECT_URL, PAYMENT_URL, USER_DASHBOARD_URL, NOT_FOUND_URL } from 'constants/urls'

type Route = {
    path: string, 
    element: JSX.Element
}

type Routes = Route[]

export const routes: Routes = [
    {
        path: HOMEPAGE_URL,
        element: <LandingPage />
    }, 
    {
        path: REGISTRATION_URL,
        element: <SignUpPage />
    }, 
    {
        path: LOGIN_URL,
        element: <LoginPage />
    }, 
    {
        path: LOGOUT_URL,
        element: <LogoutPage />
    }, 
    {
        path: FORGET_PASS_URL,
        element: <ForgetPassPage />
    }, 
    {
        path: COURSES_URL,
        element: <CatalogPage />
    }, 
    {
        path: COURSES_URL + "/:id",
        element: <CoursePage />
    }, 
    {
        path: AUTHORS_URL,
        element: <AuthorsPage />
    }, 
    {
        path: AUTHORS_URL + "/:id",
        element: <AuthorPage />
    }, 
    {
        path: PRICING_URL,
        element: <PricingPage />
    }, 
    {
        path: PROJECT_URL,
        element: <ProjectInfoPage />
    }, 
    {
        path: PAYMENT_URL,
        element: <PaymentPage />
    }, 
    {
        path: USER_DASHBOARD_URL,
        element: <DashboardPage />
    }, 
    {
        path: ADMIN_ADD_COURSE_URL,
        element: <AddCoursePage />
    }, 
    {
        path: ADMIN_ADD_MODULE_URL,
        element: <AddModulePage />
    }, 
    {
        path: ADMIN_ADD_AUTHOR_URL,
        element: <AddAuthorPage />
    }, 
    {
        path: ADMIN_HOMEPAGE_URL,
        element: <AdminDashboardPage />
    }, 
    {
        path: NOT_FOUND_URL,
        element: <Page404 />
    }
]
