import { ActionIcon, SyntaxHighlighter } from '@lobehub/ui';
import { Code2, Copy, ExternalLink } from 'lucide-react';
import { type CSSProperties, useState } from 'react';

interface GoDemoProps {
  height: number;
  name: string;
  scrollable?: boolean;
  source: string;
  title: string;
}

const frame: CSSProperties = {
  backgroundColor: 'var(--docs-surface-raised)',
  borderColor: 'var(--docs-border-default)',
  borderRadius: 'var(--docs-radius-lg)',
  borderStyle: 'solid',
  borderWidth: 1,
  boxShadow: 'var(--docs-shadow-control)',
  marginBlock: '1.75rem',
  overflow: 'hidden',
};

const caption: CSSProperties = {
  alignItems: 'center',
  borderBlockEnd: '1px solid var(--docs-border-subtle)',
  display: 'flex',
  gap: '0.5rem',
  minHeight: '2.25rem',
  paddingBlock: '0.25rem',
  paddingInline: '0.875rem 0.5rem',
};

const captionTitle: CSSProperties = {
  color: 'var(--docs-text-primary)',
  fontSize: '0.8438rem',
  fontWeight: 620,
  letterSpacing: '-0.01em',
  margin: 0,
};

export function GoDemo({ height, name, scrollable = false, source, title }: GoDemoProps) {
  const [expanded, setExpanded] = useState(false);
  const src = `/mygo/index.html?demo=${encodeURIComponent(name)}${scrollable ? '&scroll=1' : ''}`;

  return (
    <section aria-label={title} data-demo-layout="go" data-toc-ignore="" style={frame}>
      <header style={caption}>
        <p style={captionTitle}>{title}</p>
        <div style={{ display: 'flex', gap: '0.125rem', marginInlineStart: 'auto' }}>
          <ActionIcon
            active={expanded}
            aria-expanded={expanded}
            aria-label={expanded ? 'Hide Go source' : 'Show Go source'}
            icon={Code2}
            onClick={() => setExpanded((open) => !open)}
            size="small"
            title={expanded ? 'Hide Go source' : 'Show Go source'}
          />
          <ActionIcon
            aria-label="Copy Go source"
            icon={Copy}
            onClick={() => void navigator.clipboard.writeText(source)}
            size="small"
            title="Copy Go source"
          />
          <ActionIcon
            aria-label="Open preview"
            icon={ExternalLink}
            onClick={() => window.open(src, '_blank', 'noopener')}
            size="small"
            title="Open preview"
          />
        </div>
      </header>
      <iframe
        loading="lazy"
        src={src}
        style={{ border: 0, display: 'block', height, width: '100%' }}
        title={`${title} preview`}
      />
      {expanded ? (
        <div className="docs-demo-source-panel" data-pagefind-ignore="all">
          <SyntaxHighlighter className="docs-demo-source" language="go" style={{ tabSize: 4 }} tabIndex={0}>
            {source.trimEnd()}
          </SyntaxHighlighter>
        </div>
      ) : null}
    </section>
  );
}
