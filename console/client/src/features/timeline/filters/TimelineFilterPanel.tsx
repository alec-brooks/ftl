import { PhoneIcon, RocketLaunchIcon } from '@heroicons/react/24/outline'
import React from 'react'
import { EventType, EventsQuery_Filter, LogLevel } from '../../../protos/xyz/block/ftl/v1/console/console_pb'
import { modulesContext } from '../../../providers/modules-provider'
import { eventTypesFilter, logLevelFilter, modulesFilter } from '../../../services/console.service'
import { textColor } from '../../../utils'
import { LogLevelBadgeSmall } from '../../logs/LogLevelBadgeSmall'
import { logLevelBgColor, logLevelColor, logLevelRingColor } from '../../logs/log.utils'
import { FilterPanelSection } from './FilterPanelSection'

interface EventFilter {
  label: string
  type: EventType
  icon: React.ReactNode
}

const EVENT_TYPES: Record<string, EventFilter> = {
  call: { label: 'Call', type: EventType.CALL, icon: <PhoneIcon className='w-4 h-4 text-indigo-600 ml-1' /> },
  log: { label: 'Log', type: EventType.LOG, icon: <LogLevelBadgeSmall logLevel={LogLevel.INFO} /> },
  deployment: {
    label: 'Deployment',
    type: EventType.DEPLOYMENT,
    icon: <RocketLaunchIcon className='w-4 h-4 text-indigo-600 ml-1' />,
  },
}

const LOG_LEVELS: Record<number, string> = {
  1: 'Trace',
  5: 'Debug',
  9: 'Info',
  13: 'Warn',
  17: 'Error',
}

interface Props {
  onFiltersChanged: (filters: EventsQuery_Filter[]) => void
}

export const TimelineFilterPanel = ({ onFiltersChanged }: Props) => {
  const modules = React.useContext(modulesContext)
  const [selectedEventTypes, setSelectedEventTypes] = React.useState<string[]>(Object.keys(EVENT_TYPES))
  const [selectedModules, setSelectedModules] = React.useState<string[]>([])
  const [selectedLogLevel, setSelectedLogLevel] = React.useState<number>(1)

  React.useEffect(() => {
    if (selectedModules.length === 0) {
      setSelectedModules(modules.modules.map((module) => module.deploymentName))
    }
  }, [modules])

  React.useEffect(() => {
    const filter: EventsQuery_Filter[] = []
    if (selectedEventTypes.length !== Object.keys(EVENT_TYPES).length) {
      const selectedTypes = selectedEventTypes.map((key) => EVENT_TYPES[key].type)

      filter.push(eventTypesFilter(selectedTypes))
    }
    if (selectedLogLevel !== LogLevel.TRACE) {
      filter.push(logLevelFilter(selectedLogLevel))
    }

    filter.push(modulesFilter(selectedModules))

    onFiltersChanged(filter)
  }, [selectedEventTypes, selectedLogLevel, selectedModules])

  const handleTypeChanged = (eventType: string, checked: boolean) => {
    if (checked) {
      setSelectedEventTypes((prev) => [...prev, eventType])
    } else {
      setSelectedEventTypes((prev) => prev.filter((filter) => filter !== eventType))
    }
  }

  const handleModuleChanged = (deploymentName: string, checked: boolean) => {
    if (checked) {
      setSelectedModules((prev) => [...prev, deploymentName])
    } else {
      setSelectedModules((prev) => prev.filter((filter) => filter !== deploymentName))
    }
  }

  const handleLogLevelChanged = (logLevel: string) => {
    setSelectedLogLevel(Number(logLevel))
  }

  return (
    <div className='flex-shrink-0 w-52'>
      <div className='w-full'>
        <div className='mx-auto w-full max-w-md p-2'>
          <FilterPanelSection title='Event types'>
            {Object.keys(EVENT_TYPES).map((key) => (
              <div key={key} className='relative flex items-start'>
                <div className='flex h-6 items-center'>
                  <input
                    id={`event-type-${key}`}
                    name={`event-type-${key}`}
                    type='checkbox'
                    checked={selectedEventTypes.includes(key)}
                    onChange={(e) => handleTypeChanged(key, e.target.checked)}
                    className='h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600 cursor-pointer'
                  />
                </div>
                <div className='ml-2 text-sm leading-6 w-full'>
                  <label
                    htmlFor={`event-type-${key}`}
                    className={`flex justify-between items-center ${textColor} cursor-pointer`}
                  >
                    {EVENT_TYPES[key].label}
                    <span>{EVENT_TYPES[key].icon}</span>
                  </label>
                </div>
              </div>
            ))}
          </FilterPanelSection>

          <FilterPanelSection title='Log level'>
            <ul role='list' className='space-y-1'>
              {Object.keys(LOG_LEVELS).map((key) => (
                <li
                  key={key}
                  onClick={() => handleLogLevelChanged(key)}
                  className='relative flex gap-x-2 cursor-pointer'
                >
                  <div className='relative flex h-5 w-3 flex-none items-center justify-center'>
                    <div
                      className={`${selectedLogLevel <= Number(key) ? 'h-2.5 w-2.5' : 'h-0.5 w-0.5'} ${
                        selectedLogLevel <= Number(key)
                          ? `${logLevelBgColor[Number(key)]} ${logLevelRingColor[Number(key)]}`
                          : 'bg-gray-300 ring-gray-300'
                      } rounded-full ring-1`}
                    ></div>
                  </div>
                  <p className='flex-auto text-sm leading-5 text-gray-500'>
                    <span className={`${logLevelColor[Number(key)]} flex`}>{LOG_LEVELS[Number(key)]}</span>
                  </p>
                </li>
              ))}
            </ul>
          </FilterPanelSection>

          <FilterPanelSection title='Modules'>
            <div className='relative flex items-center mb-2'>
              <button
                onClick={() => setSelectedModules(modules.modules.map((module) => module.deploymentName))}
                className='text-indigo-600 cursor-pointer hover:text-indigo-500'
              >
                Select All
              </button>
              <span className='px-1 text-indigo-700'>|</span>
              <button
                onClick={() => setSelectedModules([])}
                className='text-indigo-600 cursor-pointer hover:text-indigo-500'
              >
                Deselect All
              </button>
            </div>
            {modules.modules.map((module) => (
              <div key={module.deploymentName} className='relative flex items-start'>
                <div className='flex h-6 items-center'>
                  <input
                    id={`module-${module.deploymentName}`}
                    name={`module-${module.deploymentName}`}
                    type='checkbox'
                    checked={selectedModules.includes(module.deploymentName)}
                    onChange={(e) => handleModuleChanged(module.deploymentName, e.target.checked)}
                    className='h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600 cursor-pointer'
                  />
                </div>
                <div className='ml-2 text-sm leading-6 w-full'>
                  <label htmlFor={`module-${module.deploymentName}`} className={`${textColor} flex cursor-pointer`}>
                    {module.name}
                  </label>
                </div>
              </div>
            ))}
          </FilterPanelSection>
        </div>
      </div>
    </div>
  )
}