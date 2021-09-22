import React from 'react'
import { Link } from 'react-router-dom'
import { Logo } from 'components/common'
import { Menu, Login, Logout } from './components'
import styles from './Header.module.css'

interface Props {
	inverted?: boolean
}

export const Header: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {

	/* TODO */
	const user = false

	return (
		<div className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			{ inverted ? <Menu inverted /> : <Menu /> }
			{user ?
			<Link to="/auth/login" className={styles.right}>
				{ inverted ? <Logout inverted /> : <Logout /> }
			</Link>
			: <Link to="/" className={styles.right}>
			{ inverted ? <Login inverted /> : <Login /> }
			</Link>
			}
		</div>
	)
}
