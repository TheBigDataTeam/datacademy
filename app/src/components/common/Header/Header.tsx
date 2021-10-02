import React from 'react'
import { Link } from 'react-router-dom'
import { Logo } from 'components/common'
import { Menu, Login, Logout } from './components'
import styles from './Header.module.css'
import { useSelector } from 'react-redux'

interface Props {
	inverted?: boolean
}

export const Header: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	const isLoaded: boolean = useSelector((state: any) => state.userAuth.isLoaded)

	return (
		<div className={styles.root}>
			<div className={styles.left}>
				<Logo inverted={inverted} />
			</div>
			{ inverted ? <Menu inverted /> : <Menu /> }
			{isLoaded ?
			<Link to="/auth/login" className={styles.right}>
				{ inverted ? <Logout inverted /> : <Logout /> }
			</Link>
			: <Link to="/auth/logout" className={styles.right}>
			{ inverted ? <Login inverted /> : <Login /> }
			</Link>
			}
		</div>
	)
}
