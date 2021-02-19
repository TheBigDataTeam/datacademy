import React from 'react';
import { Link } from 'react-router-dom';
import {ReactComponent as LogoIcon} from './resources/logo_datalearn.svg';
import { HOMEPAGE_URL } from '../../../constants/common';
import styles from './Logo.module.css';
import classNames from 'classnames';

interface Props {
  inverted?: boolean
}

export const Logo: React.FunctionComponent<Props> = ({ inverted }): JSX.Element => {

  const className = classNames(styles.logo, {
    [styles.logo_invertedColor]: inverted
  })

  return (
      <Link to={HOMEPAGE_URL} className={styles.link} >
        <LogoIcon className={className} />
      </Link>
  );
};
