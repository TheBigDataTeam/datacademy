import React from 'react'
import { Icon } from 'components/common'

interface Props {
	inverted?: boolean
}

export const Logout: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {
	return (
		inverted ? <Icon type="logout" size="l" inverted /> : <Icon type="logout" size="l" />
	)
}
