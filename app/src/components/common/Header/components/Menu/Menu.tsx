import React from 'react';
import { Link } from 'react-router-dom';
import styles from './Menu.module.css';
import { menuItems } from './menuItems';
import { Text } from 'components/ui';

interface Props {
	inverted?: boolean
}

export const Menu: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {
	return (
		<ul className={styles.root}>
			{menuItems.map((item, i) => (
				<li key={i}>
					<Link to={item.url} className={styles.root_item}>
						{ inverted ? <Text color="inverted" size="l">{item.title}</Text>
						: <Text size="l">{item.title}</Text> }
					</Link>
				</li>
			))}
		</ul>
	);
};
