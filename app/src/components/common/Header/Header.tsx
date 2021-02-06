/* import { Logo } from '../index'; */
import { Menu } from './MenuItems';
/* import { Login } from '../index'; */
import styles from './Header.module.css';

export const Header: React.FunctionComponent = () => {
    return (
        <div className={styles.root}>
            {/* TODO LogoIcon */}
            <div>Logo</div>
            <Menu />
            {/* TODO Login/Logout Icon */}
            <div>LoginIcon</div>
            {/* <Login /> */}
        </div>
    )
}
