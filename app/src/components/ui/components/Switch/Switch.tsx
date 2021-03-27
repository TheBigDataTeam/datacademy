import React from 'react';
import className from 'classnames';
import styles from './Switch.module.css';

type Size = 's' | 'm' | 'l';

interface Props {
    size?: Size
    rounded?: boolean
    toggled: boolean
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    onToggle: any
}

export const Switch: React.FunctionComponent<Props> = ({ size = 'm', rounded, toggled, onToggle }): JSX.Element => {

    const sliderClassNames = className(styles.slider, styles[`slider_size_${size}`], {
        [styles.rounded]: rounded });

    const switchClassNames = className(styles.switch, styles[`switch_size_${size}`]);

    return (
        <label className={switchClassNames}>
            <input type="checkbox" checked={toggled} onChange={onToggle}/>
            <span className={sliderClassNames} />
        </label>
    );
}