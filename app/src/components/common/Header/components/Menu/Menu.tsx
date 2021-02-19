import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Menu.module.css';
import { menuItems } from './menuItems';
import { Text } from '../../../../ui';

export const Menu: React.FunctionComponent = (): JSX.Element => {
	return (
		<ul className={styles.root}>
			{menuItems.map((item, i) => (
				<li key={i}>
					<Link to={item.url} className={styles.root_item}>
						<Text color="inverted" size="l">{item.title}</Text>
					</Link>
				</li>
			))}
		</ul>
	);
};
