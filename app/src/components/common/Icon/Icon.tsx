import React from 'react';
import styles from './Icon.module.css';
import classNames from 'classnames';
import { ReactComponent as LoginIcon } from './resources/login.svg';
import { ReactComponent as LogoutIcon } from './resources/logout.svg';
import { ReactComponent as BurgerIcon } from './resources/burger.svg';
import { ReactComponent as CloseIcon} from './resources/close.svg';
import { ReactComponent as Twitter } from './resources/twitter.svg';
import { ReactComponent as Facebook } from './resources/facebook.svg';
import { ReactComponent as Google } from './resources/google.svg';
import { ReactComponent as Github } from './resources/github.svg';
import { ReactComponent as Linkedin} from './resources/linkedin.svg';
import { ReactComponent as Dollar } from './resources/dollar.svg';
import { ReactComponent as Copyright } from './resources/copyright.svg'

interface Props {
    type: 'login' | 'logout' | 'burger' | 'close' | 'twitter' | 
    'facebook' | 'google' | 'github' | 'linkedin' | 'dollar' | 
    'copyright';
    size?: 's' | 'm' | 'l';
    color?: boolean;
    inverted?: boolean;
} 

export const Icon: React.FunctionComponent<Props> = ({type, size='m', color, inverted}): JSX.Element => {

    const className = classNames(styles[`root_size_${size}`], {
        [styles.root_color_facebook]: color && type === 'facebook',
        [styles.root_color_twitter]: color && type === 'twitter',
        [styles.root_color_inverted]: inverted,
    })

    switch (type) {
        case 'login': {
            return <LoginIcon className={className}/>
        }
        case 'logout': {
            return <LogoutIcon className={className}/>
        }
        case 'burger': {
            return <BurgerIcon className={className}/>
        }
        case 'close': {
            return <CloseIcon className={className}/>
        }
        case 'twitter': {
            return <Twitter className={className}/>
        }
        case 'facebook': {
            return <Facebook className={className}/>
        }
        case 'google': {
            return <Google className={className}/>
        }
        case 'github': {
            return <Github className={className}/>
        }
        case 'linkedin': {
            return <Linkedin className={className}/>
        }
        case 'dollar': {
            return <Dollar className={className}/>
        }
        /* TODO: possibly remove if find no usage */
        case 'copyright': {
            return <Copyright className={className}/>
        }
        default: {
            return null;
        }
    }
}
