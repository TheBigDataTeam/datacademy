import React from 'react'
import { Link } from 'react-router-dom'
import styles from './Menu.module.css'
import { menuItems } from './menuItems'

interface Props {
	inverted?: boolean
}

export const Menu: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {
	return (
		<ul className={styles.root}>
			{menuItems.map((item, i) => (
				<li key={i}>
					<Link to={item.url}>
						{ inverted ? <span className={styles.menu_item_inverted}>{item.title}</span>
						: <span className={styles.menu_item}>{item.title}</span> }
					</Link>
				</li>
			))}
		</ul>
	)
}
