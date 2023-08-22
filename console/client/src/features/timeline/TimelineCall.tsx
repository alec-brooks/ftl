import { ArrowRightOnRectangleIcon } from '@heroicons/react/24/outline'
import { Call } from '../../protos/xyz/block/ftl/v1/console/console_pb'
import { formatDuration, formatTimestamp } from '../../utils/date.utils'
import { panelColor, textColor } from '../../utils/style.utils'

type Props = {
  call: Call
}

export const TimelineCall: React.FC<Props> = ({ call }) => {
  return (
    <>
      <div className={`relative flex h-6 w-6 flex-none items-center justify-center ${panelColor}`}>
        <ArrowRightOnRectangleIcon className='h-6 w-6 text-indigo-500'
          aria-hidden='true'
        />
      </div>
      <div className={`flex-auto py-0.5 text-xs leading-5 ${textColor}`}>
        {call.destinationVerbRef && (
          <div className={`inline-block rounded-md dark:bg-gray-700/40 px-2 py-1 mr-1 text-xs font-medium text-gray-500 dark:text-gray-400 ring-1 ring-inset ring-black/10 dark:ring-white/10`}>
            {call.destinationVerbRef?.module}:{call.destinationVerbRef?.name}
          </div>
        )}

        <span className=' mr-1'>
          called
        </span>

        {call.sourceVerbRef?.module && (
          <>
            from
            <div className={`inline-block rounded-md dark:bg-gray-700/40 px-2 py-1 ml-1 mr-1 text-xs font-medium text-gray-500 dark:text-gray-400 ring-1 ring-inset ring-black/10 dark:ring-white/10`}>
              {call.sourceVerbRef?.module}:{call.sourceVerbRef?.name}
            </div>
          </>
        )}
        ({formatDuration(call.duration)}).
      </div>
      <time
        dateTime={formatTimestamp(call.timeStamp)}
        className='flex-none py-0.5 text-xs leading-5 text-gray-500'
      >
        {formatTimestamp(call.timeStamp)}
      </time>
    </>
  )
}