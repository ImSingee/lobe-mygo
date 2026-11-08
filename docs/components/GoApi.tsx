import type { CSSProperties } from 'react';

interface ApiField {
  default: string;
  doc: string;
  name: string;
  type: string;
  values?: string[];
}

type ApiData = Record<string, { fields: ApiField[] }>;

interface GoApiProps {
  data: ApiData;
  name: string;
}

const mono: CSSProperties = { fontFamily: 'var(--docs-font-mono)', fontSize: '0.75rem' };

const row: CSSProperties = {
  borderBlockStart: '1px solid var(--docs-border-subtle)',
  display: 'grid',
  gap: '0.25rem 1.5rem',
  gridTemplateColumns: 'minmax(8rem, 12rem) 1fr',
  paddingBlock: '0.875rem',
};

export function GoApi({ data, name }: GoApiProps) {
  const api = data[name];
  if (!api) throw new Error(`GoApi: no ${name} in the generated api.json`);

  return (
    <dl style={{ color: 'var(--docs-text-secondary)', marginBlock: '0.75rem 2rem' }}>
      {api.fields.map((field) => (
        <div key={field.name} style={row}>
          <dt>
            <code style={{ ...mono, color: 'var(--docs-text-primary)', fontWeight: 600 }}>{field.name}</code>
          </dt>
          <dd style={{ display: 'grid', gap: '0.375rem', margin: 0 }}>
            <code style={{ ...mono, color: 'var(--docs-accent)' }}>
              {field.values ? field.values.join(' | ') : field.type}
            </code>
            {field.doc ? <p style={{ fontSize: '0.875rem', margin: 0 }}>{field.doc.replaceAll('\n', ' ')}</p> : null}
            {field.default ? (
              <p style={{ fontSize: '0.8125rem', margin: 0 }}>
                Defaults to <code style={{ ...mono, color: 'var(--docs-text-primary)' }}>{field.default}</code>.
              </p>
            ) : null}
          </dd>
        </div>
      ))}
    </dl>
  );
}
