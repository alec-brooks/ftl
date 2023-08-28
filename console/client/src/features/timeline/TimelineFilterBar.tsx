import {Popover, Transition} from '@headlessui/react'
import {ChevronDownIcon} from '@heroicons/react/20/solid'
import {Fragment} from 'react'
import {logLevelText, panelColor, textColor} from '../../utils/style.utils'

const logLevels = [0, 1, 5, 9, 13, 17]

type Props = {
  selectedEventTypes: string[]
  onEventTypesChanged: (eventType: string, checked: boolean) => void
  selectedLogLevels: number[]
  onLogLevelsChanged: (logLevel: number, checked: boolean) => void
}

export const TimelineFilterBar = ({
  selectedEventTypes,
  onEventTypesChanged,
  selectedLogLevels,
  onLogLevelsChanged,
}: Props) => {
  const eventButtonStyles = `relative inline-flex items-center px-2 py-1 text-xs font-semibold ring-1 ring-inset focus:z-10`

  const eventSelectedStyles = (eventType: string) => {
    return selectedEventTypes.includes(eventType)
      ? 'bg-indigo-600 text-white ring-indigo-200 dark:ring-indigo-400'
      : 'bg-white text-gray-700 ring-indigo-200 dark:ring-indigo-400'
  }

  const toggleEventType = (eventType: string) => {
    onEventTypesChanged(eventType, !selectedEventTypes.includes(eventType))
  }

  return (
    <>
      <div className={`sticky top-0 z-10 ${panelColor} shadow`}>
        <div className='flex items-center justify-between p-4'>
          <span className='isolate inline-flex rounded-md shadow-sm'>
            <button
              type='button'
              className={`${eventButtonStyles} ${eventSelectedStyles(
                'log'
              )} rounded-l-md`}
              onClick={() => toggleEventType('log')}
            >
              Logs
            </button>
            <button
              type='button'
              className={`${eventButtonStyles} ${eventSelectedStyles(
                'call'
              )} -ml-px`}
              onClick={() => toggleEventType('call')}
            >
              Calls
            </button>
            <button
              type='button'
              className={`${eventButtonStyles} ${eventSelectedStyles(
                'deployment'
              )} -ml-px rounded-r-md`}
              onClick={() => toggleEventType('deployment')}
            >
              Deployments
            </button>
          </span>
          <Popover.Group className='hidden sm:flex sm:items-baseline sm:space-x-8'>
            <Popover
              as='div'
              key='log-levels'
              id={`desktop-menu-log-levels`}
              className='relative inline-block text-left'
            >
              <div>
                <Popover.Button
                  className={`group inline-flex items-center justify-center text-sm font-medium ${textColor} hover:text-gray-900`}
                >
                  <span>Log levels</span>
                  <span className='ml-1.5 rounded text-white bg-indigo-600 dark:bg-indigo-600 px-1.5 py-0.5 text-xs font-semibold tabular-nums'>
                    {selectedLogLevels.length}
                  </span>
                  <ChevronDownIcon
                    className='-mr-1 ml-1 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500'
                    aria-hidden='true'
                  />
                </Popover.Button>
              </div>

              <Transition
                as={Fragment}
                enter='transition ease-out duration-100'
                enterFrom='transform opacity-0 scale-95'
                enterTo='transform opacity-100 scale-100'
                leave='transition ease-in duration-75'
                leaveFrom='transform opacity-100 scale-100'
                leaveTo='transform opacity-0 scale-95'
              >
                <Popover.Panel
                  className={`absolute right-0 z-10 mt-2 origin-top-right rounded-md bg-white p-4 shadow-2xl ring-1 ring-black ring-opacity-5 focus:outline-none`}
                >
                  <form className='space-y-4'>
                    {logLevels.map(level => (
                      <div
                        key={level}
                        className='flex items-center'
                      >
                        <input
                          id={`log-level-${level}`}
                          name={`log-level-${level}`}
                          defaultValue={level}
                          type='checkbox'
                          className='h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600'
                          checked={selectedLogLevels.includes(level)}
                          onChange={e =>
                            onLogLevelsChanged(level, e.target.checked)
                          }
                        />
                        <label
                          htmlFor={`log-level-${level}`}
                          className='ml-3 whitespace-nowrap pr-6 text-sm font-medium text-gray-900'
                        >
                          {logLevelText[level]}
                        </label>
                      </div>
                    ))}
                  </form>
                </Popover.Panel>
              </Transition>
            </Popover>
          </Popover.Group>
        </div>
      </div>
    </>
  )
}
