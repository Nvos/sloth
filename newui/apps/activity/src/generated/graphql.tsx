import gql from 'graphql-tag';
import * as Urql from 'urql';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: string;
};

export type Activity = Node & {
  __typename?: 'Activity';
  id: Scalars['ID'];
  /** Date time range indicating when activity happend */
  range: DateTimeRange;
  description: Scalars['String'];
};

export type NotFound = Error & {
  __typename?: 'NotFound';
  message: Scalars['String'];
};

export type ActivityCreateResult = Activity | ValidationError;

export type ActivityEdge = {
  __typename?: 'ActivityEdge';
  node: Activity;
  cursor: Scalars['String'];
};

export type User = Node & {
  __typename?: 'User';
  id: Scalars['ID'];
  username: Scalars['String'];
};

export type Node = {
  id: Scalars['ID'];
};

export type Error = {
  message: Scalars['String'];
};

export type DateTimeRange = {
  __typename?: 'DateTimeRange';
  startedAt: Scalars['Time'];
  endedAt: Scalars['Time'];
};

export type UserCreateInput = {
  username: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  node?: Maybe<Node>;
  activities: ActivityConnection;
};

export type QueryNodeArgs = {
  id: Scalars['ID'];
};

export type QueryActivitiesArgs = {
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
};

export type ActivityCreateInput = {
  startedAt: Scalars['Time'];
  endedAt: Scalars['Time'];
  description: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  activityCreate: ActivityCreateResult;
  userCreate: User;
};

export type MutationActivityCreateArgs = {
  input: ActivityCreateInput;
};

export type MutationUserCreateArgs = {
  input: UserCreateInput;
};

export type ActivityConnection = {
  __typename?: 'ActivityConnection';
  edges: Array<ActivityEdge>;
};

export type ValidationError = Error & {
  __typename?: 'ValidationError';
  message: Scalars['String'];
};

export type AppActivityCreateMutationVariables = Exact<{
  startedAt: Scalars['Time'];
  endedAt: Scalars['Time'];
  description: Scalars['String'];
}>;

export type AppActivityCreateMutation = { __typename?: 'Mutation' } & {
  activityCreate:
    | ({ __typename: 'Activity' } & Pick<Activity, 'id' | 'description'> & {
          range: { __typename?: 'DateTimeRange' } & Pick<
            DateTimeRange,
            'startedAt' | 'endedAt'
          >;
        })
    | ({ __typename: 'ValidationError' } & Pick<ValidationError, 'message'>);
};

export type AppActivitiesQueryVariables = Exact<{
  first?: Maybe<Scalars['Int']>;
  last?: Maybe<Scalars['Int']>;
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
}>;

export type AppActivitiesQuery = { __typename?: 'Query' } & {
  activities: { __typename?: 'ActivityConnection' } & {
    edges: Array<
      { __typename?: 'ActivityEdge' } & Pick<ActivityEdge, 'cursor'> & {
          node: { __typename?: 'Activity' } & Pick<
            Activity,
            'id' | 'description'
          > & {
              range: { __typename?: 'DateTimeRange' } & Pick<
                DateTimeRange,
                'startedAt' | 'endedAt'
              >;
            };
        }
    >;
  };
};

export const AppActivityCreateDocument = gql`
  mutation AppActivityCreate(
    $startedAt: Time!
    $endedAt: Time!
    $description: String!
  ) {
    activityCreate(
      input: {
        startedAt: $startedAt
        endedAt: $endedAt
        description: $description
      }
    ) {
      __typename
      ... on Activity {
        id
        range {
          startedAt
          endedAt
        }
        description
      }
      ... on ValidationError {
        message
      }
    }
  }
`;

export function useAppActivityCreateMutation() {
  return Urql.useMutation<
    AppActivityCreateMutation,
    AppActivityCreateMutationVariables
  >(AppActivityCreateDocument);
}
export const AppActivitiesDocument = gql`
  query AppActivities(
    $first: Int
    $last: Int
    $after: String
    $before: String
  ) {
    activities(first: $first, last: $last, after: $after, before: $before) {
      edges {
        node {
          id
          range {
            startedAt
            endedAt
          }
          description
        }
        cursor
      }
    }
  }
`;

export function useAppActivitiesQuery(
  options: Omit<Urql.UseQueryArgs<AppActivitiesQueryVariables>, 'query'> = {}
) {
  return Urql.useQuery<AppActivitiesQuery>({
    query: AppActivitiesDocument,
    ...options,
  });
}
