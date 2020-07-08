import {
  createClient,
  dedupExchange,
  fetchExchange,
  debugExchange,
} from 'urql';
import {
  cacheExchange as createCacheExchange,
  UpdateResolver,
  Variables,
  ResolveInfo,
} from '@urql/exchange-graphcache';
import { devtoolsExchange } from '@urql/devtools';
import schema from '../graphql.schema.json';
import {
  ActivityCreateResult,
  AppActivitiesDocument,
  AppActivitiesQuery,
  Activity,
} from './generated/graphql';

const cacheExchange = createCacheExchange({
  // TODO: verify if generated schema is really incompatible with urql
  schema: schema as any,
  keys: {
    DateTimeRange: () => null,
  },
  updates: {
    Mutation: {
      activityCreate: (result, args, cache, info) => {
        const typed = (result as unknown) as {
          activityCreate: ActivityCreateResult;
        };
        console.log(typed);
        if (typed.activityCreate.__typename === 'Activity') {
          cache.updateQuery({ query: AppActivitiesDocument }, (data) => {
            console.log('?!', data);
            if (!data) return data;
            console.log(data);
            (data as AppActivitiesQuery).activities.edges.push({
              cursor: result.id as string,
              node: result as Activity,
            });

            return data;
          });
        }
      },
    },
  },
});

export const client = createClient({
  preferGetMethod: true,
  url: 'http://localhost:8080/query',
  exchanges: [
    debugExchange,
    devtoolsExchange,
    dedupExchange,
    cacheExchange,
    fetchExchange,
  ],
});
