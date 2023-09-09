import { LandingPage, CatalogPage, AuthorsPage, ProjectInfoPage, PricingPage, PaymentPage,
    CoursePage, AuthorPage, DashboardPage, Page404, SignUpPage, LoginPage, LogoutPage, ForgetPassPage } from 'components/pages'
import { AddAuthorPage, AddCoursePage, AddModulePage, AdminDashboardPage } from 'components/admin'

type Route = {
    path: string, 
    element: JSX.Element
}

type Routes = Route[]

export const routes: Routes = [{
    path: "/",
    element: <LandingPage />
}, {
    path: "/auth/signup",
    element: <SignUpPage />
}, {
    path: "/auth/login",
    element: <LoginPage />
}, {
    path: "/auth/logout",
    element: <LogoutPage />
}, {
    path: "/auth/forget",
    element: <ForgetPassPage />
}, {
    path: "/courses",
    element: <CatalogPage />
}, {
    path: "/courses/:id",
    element: <CoursePage />
}, {
    path: "/authors",
    element: <AuthorsPage />
}, {
    path: "/authors/:id",
    element: <AuthorPage />
}, {
    path: "/pricing",
    element: <PricingPage />
}, {
    path: "/project",
    element: <ProjectInfoPage />
}, {
    path: "/payment",
    element: <PaymentPage />
}, {
    path: "/dashboard",
    element: <DashboardPage />
}, {
    path: "/admin/add/course",
    element: <AddCoursePage />
}, {
    path: "/admin/add/module",
    element: <AddModulePage />
}, {
    path: "/admin/add/author",
    element: <AddAuthorPage />
}, {
    path: "/admin/dashboard",
    element: <AdminDashboardPage />
}, {
    path: "*",
    element: <Page404 />
}]
