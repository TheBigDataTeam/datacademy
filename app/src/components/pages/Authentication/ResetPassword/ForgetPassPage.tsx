import React, { useCallback, useState } from 'react'
import { AuthLayout } from 'components/layouts'
import { EmailEntry, Confirmation } from './components'

type PageState = 'EmailEntry' | 'Confirmation'

export const ForgetPassPage: React.FunctionComponent = (): JSX.Element => {
    const [pageState, setPageState] = useState<PageState>('EmailEntry')

    const handlePageChangeState = useCallback(() => {
        setPageState('Confirmation')
    }, [setPageState])

    return (
        <AuthLayout>
            { pageState === 'EmailEntry' ? <EmailEntry handlePageChangeState={handlePageChangeState} /> : <Confirmation /> }
        </AuthLayout>
    )
}

