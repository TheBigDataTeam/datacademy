import React from 'react';
import classNames from 'classnames';
import NextLink, { LinkProps } from 'next/link';
import { ButtonDesign, ButtonSize } from '../../common';
import designStyles from '../../styles/design.module.css';
import styles from './Link.module.css';

type Color = 'light' | 'dark';

// TODO Extends add Link attributes
interface Props extends LinkProps {
  design?: ButtonDesign;
  size?: ButtonSize;
  color?: Color;
  rounded?: boolean;
  underline?: boolean;
  children?: React.ReactNode;
  target?: string;
  icon?: React.ReactNode;
  onlyIcon?: boolean;
}

export const Link: React.FunctionComponent<Props> = ({
  design = 'link',
  color,
  rounded,
  size = 'm',
  underline = true,
  href,
  as,
  onlyIcon,
  children,
  icon,
  ...props
}) => {
  const isButtonDesign = design !== 'link';
  const className = classNames(styles.root, {
    [styles[`root_design_${design}`]]: !isButtonDesign,
    [designStyles[design]]: isButtonDesign,
    [styles[`root_underlined`]]: !isButtonDesign && underline,
    [styles[`root_color_${color}`]]: color,
    [styles[`root_size_${size}`]]: size && isButtonDesign,
    [styles.root_rounded]: rounded && isButtonDesign,
    [styles.root_onlyIcon]: icon && isButtonDesign,
  });

  const iconClassName = classNames(styles.icon, {
    [styles.icon_onlyIcon]: !children,
  });

  if (
    (typeof href === 'string' && href.startsWith('mailto:')) ||
    props.target
  ) {
    return <a {...props} href={href as string} className={className} />;
  }

  return (
    <NextLink href={href} as={as} passHref>
      <a {...props} href={href as string} className={className}>
        {icon && <span className={iconClassName}>{icon}</span>}
        {children}
      </a>
    </NextLink>
  );
};
