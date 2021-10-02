import React from 'react'
import { AuthLayout } from 'components/layouts'
import { Buttons } from './components'

export const LogoutPage: React.FunctionComponent = (): JSX.Element => {
    return (
        <AuthLayout>
            <h3>Are you sure you want to log out?</h3>
            <Buttons />
        </AuthLayout>
    )
}
