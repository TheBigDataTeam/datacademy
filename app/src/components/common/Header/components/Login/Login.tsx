import React from 'react';
import { Icon } from 'components/common';

interface Props {
	inverted?: boolean
}

export const Login: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {
	return (
		inverted ? <Icon type="login" size="l" inverted /> : <Icon type="login" size="l" />
	);
};
