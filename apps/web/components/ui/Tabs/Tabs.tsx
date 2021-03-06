import type { FC, ReactElement } from 'react'
import { useState } from 'react'
import { TabButton } from './TabButton'

interface Props {
  children: (ReactElement | null)[]
  initialTabIndex?: number
}

export const Tabs: FC<Props> = ({ children, initialTabIndex = 0 }) => {
  const [selectedTabIndex, setSelectedTabIndex] = useState(initialTabIndex)

  return (
    <>
      <ul className="tabs">
        {children.map((item, index) => {
          if (!item) return null
          return (
            <li key={index}>
              <TabButton
                title={item.props.title}
                index={index}
                setSelectedTab={setSelectedTabIndex}
                selected={index === selectedTabIndex}
              />
            </li>
          )
        })}
      </ul>
      <div className="h-full">{children[selectedTabIndex]}</div>
    </>
  )
}
