import { VerbRef } from '../../protos/xyz/block/ftl/v1/schema/schema_pb'

export const buildVerbSchema = (verbSchema: string, dataScemas: string[]): string => {
  return dataScemas.join('\n\n') + '\n\n' + verbSchema
}

export const verbRefString = (verb: VerbRef): string => {
  return `${verb.module}.${verb.name}`
}