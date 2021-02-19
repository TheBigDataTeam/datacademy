import React from 'react';
import { Link } from 'react-router-dom';
import { Logo } from '../';
import { Menu } from './components';
import { Login } from './components';
import styles from './Header.module.css';

interface Props {
	inverted?: boolean;
}

export const Header: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {

	const user = true;

	return (
		<div className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			<Menu />
			{/* TODO Logout Icon Component */}
			{user && <Link to="/login" className={styles.right}>
				<Login />
			</Link>}
		</div>
	);
};
