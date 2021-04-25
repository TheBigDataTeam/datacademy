import React from 'react';
import { Link } from 'react-router-dom';
import {ReactComponent as LogoIcon} from './resources/logo_datalearn.svg';
import { HOMEPAGE_URL } from 'constants/common';
import styles from './Logo.module.css';
import classNames from 'classnames';

interface Props {
  inverted?: boolean;
  size?: 's' | 'm' | 'l';
}

export const Logo: React.FunctionComponent<Props> = ({ inverted, size = 'm' }): JSX.Element => {

  const className = classNames(styles.logo, styles[`logo_size_${size}`], {
    [styles.logo_invertedColor]: inverted})

  return (
      <Link to={HOMEPAGE_URL} className={styles.link} >
        <LogoIcon className={className} />
      </Link>
  );
};
