import React, { useCallback, useState } from 'react'
import { AuthLayout } from 'components/layouts'
import { EmailEntry, Confirmation } from './components'
import { useDocTitle } from 'components/hooks'
import { TITLE_PREFIX } from 'constants/common'


type PageState = 'EmailEntry' | 'Confirmation'

export const ForgetPassPage: React.FunctionComponent = (): JSX.Element => {

    const [pageState, setPageState] = useState<PageState>('EmailEntry')
    
    useDocTitle(TITLE_PREFIX + pageState)

    const handlePageChangeState = useCallback(() => {
        setPageState('Confirmation')
    }, [setPageState])

    return (
        <AuthLayout>
            { pageState === 'EmailEntry' ? <EmailEntry handlePageChangeState={handlePageChangeState} /> : <Confirmation /> }
        </AuthLayout>
    )
}

