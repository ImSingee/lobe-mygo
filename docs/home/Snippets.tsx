import { CodeShowcase, type CodeShowcaseItem } from '@lobehub/ui/awesome';

import { Shot } from './Shot';
import snippets from './snippets.json';

const ITEMS: CodeShowcaseItem[] = snippets.map((snippet) => ({
  code: snippet.code,
  key: snippet.key,
  label: snippet.label,
  language: 'go',
  preview: <Shot alt={`${snippet.label} rendered by Lobe MyGo`} name={snippet.image} width={480} />,
}));

export function Snippets() {
  return (
    <div className="mygo-snippets">
      <CodeShowcase items={ITEMS} />
    </div>
  );
}
