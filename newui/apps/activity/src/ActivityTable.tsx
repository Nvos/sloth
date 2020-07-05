import React, { useState, FC, useEffect } from 'react';
import {
  useAppActivityCreateMutation,
  useAppActivitiesQuery,
  ActivityCreateInput,
  Activity,
} from './generated/graphql';
import { useForm, SubmitHandler } from 'react-hook-form';

const ActivityTable: FC = () => {
  const [error, setError] = useState('');
  const { register, handleSubmit, errors } = useForm<ActivityCreateInput>();
  const [activities] = useAppActivitiesQuery();
  const [_, execute] = useAppActivityCreateMutation();

  const onSubmit: SubmitHandler<ActivityCreateInput> = async (data) => {
    console.log(data);
    await execute({
      description: data.description,
      endedAt: data.endedAt + 'Z',
      startedAt: data.startedAt + 'Z',
    }).then((result) => {
      if (result.data) {
        if (result.data.activityCreate.__typename === 'Activity') {
          setError('');
        } else {
          setError(result.data?.activityCreate.message);
        }
      }
    });
  };

  return (
    <div>
      <div>
        <form onSubmit={handleSubmit(onSubmit)}>
          <label style={{ display: 'block' }}>
            description
            <input
              ref={register}
              name="description"
              placeholder="description"
              type="text"
            />
          </label>
          <label style={{ display: 'block' }}>
            startedAt
            <input
              ref={register}
              name="startedAt"
              placeholder="startedAt"
              step=".1"
              type="datetime-local"
            />
          </label>
          <label style={{ display: 'block' }}>
            endedAt
            <input
              ref={register}
              name="endedAt"
              step=".1"
              placeholder="endedAt"
              type="datetime-local"
            />
          </label>

          <button>Submit</button>
        </form>
      </div>
      <div>{error && `error=${error}`}</div>
      <table>
        <tbody>
          {activities.data?.activities.edges
            .map((it) => it.node)
            .map((it) => (
              <tr>
                <td>{it.id}</td>
                <td>{it.range.startedAt}</td>
                <td>{it.range.endedAt}</td>
                <td>{it.description}</td>
              </tr>
            ))}
        </tbody>
      </table>
    </div>
  );
};

export { ActivityTable };
