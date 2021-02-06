import { Link } from 'react-router-dom';
import styles from './Menu.module.css';
import { Grid } from '../../../ui';
import { menuItems } from './menuItems/menuItems';

export const Menu: React.FunctionComponent = () => {
    return (
            <Grid.Col>
                <ul className={styles.root}>
                    {menuItems.map((item, i) => (
                        <li key={i}>
                            <Link to={item.url} className={styles.root_item}>
                                {item.title}
                            </Link>
                        </li>
                    ))}
                </ul>
            </Grid.Col>
    )
};
