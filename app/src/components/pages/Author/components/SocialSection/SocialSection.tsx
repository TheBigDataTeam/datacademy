import React from 'react';
import { Icon } from 'components/common';
import styles from './SocialSection.module.css';
import { Author } from 'models';

interface Props {
    author: Author
}

export const SocialSection: React.FunctionComponent<Props> = ({ author }): JSX.Element => {
    return (
        <div className={styles.root}>
            <a href={`www.facebook.com/${author.facebook}`} target="_blank" rel="noreferrer" className={styles.root_link}>
                <Icon type="facebook" color />
            </a>
            <a href={`www.twitter.com/${author.twitter}`} target="_blank" rel="noreferrer" className={styles.root_link}>
                <Icon type="twitter" color/>
            </a>
            <a href={`www.instagram.com/${author.instagram}`} target="_blank" rel="noreferrer" className={styles.root_link}>
                <Icon type="instagram"/>
            </a>
        </div>
    )
}
