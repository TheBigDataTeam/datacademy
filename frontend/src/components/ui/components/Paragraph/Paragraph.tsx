import React from 'react'
import classNames from 'classnames'
import { Color, Size, Text } from 'components/ui/components/Text/Text'
import { Margin } from 'components/ui/common/margins'
import marginStyles from 'components/ui/styles/margins.module.css'
import lineHeightStyles from 'components/ui/styles/line-heights.module.css'
import styles from './Paragraph.module.css'

export type Align = 'left' | 'center' | 'right'

type WithoutClassName = Omit<React.HTMLAttributes<HTMLElement>, 'className'>

export interface Props extends WithoutClassName {
  /** Text size */
  size?: Size
  /** Text color */
  color?: Color
  /** Text alignment */
  align?: Align
  /** Top and bottom margin */
  margin?: Margin
  /** Top margin */
  marginTop?: Margin
  /** Bottom margin */
  marginBottom?: Margin
  /** Data marker */
  marker?: string
}

export const Paragraph: React.FunctionComponent<Props> = ({
  children,
  size = 'm',
  color = 'default',
  align = 'left',
  margin = 'xs',
  marginTop,
  marginBottom,
  ...props
}) => {
  const className = classNames(
    styles[`root_align_${align}`],
    lineHeightStyles[`line-height_${size}`],
    marginStyles[`top_${marginTop || margin}`],
    marginStyles[`bottom_${marginBottom || margin}`]
  )

  return (
    <p {...props} className={className}>
      <Text size={size} color={color}>
        {children}
      </Text>
    </p>
  )
}
