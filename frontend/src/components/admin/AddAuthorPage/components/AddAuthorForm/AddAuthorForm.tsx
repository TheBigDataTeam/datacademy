import React from 'react'
import { Button, Form } from 'components/ui'
import { InputRow } from './components'
import { Author } from 'models'

interface Fields {
  email: string
  fullname: string
  bio: string
  location: string
  facebook: string
  instagram: string
  twitter: string
  shortdescription: string
  features: string

}

interface Props {
  data: Author
  errors: Partial<Fields>
  error?: string
  isLoading: boolean
  onChange: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>
  onSubmit: React.FormEventHandler<HTMLFormElement>
}

const FIELDS: (keyof Fields)[] = [
  'email',
  'fullname',
  'bio',
  'location',
  'facebook',
  'instagram',
  'twitter',
  'shortdescription',
  'features',
]
const LABELS = [
  'Email',
  'Name and Surname',
  'BIO',
  'Place of living',
  'Facebook',
  'Instagram',
  'Twitter',
  'Short description',
  'List of features'
]

export const AddAuthorForm: React.FunctionComponent<Props> = ({ data, errors, error, isLoading, onChange, onSubmit }: Props): JSX.Element => {
  
  const focusField = getFocusField(errors, isLoading)

  return (
    <Form title='New author' error={error} onSubmit={onSubmit}>
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
          Add author
        </Button>
      </Form.ActionRow>
    </Form>
  )
}

function getFocusField(errors: Partial<Fields>, isLoading: boolean): keyof Fields | null {
  if (isLoading) {
    return null
  }

  for (const f of FIELDS) {
    if (f in errors) {
      return f
    }
  }

  return FIELDS[0]
}
