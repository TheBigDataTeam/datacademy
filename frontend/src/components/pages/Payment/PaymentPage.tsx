import { useDocTitle } from 'components/hooks'
import { TITLE_PREFIX } from 'constants/common'
import React from 'react'

export const PaymentPage: React.FunctionComponent = (): JSX.Element => {

    useDocTitle(TITLE_PREFIX + 'Payment')

    return (
        <h2>We are working on this service!</h2>
    )
}