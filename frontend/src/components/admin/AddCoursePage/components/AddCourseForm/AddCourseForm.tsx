import React from 'react'
import { Button, Form } from 'components/ui'
import { InputRow } from './components'
import { Course } from 'models'

interface Fields {
  title: string
  author: string
  description: string
  techstack: string
  moduleQuantity: string
  workshopQuantity: string
}

interface Props {
  data: Course;
  errors: Partial<Fields>;
  error?: string;
  isLoading: boolean;
  onChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;
  onSubmit: React.FormEventHandler<HTMLFormElement>;
}

const FIELDS: (keyof Fields)[] = [
  'title',
  'author',
  'description',
  'techstack',
  'moduleQuantity',
  'workshopQuantity'
]
const LABELS = [
  'Title',
  'Author',
  'Description',
  'Techstack',
  'Quantity of modules',
  'Quantity of workshops'
]

export const AddCourseForm: React.FunctionComponent<Props> = ({ data, errors, error, isLoading, onChange, onSubmit }: Props): JSX.Element => {
  
  const focusField = getFocusField(errors, isLoading)

  return (
    <Form title='New course' error={error} onSubmit={onSubmit}>
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
          Add course
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
