import React from 'react';
import { Link } from 'react-router-dom';
import { Logo } from 'components/common';
import { Menu } from './components';
import { Login } from './components';
import styles from './Header.module.css';

interface Props {
	inverted?: boolean;
}

export const Header: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {

	/* TODO */
	const user = true;

	return (
		<div className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			{ inverted ? <Menu inverted /> : <Menu /> }
			{/* TODO Logout Icon Component when user is false */}
			{user && <Link to="/auth/login" className={styles.right}>
				{ inverted ? <Login inverted /> : <Login /> }
			</Link>}
		</div>
	);
}
