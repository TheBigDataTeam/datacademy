import React from 'react'
import { Button, Form } from 'components/ui'
import { InputRow } from './components'
import { Module } from 'models'

interface Fields {
  title: string
  body: string
  badge: string
}

interface Props {
  data: Module;
  errors: Partial<Fields>;
  error?: string;
  isLoading: boolean;
  onChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;
  onSubmit: React.FormEventHandler<HTMLFormElement>;
}

const FIELDS: (keyof Fields)[] = [
  'title',
  'body',
  'badge',
]
const LABELS = [
  'Title',
  'Body',
  'Badge',
]

export const AddModuleForm: React.FunctionComponent<Props> = ({ data, errors, error, isLoading, onChange, onSubmit }: Props): JSX.Element => {
  
  const focusField = getFocusField(errors, isLoading)

  return (
    <Form title='New module' error={error} onSubmit={onSubmit}>
      {FIELDS.map((f, i) => (
        <InputRow
          key={f}
          label={LABELS[i]}
          name={f}
          value={data[f]}
          disabled={isLoading}
          autoFocus={focusField === f}
          onChange={onChange}
        />
      ))}
      <Form.ActionRow>
        <Button type='submit' loading={isLoading} design='primary'>
          Add module
        </Button>
      </Form.ActionRow>
    </Form>
  )
}

function getFocusField(errors: Partial<Fields>, isLoading: boolean): keyof Fields | null {
  if (isLoading) {
    return null;
  }

  for (const f of FIELDS) {
    if (f in errors) {
      return f;
    }
  }

  return FIELDS[0];
}
